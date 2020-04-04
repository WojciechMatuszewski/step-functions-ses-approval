package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
	"step-functions/internal/env"
	"step-functions/platform/stepFunctions"
)

func main() {
	sess := session.Must(session.NewSession())
	sfnClient := sfn.New(sess)
	sfnService := stepFunctions.NewService(sfnClient, env.Get(env.STATE_MACHINE))

	handler := NewHandler(sfnService)
	lambda.Start(handler)
}
