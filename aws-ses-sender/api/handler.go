package api

import (
	"aws-ses-sender/config"
	"aws-ses-sender/model"
	"bytes"
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

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
		var scheduledAt *time.Time
		if msg.ScheduledAt != "" {
			if t, err := time.Parse(time.DateTime, msg.ScheduledAt); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid scheduledAt format"})
			} else {
				scheduledAt = &t
			}
		}
		for _, email := range msg.Emails {
			req := &model.Request{
				TopicId:     msg.TopicId,
				To:          email,
				Subject:     msg.Subject,
				Content:     msg.Content,
				ScheduledAt: scheduledAt,
				Status:      model.EmailMessageStatusCreated,
			}
			reqs = append(reqs, req)
			totCnt++
		}
	}

	chunkSize := 1000
	for i := 0; i < len(reqs); i += chunkSize {
		end := i + chunkSize
		if end > len(reqs) {
			end = len(reqs)
		}
		if err := db.Create(reqs[i:end]).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	// Return the result
	return c.JSON(fiber.Map{
		"count":   len(reqBody.Messages),
		"elapsed": time.Since(start).String(),
	})
}

// createOpenEventHandler Open Event Handler
// Attach an image script to the email and assume it has been read when the image is accessed
func createOpenEventHandler(c fiber.Ctx) error {
	reqId := c.Query("requestId")
	if reqId != "" {
		// Consider email as opened and create data
		db := config.GetDB()
		reqIdInt, _ := strconv.Atoi(reqId)
		_ = db.Create(&model.Result{
			RequestId: uint(reqIdInt),
			Status:    "Open",
			Raw:       "{}",
		}).Error
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
// Handler that receives AWS SES results
func createResultEventHandler(c fiber.Ctx) error {
	msgType := c.Get("x-amz-sns-message-type")
	if msgType != "Notification" && msgType != "SubscriptionConfirmation" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid SNS Message Type"})
	}

	var reqBody struct {
		Type         string `json:"Type"`
		Message      string `json:"Message"`
		MessageId    string `json:"MessageId"`
		SubscribeURL string `json:"SubscribeURL"`
	}
	if err := c.Bind().JSON(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse SNS message"})
	}

	if reqBody.Type == "SubscriptionConfirmation" {
		// Subscription confirmation
		log.Printf("Subscription confirmation required. Visiting: %s", reqBody.SubscribeURL)
		return c.JSON(fiber.Map{"message": "Subscription confirmation required"})
	}

	if reqBody.Type != "Notification" {
		return c.JSON(fiber.Map{"message": "Other message type received"})
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
		return c.JSON(fiber.Map{"message": "Non-SES notification received"})
	}
	if sesNotification.Mail.MessageId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "SES message_id not found"})
	}

	// Find request_id from headers
	var reqId uint
	for _, header := range sesNotification.Mail.Headers {
		if header.Name == "X-Request-ID" {
			reqIdInt, _ := strconv.Atoi(header.Value)
			reqId = uint(reqIdInt)
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
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
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
