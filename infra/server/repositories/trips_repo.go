package repositories

import (
	"mime/multipart"
	"tripmemories/server/entities"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/s3"
)

type TripRepository interface {
	CreateTrip(user_pk string, trip *entities.Trip) error
	ListUserTrips(user_pk string) (*dynamodb.QueryOutput, error)
	FindTrip(user_pk string, trip_sk string) (*dynamodb.QueryOutput, error)
	UploadTripImage(f multipart.File, key string) (*s3.PutObjectOutput, error)
	AddTripBanner(userpk string, trip_sk string, picture string) error
	FindTripImages(user_pk string, trip_sk string) (*s3.ListObjectsV2Output, error)
}

func NewTripRepository() TripRepository {
	return &repo{
		tableName: "MainTable-DEV",
	}
}

func (*repo) CreateTrip(user_pk string, trip *entities.Trip) error {
	// Get a new dynamo client
	dynamoDBClient := createDynamoDBClient()

	// Transform trip into map[string]*dynamodb.AttributeValue
	attributeValue, err := dynamodbattribute.MarshalMap(trip)
	if err != nil {
		return err
	}

	// Set datatype
	attributeValue["datatype"] = &dynamodb.AttributeValue{
		S: aws.String("trip"),
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
		KeyConditionExpression: aws.String("partition_key = :partition_key AND sort_key = :sort_key"),
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
		IndexName:              aws.String("datatype-pk-index"),
		KeyConditionExpression: aws.String("partition_key = :partition_key AND datatype = :datatype"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":partition_key": {
				S: aws.String(user_pk),
			},
			":datatype": {
				S: aws.String("trip"),
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

func (*repo) UploadTripImage(f multipart.File, key string) (*s3.PutObjectOutput, error) {
	s3Client := createS3Client()
	resp, err := s3Client.PutObject(&s3.PutObjectInput{
		Body:   f,
		Bucket: aws.String(images_bucket),
		Key:    aws.String(key),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (*repo) FindTripImages(user_pk string, trip_sk string) (*s3.ListObjectsV2Output, error) {
	s3Client := createS3Client()
	res, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(images_bucket),
		Prefix: aws.String(user_pk + "/" + trip_sk),
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (*repo) AddTripBanner(user_pk string, trip_sk string, picture string) error {
	// Get a new dynamo client
	dynamoDBClient := createDynamoDBClient()

	// Set a trip banner
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("MainTable-DEV"),
		Key: map[string]*dynamodb.AttributeValue{
			"partition_key": {
				S: aws.String(user_pk),
			},
			"sort_key": {
				S: aws.String(trip_sk),
			},
		},
		UpdateExpression: aws.String("set #picture = :picture"),
		ExpressionAttributeNames: map[string]*string{
			"#picture": aws.String("picture"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":picture": {
				S: aws.String(picture),
			},
		},
	}

	// Update!
	_, err := dynamoDBClient.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}
