package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"aws-ses-sender/config"
	"aws-ses-sender/model"
	"aws-ses-sender/pkg/sns"
	"aws-ses-sender/pkg/tracking"

	"github.com/gofiber/fiber/v3"
)

// snsVerifier checks that the messages posted to /v1/events/results come from AWS SNS and from
// the topic in SNS_TOPIC_ARN. setV1Routes creates it.
var snsVerifier *sns.Verifier

// newSNSVerifier creates the verifier from SNS_TOPIC_ARN, a comma separated list of topic ARNs.
func newSNSVerifier() *sns.Verifier {
	var topics []string
	for arn := range strings.SplitSeq(config.GetEnv("SNS_TOPIC_ARN"), ",") {
		topics = append(topics, strings.TrimSpace(arn))
	}
	v := sns.NewVerifier(topics...)
	if !v.Configured() {
		log.Printf("SNS_TOPIC_ARN is not set: every message posted to /v1/events/results is rejected")
	}
	return v
}

// createMessageHandler Message Handler
// Handler that receives email sending requests
func createMessageHandler(c fiber.Ctx) error {
	start := time.Now()
	var reqBody struct {
		Messages []struct {
			TopicId     string   `json:"topicId"`
			Emails      []string `json:"emails"`
			Subject     string   `json:"subject"`
			Content     string   `json:"content"`
			ScheduledAt string   `json:"scheduledAt"`
		} `json:"messages"`
	}
	if err := c.Bind().JSON(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db := config.GetDB()
	reqs := make([]*model.Request, 0)
	totCnt := 0
	for _, msg := range reqBody.Messages {
		scheduledAt := time.Now().UTC()

		if msg.ScheduledAt != "" {
			if t, err := time.Parse(time.RFC3339, msg.ScheduledAt); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid scheduledAt format"})
			} else {
				scheduledAt = t.UTC()
				// Ensure scheduled time is in the future
				if scheduledAt.Before(time.Now().UTC()) {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "scheduledAt must be in the future"})
				}
			}
		}
		for _, email := range msg.Emails {
			// Validate email address
			if _, err := mail.ParseAddress(email); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("invalid email: %s", email)})
			}
			// Validate essential message fields
			if strings.TrimSpace(msg.Subject) == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Subject cannot be empty"})
			}
			if strings.TrimSpace(msg.Content) == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Content cannot be empty"})
			}
			req := &model.Request{
				TopicId:     msg.TopicId,
				To:          email,
				Subject:     msg.Subject,
				Content:     msg.Content,
				ScheduledAt: &scheduledAt,
				Status:      model.EmailMessageStatusCreated,
			}
			reqs = append(reqs, req)
			totCnt++
		}
	}

	chunkSize := 1000
	for i := 0; i < len(reqs); i += chunkSize {
		end := min(i+chunkSize, len(reqs))
		if err := db.Create(reqs[i:end]).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	// Return the result
	return c.JSON(fiber.Map{
		"count":   totCnt,
		"elapsed": time.Since(start).String(),
	})
}

// createOpenEventHandler Open Event Handler
// Attach an image script to the email and assume it has been read when the image is accessed.
// The link carries a token that the sender signs with API_KEY, so only the links in the emails
// that were sent count. Every request gets the image, so a mail client never shows a broken one.
func createOpenEventHandler(c fiber.Ctx) error {
	reqId := c.Query("requestId")
	if reqId != "" {
		reqIdInt, err := strconv.ParseUint(reqId, 10, 63)
		switch {
		case err != nil:
			log.Printf("Invalid requestId: %q", reqId)
		case !tracking.Valid(config.GetEnv("API_KEY"), uint(reqIdInt), c.Query("token")):
			log.Printf("Ignoring an open event for request %d: the token is missing or wrong", reqIdInt)
		default:
			// Consider email as opened and create data
			db := config.GetDB()
			if err := db.Create(&model.Result{
				RequestId: uint(reqIdInt),
				Status:    "Open",
				Raw:       "{}",
			}).Error; err != nil {
				log.Printf("Failed to create open event: %s", err.Error())
			}
		}
	}

	// Return a blank image
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{R: 0, G: 0, B: 0, A: 0})
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	c.Set("Content-Type", "image/png")
	return c.Send(buf.Bytes())
}

