package main

import (
	"buy-btc/bitflyer"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ticker, err := bitflyer.GetTicker(bitflyer.Btcjpy)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Bad Request!!!",
			StatusCode: 400,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Ticker:%+v", ticker),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
