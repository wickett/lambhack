package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	sparta "github.com/mweagle/Sparta"
	runner "github.com/wickett/lambhack/runner"
)

////////////////////////////////////////////////////////////////////////////////
func paramVal(keyName string, defaultValue string) string {
	value := os.Getenv(keyName)
	if "" == value {
		value = defaultValue
	}
	return value
}

var s3Bucket = paramVal("S3_TEST_BUCKET", "arn:aws:s3:::lambhack")

func lambhackEvent(event *json.RawMessage,
	context *sparta.LambdaContext,
	w http.ResponseWriter,
	logger *logrus.Logger) {

	//code here for lambhack

	var lambdaEvent sparta.APIGatewayLambdaJSONEvent
	_ = json.Unmarshal([]byte(*event), &lambdaEvent)

	command := lambdaEvent.QueryParams["args"]
	output := runner.Run(command)
	logger.WithFields(logrus.Fields{
		"Event":   string(*event),
		"Command": string(command),
		"Output":  string(output),
	}).Info("Request received")

	fmt.Fprintf(w, output)
	time.Sleep(time.Second)
}

func appendS3Lambda(api *sparta.API, lambdaFunctions []*sparta.LambdaAWSInfo) []*sparta.LambdaAWSInfo {
	options := new(sparta.LambdaFunctionOptions)
	options.Timeout = 30
	lambdaFn := sparta.NewLambda(sparta.IAMRoleDefinition{}, lambhackEvent, options)
	apiGatewayResource, _ := api.NewResource("/lambhack/{command+}", lambdaFn)
	apiGatewayResource.NewMethod("GET", http.StatusOK)

	lambdaFn.Permissions = append(lambdaFn.Permissions, sparta.S3Permission{
		BasePermission: sparta.BasePermission{
			SourceArn: s3Bucket,
		},
		Events: []string{"s3:ObjectCreated:*", "s3:ObjectRemoved:*"},
	})
	return append(lambdaFunctions, lambdaFn)
}

////////////////////////////////////////////////////////////////////////////////
// Return the *[]sparta.LambdaAWSInfo slice
//
func spartaLambdaData(api *sparta.API) []*sparta.LambdaAWSInfo {

	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFunctions = appendS3Lambda(api, lambdaFunctions)
	return lambdaFunctions
}

func main() {
	stage := sparta.NewStage("prod")
	apiGateway := sparta.NewAPIGateway("lambhackAPI", stage)
	apiGateway.CORSEnabled = true

	//lambda info
	os.Setenv("AWS_PROFILE", "sparta")
	os.Setenv("AWS_REGION", "us-east-1")

	stackName := "LambhackApplication"
	sparta.Main(stackName,
		"Lambhack Application",
		spartaLambdaData(apiGateway),
		apiGateway,
		nil)

}
