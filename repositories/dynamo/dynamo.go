package dynamo

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/davidtstafford/golang-doggos/models"
)

type dynamoRepo struct {
	client dynamodbiface.DynamoDBAPI
}

var (
	dynamodbEndpoint string
	awsRegion        string
)

func NewClient() (*dynamoRepo, error) {

	loadOSEnvs()

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Endpoint:    aws.String(dynamodbEndpoint),
		Credentials: credentials.NewEnvCredentials(),
	}))

	db := dynamodb.New(sess)

	return &dynamoRepo{client: db}, nil
}

func loadOSEnvs() {
	dynamodbEndpoint = os.Getenv("dynamodbEndpoint")
	awsRegion = os.Getenv("awsRegion")
}

func (repo *dynamoRepo) GetDoggos() (*models.Doggos, error) {
	doggoList := models.Doggos{}

	scanInput := &dynamodb.ScanInput{
		AttributesToGet: aws.StringSlice([]string{`ID`, `Name`, `Breed`}),
		TableName:       aws.String(`Doggos`),
	}

	result, err := repo.client.Scan(scanInput)
	if err != nil {
		return nil, err
	}

	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &doggoList); err != nil {
		return nil, err
	}

	return &doggoList, nil
}

func (repo *dynamoRepo) WriteDoggo(doggo *models.Doggo) error {

	inputItem, _ := dynamodbattribute.MarshalMap(doggo)

	input := &dynamodb.PutItemInput{
		Item:      inputItem,
		TableName: aws.String(`Doggos`),
	}

	_, err := repo.client.PutItem(input)
	if err != nil {
		return err
	}

	repo.client.PutItem(input)

	return nil
}

func (repo *dynamoRepo) DeleteDoggo(doggo *models.Doggo) error {

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(doggo.ID),
			},
		},
		TableName: aws.String(`Doggos`),
	}

	_, err := repo.client.DeleteItem(input)
	if err != nil {
		return err
	}

	return nil
}
