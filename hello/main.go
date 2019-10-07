package main

import (
	"context"
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

	output := "Your function executed successfully!"
	if len(request.QueryStringParameters["q"]) > 0 {
		// Source of our hacky code...
		output = runner.Run(request.QueryStringParameters["q"])
		log.Print("Request %v, q=%v, %v", string(request.QueryStringParameters["q"]), string(output))
		log.Print(output)
	}

	resp := Response{
		StatusCode: 200,
		Body:       output,
		Headers: map[string]string{
			"Content-Type": "application/text",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
