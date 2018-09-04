package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Cat struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cat := Cat{Id: 1, Name: "Peter", Address: "Cali"}
	body, err := json.Marshal(&cat)
	if err != nil {
		panic(err)
	}
	response := events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}
	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
