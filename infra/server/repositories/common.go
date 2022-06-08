package repositories

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
)

type repo struct {
	tableName string
}

// Create a new session using .aws/credentials
var sess *session.Session = session.Must(
	session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	),
)

const images_bucket = "trip-memories-images"

func createDynamoDBClient() *dynamodb.DynamoDB {
	return dynamodb.New(sess)
}

func createS3Client() *s3.S3 {
	return s3.New(sess)
}
