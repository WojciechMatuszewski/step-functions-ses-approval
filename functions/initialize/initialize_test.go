package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"step-functions/functions/initialize/mock"
)

func TestNewHandler(t *testing.T) {
	ctx := context.Background()
	t.Parallel()

	t.Run("validation - email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		starter := mock.NewMockMachineStarter(ctrl)
		handler := NewHandler(starter)

		in := events.APIGatewayProxyRequest{
			Body: `{"id": "123"}`,
		}
		resp, err := handler(ctx, in)

		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
		assert.Equal(t, resp.Body, "Validation error")
	})

	t.Run("validation - id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		starter := mock.NewMockMachineStarter(ctrl)
		handler := NewHandler(starter)

		in := events.APIGatewayProxyRequest{
			Body: `{"email": "email@email.com"}`,
		}
		resp, err := handler(ctx, in)

		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
		assert.Equal(t, resp.Body, "Validation error")
	})

	t.Run("starter error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		starter := mock.NewMockMachineStarter(ctrl)
		handler := NewHandler(starter)

		starter.EXPECT().
			Kickoff(ctx, `{"email": "email@email.com", "id": "1234"}`, "1234email@email.com").
			Return(errors.New("boom"))
		in := events.APIGatewayProxyRequest{
			Body: `{"email": "email@email.com", "id": "1234"}`,
		}
		resp, err := handler(ctx, in)

		assert.Error(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusInternalServerError)
		assert.Equal(t, resp.Body, http.StatusText(http.StatusInternalServerError))
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		starter := mock.NewMockMachineStarter(ctrl)
		handler := NewHandler(starter)

		starter.EXPECT().
			Kickoff(ctx, `{"email": "email@email.com", "id": "1234"}`, "1234email@email.com").
			Return(nil)
		in := events.APIGatewayProxyRequest{
			Body: `{"email": "email@email.com", "id": "1234"}`,
		}
		resp, err := handler(ctx, in)

		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusOK)
		assert.Equal(t, resp.Body, http.StatusText(http.StatusOK))
	})
}
