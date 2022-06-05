package repositories

import (
	"tripmemories/server/entities"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type TripRepository interface {
	CreateTrip(user_pk string, trip *entities.Trip) error
	ListUserTrips(user_pk string) (*dynamodb.QueryOutput, error)
	FindTrip(user_pk string, trip_sk string) (*dynamodb.QueryOutput, error)
}

func NewTripRepository() TripRepository {
	return &repo{
		tableName: "MainTable-DEV",
	}
}

func (*repo) CreateTrip(user_pk string, trip *entities.Trip) error {
	// Get a new dynamo client
	dynamoDBClient := createDynamoDBClient()

	// Add PK and SK to trip
	trip.PartitionKey = user_pk
	trip.SortKey = "trip#" + uuid.NewString()

	// Transform trip into map[string]*dynamodb.AttributeValue
	attributeValue, err := dynamodbattribute.MarshalMap(trip)
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

func (*repo) FindTrip(user_pk string, trip_sk string) (*dynamodb.QueryOutput, error) {
	// Get a new dynamo client
	dynamoDBClient := createDynamoDBClient()

	// Create query input params
	params := &dynamodb.QueryInput{

		TableName:              aws.String("MainTable-DEV"),
		KeyConditionExpression: aws.String("partition_key = :partition_key AND sort_key = :secondary_key"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":partition_key": {
				S: aws.String(user_pk),
			},
			":sort_key": {
				S: aws.String(trip_sk),
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

func (*repo) ListUserTrips(user_pk string) (*dynamodb.QueryOutput, error) {
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
