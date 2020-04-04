package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"step-functions/internal/env"
	"step-functions/platform/emails"
)

func main() {
	sess := session.Must(session.NewSession())
	sesAPI := ses.New(sess)
	sender := emails.NewService(sesAPI, env.Get(env.SES_EMAIL_SENDER))

	handler := NewHandler(sender)

	lambda.Start(handler)
}
