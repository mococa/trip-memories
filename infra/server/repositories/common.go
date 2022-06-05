package repositories

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type repo struct {
	tableName string
}

func createDynamoDBClient() *dynamodb.DynamoDB {
	// Create a new session using .aws/credentials
	sess := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState: session.SharedConfigEnable,
			},
		),
	)

	return dynamodb.New(sess)
}
