package main

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"step-functions/functions/notify/mock"
	"step-functions/internal/user"
	"step-functions/platform/emails"
)

const (
	TASK_TOKEN         = "123"
	JUDGE_ENDPOINT_URL = "MOCK_ENDPOINT"
	JUDGE_EMAIL        = "JUDGE_MOCK_EMAIL"
	CANDIDATE_EMAIL    = "MOCK_EMAIL"
	CANDIDATE_ID       = "1234"
)

func TestNewHandler(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("sender failure - candidate", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sender := mock.NewMockSender(ctrl)
		endpointer := mock.NewMockEndpointer(ctrl)
		handler := NewHandler(sender, endpointer, JUDGE_ENDPOINT_URL, JUDGE_EMAIL)

		usr := user.User{
			ID:    CANDIDATE_ID,
			Email: CANDIDATE_EMAIL,
		}
		tmpl, err := emails.NewCandidateStepOneTemplate(usr.Email)
		if err != nil {
			t.Fatalf(err.Error())
		}

		sender.EXPECT().Send(ctx, tmpl, usr.Email).Return(errors.New("boom"))

		err = handler(ctx, Payload{
			User:      usr,
			TaskToken: TASK_TOKEN,
		})
		assert.Error(t, err)
	})

	t.Run("endpointer - approve endpoint failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sender := mock.NewMockSender(ctrl)
		endpointer := mock.NewMockEndpointer(ctrl)
		handler := NewHandler(sender, endpointer, JUDGE_ENDPOINT_URL, JUDGE_EMAIL)

		usr := user.User{
			ID:    CANDIDATE_ID,
			Email: CANDIDATE_EMAIL,
		}
		tmpl, err := emails.NewCandidateStepOneTemplate(usr.Email)
		if err != nil {
			t.Fatalf(err.Error())
		}

		sender.EXPECT().Send(ctx, tmpl, usr.Email).Return(nil)
		endpointer.EXPECT().NewApproveEndpoint(JUDGE_ENDPOINT_URL, TASK_TOKEN, CANDIDATE_EMAIL).Return("", errors.New("boom"))

		err = handler(ctx, Payload{
			User:      usr,
			TaskToken: TASK_TOKEN,
		})

		assert.Error(t, err)
	})

	t.Run("endpointer - deny endpoint failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sender := mock.NewMockSender(ctrl)
		endpointer := mock.NewMockEndpointer(ctrl)
		handler := NewHandler(sender, endpointer, JUDGE_ENDPOINT_URL, JUDGE_EMAIL)

		usr := user.User{
			ID:    CANDIDATE_ID,
			Email: CANDIDATE_EMAIL,
		}
		tmpl, err := emails.NewCandidateStepOneTemplate(usr.Email)
		if err != nil {
			t.Fatalf(err.Error())
		}

		sender.EXPECT().Send(ctx, tmpl, usr.Email).Return(nil)
		endpointer.EXPECT().NewApproveEndpoint(JUDGE_ENDPOINT_URL, TASK_TOKEN, CANDIDATE_EMAIL).Return("endpoint", nil)
		endpointer.EXPECT().NewDenyEndpoint(JUDGE_ENDPOINT_URL, TASK_TOKEN, CANDIDATE_EMAIL).Return("", errors.New("boom"))

		err = handler(ctx, Payload{
			User:      usr,
			TaskToken: TASK_TOKEN,
		})

		assert.Error(t, err)
	})

	t.Run("sender failure - judge", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sender := mock.NewMockSender(ctrl)
		endpointer := mock.NewMockEndpointer(ctrl)
		handler := NewHandler(sender, endpointer, JUDGE_ENDPOINT_URL, JUDGE_EMAIL)

		usr := user.User{
			ID:    CANDIDATE_ID,
			Email: CANDIDATE_EMAIL,
		}
		_, err := emails.NewJudgeDecisionTemplate(JUDGE_EMAIL, "dupa.pl", "dupa.pl")
		if err != nil {
			t.Fatalf(err.Error())
		}

		sender.EXPECT().Send(ctx, gomock.Any(), usr.Email).Return(nil)
		endpointer.EXPECT().NewApproveEndpoint(JUDGE_ENDPOINT_URL, TASK_TOKEN, CANDIDATE_EMAIL).Return("endpoint", nil)
		endpointer.EXPECT().NewDenyEndpoint(JUDGE_ENDPOINT_URL, TASK_TOKEN, CANDIDATE_EMAIL).Return("endpoint2", nil)
		sender.EXPECT().Send(ctx, gomock.Any(), JUDGE_EMAIL).Return(errors.New("boom"))

		err = handler(ctx, Payload{
			User:      usr,
			TaskToken: TASK_TOKEN,
		})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sender := mock.NewMockSender(ctrl)
		endpointer := mock.NewMockEndpointer(ctrl)
		handler := NewHandler(sender, endpointer, JUDGE_ENDPOINT_URL, JUDGE_EMAIL)

		usr := user.User{
			ID:    CANDIDATE_ID,
			Email: CANDIDATE_EMAIL,
		}

		sender.EXPECT().Send(ctx, gomock.Any(), usr.Email).Return(nil)
		endpointer.EXPECT().NewApproveEndpoint(JUDGE_ENDPOINT_URL, TASK_TOKEN, CANDIDATE_EMAIL).Return("endpoint", nil)
		endpointer.EXPECT().NewDenyEndpoint(JUDGE_ENDPOINT_URL, TASK_TOKEN, CANDIDATE_EMAIL).Return("endpoint2", nil)
		sender.EXPECT().Send(ctx, gomock.Any(), JUDGE_EMAIL).Return(nil)

		err := handler(ctx, Payload{
			User:      usr,
			TaskToken: TASK_TOKEN,
		})
		assert.NoError(t, err)
	})
}
