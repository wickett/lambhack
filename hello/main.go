package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	runner "github.com/karthequian/lambhack/runner"
	"log"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	log.Print("Query: ", request.QueryStringParameters["q"])
	log.Print("Headers: %v", request.Headers)
	log.Print("context ", ctx)
	headers := map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"}

	log.Print(headers)
	var output string

	if len(request.QueryStringParameters["q"]) > 0 {
		output = runner.Run(request.QueryStringParameters["q"])
		log.Print("Request %v, q=%v, %v", string(request.QueryStringParameters["q"]), string(output))
		log.Print(output)
	} else {
		body, err := json.Marshal(map[string]interface{}{
			"message": "Your function executed successfully!",
		})
		if err != nil {
			return Response{StatusCode: 404}, err
		}
		json.HTMLEscape(&buf, body)
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            output,
		Headers: map[string]string{
			"Content-Type":           "application/text",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
