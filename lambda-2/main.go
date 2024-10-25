package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// Define a custom request structure if you need to handle JSON input
type MyRequest2 struct {
	Name   string `json:"name"`
	Animal string `json:"animal"`
}

// Define a custom response structure if you need to return JSON output
type MyResponse2 struct {
	Message string `json:"message"`
}

// Lambda function handler
func handler(ctx context.Context, request MyRequest2) (MyResponse2, error) {
	// Business logic or processing based on request data
	fmt.Println("in lambda 2")
	message := fmt.Sprintf("Hello %s, you are %d years old!", request.Name, request.Animal)

	// Create a response
	return MyResponse2{Message: message}, nil
}

func main() {
	// Start the Lambda handler
	lambda.Start(handler)
}