// createResultEventHandler Result Event Handler
// Handler that receives AWS SES results. The endpoint has to be public for SNS to reach it, so
// it only accepts a message that SNS signed and that comes from the configured topic.
func createResultEventHandler(c fiber.Ctx) error {
	msgType := c.Get("x-amz-sns-message-type")
	if msgType != sns.TypeNotification && msgType != sns.TypeSubscriptionConfirmation {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid SNS Message Type"})
	}

	var reqBody sns.Message
	if err := json.Unmarshal(c.Body(), &reqBody); err != nil || reqBody.Type != msgType {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse SNS message"})
	}
	if err := snsVerifier.Verify(c.Context(), &reqBody); err != nil {
		log.Printf("Rejected an SNS message: %v", err)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Invalid SNS message"})
	}

	if reqBody.Type == sns.TypeSubscriptionConfirmation {
		// Subscription confirmation. Nothing is done with the URL here, open it to confirm.
		log.Printf("Subscription confirmation required. Open this URL to confirm: %s", reqBody.SubscribeURL)
		return c.JSON(fiber.Map{"message": "Subscription confirmation required"})
	}

	var sesNotification struct {
		NotificationType string `json:"notificationType"`
		Mail             struct {
			MessageId string `json:"messageId"`
			Headers   []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"headers"`
		} `json:"mail"`
	}
	if err := json.Unmarshal([]byte(reqBody.Message), &sesNotification); err != nil {
		return c.JSON(fiber.Map{"message": fmt.Sprintf("Non-SES notification received: %v", err)})
	}
	if sesNotification.Mail.MessageId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "SES message_id not found"})
	}

	// Log notification type for monitoring
	log.Printf("Received %s notification from SES for message %s", sesNotification.NotificationType, sesNotification.Mail.MessageId)

	// Find request_id from headers
	var reqId uint
	for _, header := range sesNotification.Mail.Headers {
		if strings.EqualFold(header.Name, "X-Request-ID") {
			if reqIdInt, err := strconv.Atoi(header.Value); err == nil {
				reqId = uint(reqIdInt)
			} else {
				log.Printf("Invalid requestId: %s", header.Value)
			}
			break
		}
	}

	if reqId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request ID not found in headers"})
	}

	db := config.GetDB()
	if err := db.Create(&model.Result{
		RequestId: reqId,
		Status:    sesNotification.NotificationType,
		Raw:       reqBody.Message,
	}).Error; err != nil {
		log.Printf("Failed to save result event: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save event"})
	}
	return c.JSON(fiber.Map{"message": "OK"})
}

// getResultCountHandler Retrieve email delivery results as counts
func getResultCountHandler(c fiber.Ctx) error {
	topicID := c.Params("topicId")
	if topicID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "topicId is required"})
	}

	db := config.GetDB()

	// Check if any requests exist for the given topicID.  Early exit if none.
	var reqCnt int64
	if err := db.Model(&model.Request{}).Where("topic_id = ?", topicID).Count(&reqCnt).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if reqCnt == 0 {
		return c.JSON(fiber.Map{
			"request": fiber.Map{"total": 0, "created": 0, "sent": 0, "failed": 0, "stopped": 0},
			"result":  fiber.Map{"total": 0, "statuses": map[string]int{}},
		})
	}

	// --- Request Counts (Efficient Single Query) ---
	var reqResults []struct {
		Status int
		Count  int
	}
	if err := db.Model(&model.Request{}).
		Select("status, COUNT(*) as count").
		Where("topic_id = ?", topicID).
		Group("status").
		Scan(&reqResults).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	reqCnts := struct {
		Total   int `json:"total"`
		Created int `json:"created"`
		Sent    int `json:"sent"`
		Failed  int `json:"failed"`
		Stopped int `json:"stopped"`
	}{Total: int(reqCnt)} // Initialize Total with requestCount

	for _, r := range reqResults {
		switch r.Status {
		case model.EmailMessageStatusCreated:
			reqCnts.Created = r.Count
		case model.EmailMessageStatusSent:
			reqCnts.Sent = r.Count
		case model.EmailMessageStatusFailed:
			reqCnts.Failed = r.Count
		case model.EmailMessageStatusStopped:
			reqCnts.Stopped = r.Count
		}
	}

	// --- Result Counts (Optimized with Subquery) ---

	// Use a subquery to get the distinct request IDs associated with the topicID.
	// This is generally the most efficient approach with GORM and avoids extra Go-side processing.
	subQuery := db.Model(&model.Request{}).Select("id").Where("topic_id = ?", topicID)

	var resultResults []struct {
		Status string
		Count  int
	}
	if err := db.Model(&model.Result{}).
		Select("status, COUNT(DISTINCT request_id) as count").
		Where("request_id IN (?)", subQuery).
		Group("status").
		Scan(&resultResults).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Use a map for flexible status handling.
	resultCounts := make(map[string]int)
	for _, r := range resultResults {
		resultCounts[r.Status] = r.Count
	}

	// --- Return Combined Result ---
	return c.JSON(fiber.Map{
		"request": reqCnts,
		"result": fiber.Map{
			"statuses": resultCounts,
		},
	})
}

// getSentCountHandler Retrieve the number of emails sent within 24 hours
func getSentCountHandler(c fiber.Ctx) error {
	// Receive hours as a query string (default 24)
	hours, err := strconv.Atoi(c.Query("hours", "24"))
	if err != nil || hours <= 0 || hours > 168 {
		log.Printf("Invalid hours: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "hours must be 1-168"})
	}

	// Calculate the time hours before the current time
	startTime := time.Now().Add(-time.Duration(hours) * time.Hour)

	// Get the number of emails after startTime from the DB
	db := config.GetDB()
	var cnt int64
	if err := db.Model(&model.Request{}).
		Where("updated_at > ?", startTime).
		Where("status = ?", model.EmailMessageStatusSent).
		Count(&cnt).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Return the result
	return c.JSON(fiber.Map{"count": cnt})
}
