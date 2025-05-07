package aws

import (
	"aws-ses-sender/config"
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

// SES is a wrapper around the AWS SES client
type SES struct {
	Client *sesv2.Client
}

// NewSESClient creates an email client
func NewSESClient(ctx context.Context) (*SES, error) {
	accessKeyID := config.GetEnv("AWS_ACCESS_KEY_ID")
	secretAccessKey := config.GetEnv("AWS_SECRET_ACCESS_KEY")
	var cfgOpts []func(*awsConfig.LoadOptions) error
	cfgOpts = append(cfgOpts, awsConfig.WithRegion(config.GetEnv("AWS_REGION", "ap-northeast-2")))

	// If accessKeyID and secretAccessKey are set, use static credentials
	if accessKeyID != "" && secretAccessKey != "" {
		cfgOpts = append(cfgOpts, awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, ""),
		))
	}
	cfg, err := awsConfig.LoadDefaultConfig(ctx, cfgOpts...)
	if err != nil {
		return nil, err
	}
	return &SES{
		Client: sesv2.NewFromConfig(cfg),
	}, nil
}

// SendEmail sends an email
func (s *SES) SendEmail(ctx context.Context, reqID int, subject, body *string, receivers []string) (string, error) {
	sender := config.GetEnv("EMAIL_SENDER")
	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(sender),
		Destination: &types.Destination{
			ToAddresses: receivers,
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Data: aws.String(*subject),
				},
				Body: &types.Body{
					Html: &types.Content{
						Data: aws.String(*body),
					},
				},
				Headers: []types.MessageHeader{
					{
						Name:  aws.String("X-Request-ID"),
						Value: aws.String(strconv.Itoa(reqID)),
					},
				},
			},
		},
	}
	result, err := s.Client.SendEmail(ctx, input)
	if err != nil {
		return "", err
	}
	if result.MessageId == nil {
		return "", fmt.Errorf("SES returned nil MessageId")
	}
	return *result.MessageId, nil
}
