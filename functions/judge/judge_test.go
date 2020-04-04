package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"step-functions/functions/judge/mock"
)

const (
	TASK_TOKEN      = "TASK_TOKEN"
	CANDIDATE_EMAIL = "MOCK_EMAIL"
)

func TestNewHandler(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("no taskToken", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		notifier := mock.NewMockMachineNotifier(ctrl)
		handler := NewHandler(notifier)

		in := events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{"foo": "bar"},
		}

		out, err := handler(ctx, in)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, out.StatusCode)
	})

	t.Run("no candidate email and notifier failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		notifier := mock.NewMockMachineNotifier(ctrl)
		handler := NewHandler(notifier)

		in := events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{"taskToken": TASK_TOKEN},
		}
		notifier.EXPECT().NotifyFailure(ctx, "no-email", gomock.Any(), TASK_TOKEN).Return(errors.New("boom"))

		out, err := handler(ctx, in)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, out.StatusCode)
	})

	t.Run("no candidate email and notifier failure success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		notifier := mock.NewMockMachineNotifier(ctrl)
		handler := NewHandler(notifier)

		in := events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{"taskToken": TASK_TOKEN},
		}
		notifier.EXPECT().NotifyFailure(ctx, "no-email", gomock.Any(), TASK_TOKEN).Return(nil)

		out, err := handler(ctx, in)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, out.StatusCode)
	})

	t.Run("no decision and notifier failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		notifier := mock.NewMockMachineNotifier(ctrl)
		handler := NewHandler(notifier)

		in := events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{"taskToken": TASK_TOKEN, "candidateEmail": CANDIDATE_EMAIL},
		}
		notifier.EXPECT().NotifyFailure(ctx, "no-decision", gomock.Any(), TASK_TOKEN).Return(errors.New("boom"))

		out, err := handler(ctx, in)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, out.StatusCode)
	})

	t.Run("no decision and notify failure success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		notifier := mock.NewMockMachineNotifier(ctrl)
		handler := NewHandler(notifier)

		in := events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{"taskToken": TASK_TOKEN, "candidateEmail": CANDIDATE_EMAIL},
		}
		notifier.EXPECT().NotifyFailure(ctx, "no-decision", gomock.Any(), TASK_TOKEN).Return(nil)

		out, err := handler(ctx, in)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, out.StatusCode)
	})

	t.Run("notify success failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		notifier := mock.NewMockMachineNotifier(ctrl)
		handler := NewHandler(notifier)

		in := events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{"taskToken": TASK_TOKEN, "decision": "approve", "candidateEmail": CANDIDATE_EMAIL},
		}
		notifier.EXPECT().NotifySuccess(ctx, fmt.Sprintf(`{"decision":"approve","email":"%v"}`, CANDIDATE_EMAIL), TASK_TOKEN).Return(errors.New("boom"))

		out, err := handler(ctx, in)
		assert.Error(t, err)
		assert.Equal(t, out.StatusCode, http.StatusInternalServerError)
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		notifier := mock.NewMockMachineNotifier(ctrl)
		handler := NewHandler(notifier)

		in := events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{"taskToken": TASK_TOKEN, "decision": "deny", "candidateEmail": CANDIDATE_EMAIL},
		}
		notifier.EXPECT().NotifySuccess(ctx, fmt.Sprintf(`{"decision":"deny","email":"%v"}`, CANDIDATE_EMAIL), TASK_TOKEN).Return(nil)

		out, err := handler(ctx, in)
		assert.NoError(t, err)
		assert.Equal(t, out.StatusCode, http.StatusOK)
	})

}
