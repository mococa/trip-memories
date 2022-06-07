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

	// Set datatype
	attributeValue["datatype"] = &dynamodb.AttributeValue{
		S: aws.String("user"),
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
		IndexName:              aws.String("datatype-pk-index"),
		TableName:              aws.String("MainTable-DEV"),
		KeyConditionExpression: aws.String("partition_key = :partition_key AND datatype = :datatype"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":partition_key": {
				S: aws.String(user_pk),
			},
			":datatype": {
				S: aws.String("user"),
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
