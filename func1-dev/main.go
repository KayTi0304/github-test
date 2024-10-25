package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// Define a custom request structure if you need to handle JSON input
type MyRequest1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Define a custom response structure if you need to return JSON output
type MyResponse struct {
	Message string `json:"message"`
}

// Lambda function handler
func handler(ctx context.Context, request MyRequest1) (MyResponse, error) {
	// Business logic or processing based on request data
	message := fmt.Sprintf("Hello %s, you are %d years old!", request.Name, request.Age)

	// Create a response
	return MyResponse{Message: message}, nil
}

func main() {
	// Start the Lambda handler
	lambda.Start(handler)
}
