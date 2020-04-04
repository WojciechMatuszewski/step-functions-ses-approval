package emails

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"step-functions/platform/emails/mock"
)

func TestService_Send(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	const recipientEmail = "cand@cand.com"
	const senderEmail = "send@send.com"

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tmpl, err := NewCandidateStepOneTemplate(recipientEmail)
		if err != nil {
			t.Fatalf(err.Error())
		}

		emailtmpl := newEmailFromTemplate(tmpl, recipientEmail, senderEmail)
		sesClient := mock.NewMockSESAPI(ctrl)
		service := NewService(sesClient, senderEmail)

		sesClient.EXPECT().SendEmailWithContext(ctx, emailtmpl).Return(nil, nil)
		err = service.Send(ctx, tmpl, recipientEmail)

		assert.NoError(t, err)
	})

	t.Run("send failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tmpl, err := NewCandidateStepOneTemplate(recipientEmail)
		if err != nil {
			t.Fatalf(err.Error())
		}

		emailtmpl := newEmailFromTemplate(tmpl, recipientEmail, senderEmail)
		sesClient := mock.NewMockSESAPI(ctrl)
		service := NewService(sesClient, senderEmail)

		sesClient.EXPECT().SendEmailWithContext(ctx, emailtmpl).Return(nil, errors.New("boom"))
		err = service.Send(ctx, tmpl, recipientEmail)

		assert.Error(t, err)
	})
}
