package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Name string `json:"name"`
}

func handler(event Event) (string, error) {
	if event.Name == "" {
		return "", errors.New("invalid input event")
	}
	output := fmt.Sprintf("Hello from Lambda %s", event.Name)

	return output, nil
}

func main() {
	lambda.Start(handler)
}
