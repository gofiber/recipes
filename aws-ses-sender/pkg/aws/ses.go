package aws

import (
	"aws-ses-sender/config"
	"context"
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
	AccessKeyId := config.GetEnv("AWS_ACCESS_KEY_ID")
	SecretAccessKey := config.GetEnv("AWS_SECRET_ACCESS_KEY")
	cfg, err := awsConfig.LoadDefaultConfig(
		ctx,
		awsConfig.WithRegion("ap-northeast-2"),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(AccessKeyId, SecretAccessKey, ""),
		),
	)
	if err != nil {
		return nil, err
	}
	return &SES{
		Client: sesv2.NewFromConfig(cfg),
	}, nil
}

// SendEmail sends an email
func (s *SES) SendEmail(ctx context.Context, reqId int, subject, body *string, receivers *[]string) (string, error) {
	sender := config.GetEnv("EMAIL_SENDER")
	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(sender),
		Destination: &types.Destination{
			ToAddresses: *receivers,
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
						Value: aws.String(strconv.Itoa(reqId)),
					},
				},
			},
		},
	}
	result, err := s.Client.SendEmail(ctx, input)
	if err != nil {
		return "", err
	}
	return *result.MessageId, nil
}
