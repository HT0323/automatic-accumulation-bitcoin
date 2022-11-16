package main

import (
	"buy-btc/bitflyer"
	"fmt"
	"math"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	tickerChan := make(chan *bitflyer.Ticker)
	errChan := make(chan error)
	defer close(tickerChan)
	defer close(errChan)

	go bitflyer.GetTicker(tickerChan, errChan, bitflyer.Btcjpy)
	ticker := <-tickerChan
	err := <-errChan

	if err != nil {
		return getErrorResponse(err.Error()), err
	}
	//購入価格
	buyPrice := RoundDecimal(ticker.Ltp * 0.85)

	apiKey, err := getParameter("buy-btc-apikey")
	if err != nil {
		return getErrorResponse(err.Error()), nil
	}

	apiSecret, err := getParameter("buy-btc-apisecret")
	if err != nil {
		return getErrorResponse(err.Error()), nil
	}

	order := bitflyer.Order{
		ProductCode:     bitflyer.Btcjpy.String(),
		ChildOrderType:  bitflyer.Limit.String(),
		Side:            bitflyer.Buy.String(),
		Price:           buyPrice,
		Size:            0.001,
		MinuteToExpires: 4320, //3days
		TimeInForce:     bitflyer.Gtc.String(),
	}

	client := bitflyer.NewAPIClient(apiKey, apiSecret)

	orderRes, err := client.PlaceOrder(&order)
	if err != nil {
		return getErrorResponse(err.Error()), err
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("res:%+v", orderRes),
		StatusCode: 200,
	}, nil
}

func RoundDecimal(num float64) float64 {
	return math.Round(num)
}

// System Mangerからパラメータを取得
func getParameter(key string) (string, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ssm.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))

	params := &ssm.GetParameterInput{
		Name:           aws.String(key),
		WithDecryption: aws.Bool(true),
	}
	res, err := svc.GetParameter(params)
	if err != nil {
		return "", err
	}

	return *res.Parameter.Value, nil
}

func getErrorResponse(message string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       message,
		StatusCode: 400,
	}
}
func main() {
	lambda.Start(handler)
}
