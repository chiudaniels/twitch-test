package model

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Entry struct {
	Year  int
	Title string
}

func MakeTable() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		// Credentials: credentials.NewSharedCredentials("", "default"),
	})

	svc := dynamodb.New(sess)

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Year"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Title"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Year"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Title"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Movies"),
	}

	result, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}

func AddEntry(yearMade int, titleName string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		// Credentials: credentials.NewSharedCredentials("", "default"),
	})

	svc := dynamodb.New(sess)

	item := Entry{
		Year:  yearMade,
		Title: titleName,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tableName := "Movies"
	fmt.Println(av)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	years := strconv.Itoa(item.Year)

	fmt.Println("Successfully added '" + item.Title + "' (" + years + ") to table " + tableName)
}

func GetEntry(titleName string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		// Credentials: credentials.NewSharedCredentials("", "default"),
	})

	svc := dynamodb.New(sess)

	// Title := titleName,
	tableName := "Movies"

	filt := expression.Name("Title").Equal(expression.Value(titleName))
	proj := expression.NamesList(expression.Name("Title"), expression.Name("Year"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	result, err := svc.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	numItems := 0

	for _, i := range result.Items {
		item := Entry{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Which ones had a higher rating than minimum?
		if item.Title == titleName {
			// Or it we had filtered by rating previously:
			//   if item.Year == year {
			numItems++

			fmt.Println("Title: ", item.Title)
			fmt.Println("Year:", item.Year)
			fmt.Println()
		}
	}

	fmt.Println("Found", numItems, "movie(s) with a name", titleName)
}

func GetAllEntry() []Entry {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		// Credentials: credentials.NewSharedCredentials("", "default"),
	})

	svc := dynamodb.New(sess)

	params := &dynamodb.ScanInput{
		TableName: aws.String("Movies"),
	}
	result, err := svc.Scan(params)
	if err != nil {
		fmt.Errorf("failed to make Query API call, %v", err)
	}
	obj := []Entry{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &obj)
	if err != nil {
		fmt.Errorf("failed to unmarshal Query result items, %v", err)
	}
	// Title := titleName,
	fmt.Println("Got all entries")
	return obj
}

// func main() {
// 	// MakeTable()
// 	// AddEntry(2015, "Sonic")
// 	GetEntry("Sonic")
// }
