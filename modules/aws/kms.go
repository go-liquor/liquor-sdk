package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

// NewKmsClient creates a new AWS KMS client using the provided AWS configuration.
//
// Parameters:
//   - awsCfg: AWS configuration object containing credentials and region settings
//
// Returns:
//   - *kms.Client: A pointer to the AWS SDK KMS client instance
func NewKmsClient(awsCfg aws.Config) *kms.Client {
	return kms.NewFromConfig(awsCfg)
}
