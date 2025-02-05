package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewS3Client creates a new AWS S3 client using the provided AWS configuration.
//
// Parameters:
//   - awsCfg: AWS configuration object containing credentials and region settings
//
// Returns:
//   - *s3.Client: A pointer to the AWS SDK S3 client instance
func NewS3Client(awsCfg aws.Config) *s3.Client {
	return s3.NewFromConfig(awsCfg)
}
