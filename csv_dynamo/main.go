package main

import (
    "context"
    // "encoding/json"
    // "log"
    // "time"
    // "github.com/aws/aws-lambda-go/events"
	"fmt"
	"os"
	"github.com/aws/aws-lambda-go/lambda"
	"./dynamo_build"
)

func HandleRequest(ctx context.Context){
	_, err := dynamo_build.GetEntries()
	if err != nil {
		fmt.Println("got error in handler")
        fmt.Println(err.Error())
        os.Exit(1)
	}
}

func main() {
	fmt.Println("started lambda")
	lambda.Start(HandleRequest)
}