package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// NewDynamoDBClient creates a new AWS DynamoDB client using the provided AWS configuration.
//
// Parameters:
//   - awsCfg: AWS configuration object containing credentials and region settings
//
// Returns:
//   - *dynamodb.Client: A pointer to the AWS SDK DynamoDB client instance
func NewDynamoDBClient(awsCfg aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(awsCfg)
}
