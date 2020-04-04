package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
	"step-functions/internal/user"
)

// MachineStarter represents and construct which is able to kickoff a state machine.
type MachineStarter interface {
	Kickoff(ctx context.Context, input, name string) error
}

// Handler represents the shape of a handler returned by `NewHandler`.
type Handler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

// NewHandler returns a lambda handler.
func NewHandler(starter MachineStarter) Handler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var u user.User
		err := json.Unmarshal([]byte(request.Body), &u)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       http.StatusText(http.StatusBadRequest),
			}, errors.Wrap(err, "problem while unmarshalling request body")
		}

		if u.Email == "" || u.ID == "" {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       "Validation error",
			}, nil
		}

		err = starter.Kickoff(ctx, request.Body, fmt.Sprintf("%v%v", u.ID, u.Email))
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       http.StatusText(http.StatusInternalServerError),
			}, errors.Wrap(err, "problem while kicking off the machine")
		}

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       http.StatusText(http.StatusOK),
		}, nil
	}
}
