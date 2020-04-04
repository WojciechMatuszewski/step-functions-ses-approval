package emails

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

// Service represents Emails service
type Service struct {
	client sesiface.SESAPI
	sender string
}

// NewService creates new email service
func NewService(client sesiface.SESAPI, sender string) Service {
	return Service{
		client: client,
		sender: sender,
	}
}

// Send sends email of a given type to a given recipient
func (s Service) Send(ctx context.Context, template EmailTemplate, recipient string) error {
	email := newEmailFromTemplate(template, recipient, s.sender)

	_, err := s.client.SendEmailWithContext(ctx, email)
	return err
}
