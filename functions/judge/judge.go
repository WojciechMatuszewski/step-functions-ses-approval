package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Handler is the handler returned by `NewHandler`
type Handler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

// MachineNotifier can send success or failure messages to a state machine
type MachineNotifier interface {
	NotifySuccess(ctx context.Context, output, taskToken string) error
	NotifyFailure(ctx context.Context, cause, outError, taskToken string) error
}

// MachineOutput is the output payload for MachineNotifier
type MachineOutput struct {
	Decision string `json:"decision"`
	Email    string `json:"email"`
}

// NewHandler returns a lambda Handler
func NewHandler(notifier MachineNotifier) Handler {
	l := zerolog.New(os.Stdout).With().Logger()

	return func(ctx context.Context, request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
		qParams := request.QueryStringParameters
		l.Info().Fields(map[string]interface{}{"queryParams": qParams}).Msg("params")

		taskToken, err := findInQueryParams(qParams, "taskToken")
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       http.StatusText(http.StatusBadRequest),
			}, errors.Wrap(err, "while getting from a query params")
		}
		l.Info().Fields(map[string]interface{}{"taskToken": taskToken})

		candidateEmail, err := findInQueryParams(qParams, "candidateEmail")
		if err != nil {
			notifierErr := notifier.NotifyFailure(ctx, "no-email", err.Error(), taskToken)
			if notifierErr != nil {
				return events.APIGatewayProxyResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       http.StatusText(http.StatusInternalServerError),
				}, errors.Wrap(notifierErr, errors.Errorf("while sending task failure: %v", err).Error())
			}

			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       http.StatusText(http.StatusBadRequest),
			}, errors.Wrap(err, "while getting from a query params")
		}

		decision, err := findInQueryParams(qParams, "decision")
		if err != nil {
			notifierErr := notifier.NotifyFailure(ctx, "no-decision", err.Error(), taskToken)
			if notifierErr != nil {
				return events.APIGatewayProxyResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       http.StatusText(http.StatusInternalServerError),
				}, errors.Wrap(notifierErr, errors.Errorf("while sending task failure: %v", err).Error())
			}

			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       http.StatusText(http.StatusBadRequest),
			}, errors.Wrap(err, "while getting from a query params")
		}

		out := MachineOutput{
			Decision: decision,
			Email:    candidateEmail,
		}
		outB, err := json.Marshal(&out)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       http.StatusText(http.StatusInternalServerError),
			}, errors.Wrap(err, "while marshalling output")
		}

		err = notifier.NotifySuccess(ctx, string(outB), taskToken)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       http.StatusText(http.StatusInternalServerError),
			}, errors.Wrap(err, "while sending task success")
		}

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       decision,
		}, nil

	}
}

func findInQueryParams(params map[string]string, key string) (string, error) {
	v, found := params[key]
	if !found {
		return "", fmt.Errorf("key %v not found", key)
	}

	parsed, err := url.PathUnescape(v)
	if err != nil {
		return "", nil
	}

	return parsed, nil
}
