package repositories

import (
	"tripmemories/server/entities"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type UsersRepository interface {
	CreateUser(user *entities.User) error
	FindUserByPk(user_pk string) (*dynamodb.QueryOutput, error)
}

func NewUsersRepository() UsersRepository {
	return &repo{
		tableName: "MainTable-DEV",
	}
}

func (*repo) CreateUser(user *entities.User) error {
	// Get a new dynamo client
	dynamoDBClient := createDynamoDBClient()

	// Transform trip into map[string]*dynamodb.AttributeValue
	attributeValue, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return err
	}

	// Create Trip item input
	item := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String("MainTable-DEV"),
	}

	// Store!
	_, err = dynamoDBClient.PutItem(item)
	if err != nil {
		return err
	}

	return nil
}

func (*repo) FindUserByPk(user_pk string) (*dynamodb.QueryOutput, error) {
	// Get a new dynamo client
	dynamoDBClient := createDynamoDBClient()

	// Create query input params
	params := &dynamodb.QueryInput{

		TableName:              aws.String("MainTable-DEV"),
		KeyConditionExpression: aws.String("partition_key = :partition_key"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":partition_key": {
				S: aws.String(user_pk),
			},
		},
	}

	// Search!
	result, err := dynamoDBClient.Query(params)

	if err != nil {
		return nil, err
	}

	return result, nil

}
