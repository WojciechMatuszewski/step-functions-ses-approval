package main

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"step-functions/platform/emails"
)

// Payload is the data that is passed to Handler
type Payload struct {
	Email string `json:"email"`
}

// Handler represents the lambda handler.
type Handler func(ctx context.Context, payload Payload) error

// Sender represents a construct which can send emails
type Sender interface {
	Send(ctx context.Context, template emails.EmailTemplate, recipient string) error
}

// NewHandler creates lambda handler
func NewHandler(sender Sender) Handler {
	l := zerolog.New(os.Stdout).With().Logger()

	return func(ctx context.Context, payload Payload) error {
		l.Info().Fields(map[string]interface{}{"payload": payload})

		tmpl, err := emails.NewApproveOutcomeTemplate(payload.Email)
		if err != nil {
			return errors.Wrap(err, "while creating approve outcome template")
		}

		err = sender.Send(ctx, tmpl, payload.Email)
		if err != nil {
			return errors.Wrap(err, "while sending")
		}

		return nil
	}
}
