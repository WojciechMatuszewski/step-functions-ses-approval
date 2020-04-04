package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"step-functions/internal/user"
	"step-functions/platform/emails"
)

// Sender represents a construct which can send emails
type Sender interface {
	Send(ctx context.Context, template emails.EmailTemplate, recipient string) error
}

// Endpointer is responsible for creating endpoints (wrong name change it)
type Endpointer interface {
	NewDenyEndpoint(rootURL, token, candidateEmail string) (string, error)
	NewApproveEndpoint(rootURL, token, candidateEmail string) (string, error)
}

// Handler represents the shape of a handler returned by `NewHandler`.
type Handler func(ctx context.Context, p Payload) error

// Payload represents the payload given to lambda by state machine
type Payload struct {
	User      user.User `json:"user"`
	TaskToken string    `json:"taskToken"`
}

// NewHandler creates a new lambda Handler
func NewHandler(sender Sender, endpointer Endpointer, judgeEndpointURL, judgeEmail string) Handler {

	l := zerolog.New(os.Stdout).With().Logger()
	return func(ctx context.Context, payload Payload) error {
		l.Info().Fields(map[string]interface{}{"payload": payload}).Msg("notify started")

		l.Info().Msg("sending to candidate")
		ctmpl, err := emails.NewCandidateStepOneTemplate(payload.User.Email)
		if err != nil {
			return errors.Wrap(err, "while generating candidate template")
		}

		err = sender.Send(ctx, ctmpl, payload.User.Email)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("problem while sending to %v", payload.User.Email))
		}

		approveEndpoint, err := endpointer.NewApproveEndpoint(judgeEndpointURL, payload.TaskToken, payload.User.Email)
		if err != nil {
			return errors.Wrap(err, "endpointer - new approve endpoint")
		}

		denyEndpoint, err := endpointer.NewDenyEndpoint(judgeEndpointURL, payload.TaskToken, payload.User.Email)
		if err != nil {
			return errors.Wrap(err, "endpointer - new deny endpoint")
		}

		jtmpl, err := emails.NewJudgeDecisionTemplate(judgeEmail, approveEndpoint, denyEndpoint)
		if err != nil {
			return errors.Wrap(err, "while generating judge template")
		}

		err = sender.Send(ctx, jtmpl, judgeEmail)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("problem while sending to %v", judgeEmail))
		}

		l.Info().Msg("notify ended")

		return nil
	}
}
