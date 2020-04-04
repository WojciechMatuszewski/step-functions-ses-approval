package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/service/ses"
	"step-functions/internal/endpoint"
	"step-functions/internal/env"
	"step-functions/platform/emails"
)

func main() {
	sess := session.Must(session.NewSession())

	sesAPI := ses.New(sess)
	sender := emails.NewService(sesAPI, env.Get(env.SES_EMAIL_SENDER))

	signer := v4.NewSigner(sess.Config.Credentials)
	endpointer := endpoint.NewService(signer)

	handler := NewHandler(sender, endpointer, env.Get(env.JUDGE_ENDPOINT_URL), env.Get(env.SES_EMAIL_SENDER))
	lambda.Start(handler)
}
