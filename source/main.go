package main

// https://github.com/aws/aws-lambda-go

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda" // go get github.com/aws/aws-lambda-go/lambda
)

// HandleRequest default handler
func HandleRequest(ctx context.Context) (string, error) {
	return "Hello world!", nil
}

func main() {
	lambda.Start(HandleRequest)
}
