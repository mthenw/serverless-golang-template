package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mthenw/serverless-golang-template/users"
)

func main() {
	service := users.Users{}

	lambda.Start(func() (string, error) {
		return service.Create(), nil
	})
}
