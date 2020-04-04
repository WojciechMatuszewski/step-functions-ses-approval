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
	sfnAPI := sfn.New(sess)
	machineNotifier := stepFunctions.NewService(sfnAPI, env.Get(env.STATE_MACHINE))

	handler := NewHandler(machineNotifier)

	lambda.Start(handler)
}
