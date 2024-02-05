package main

import 
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"


// Define a struct to hold the item data
type Item struct {
    ID   int    `json:id`
    Name string `json:name`
}

func main() {
    // Create a session with AWS
    sess, err = session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"), // Change this to your region
    })

    if err != nil {
        fmt.Println("Error creating session:", err)
        os.Exit(1)
    }

    // Create a DynamoDB client
    svc = dynamodb.New(sess)

    // Create an item to write
    item = Item{
        ID:   1,
        Name: "Bing",
    }

    // Convert the item to a map of attribute values
    av, err = dynamodbattribute.MarshalMap(item)
    if err != nil {
        fmt.Println("Error marshalling item:", err)
        os.Exit(1)
    }

    // Create an input for the PutItem operation
    input := &dynamodb.PutItemInput{
        Item:      av,
        TableName: aws.String("Items"), // Change this to your table name
    }

    // Call the PutItem method
    _, err = svc.PutItem(input)
    if err != nil {
        fmt.Println("Error putting item:", err)
        os.Exit(1)
    }

    fmt.Println("Successfully put item:", item)
}

