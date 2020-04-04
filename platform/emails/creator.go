package emails

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

func newEmailFromTemplate(tmpl EmailTemplate, recipient, sender string) *ses.SendEmailInput {
	return &ses.SendEmailInput{
		ConfigurationSetName: nil,
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(tmpl.body),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(tmpl.body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(tmpl.subject),
			},
		},
		Source: aws.String(sender),
	}
}
