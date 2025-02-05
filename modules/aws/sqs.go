package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// NewSQSClient creates a new AWS SQS client using the provided AWS configuration.
//
// Parameters:
//   - awsCfg: AWS configuration object containing credentials and region settings
//
// Returns:
//   - *sqs.Client: A pointer to the AWS SDK SQS client instance
func NewSQSClient(awsCfg aws.Config) *sqs.Client {
	return sqs.NewFromConfig(awsCfg)
}
