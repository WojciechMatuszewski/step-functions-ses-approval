package stepFunctions_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"step-functions/platform/stepFunctions"
	"step-functions/platform/stepFunctions/mock"
)

const MACHINE = "MOCK_MACHINE"

func TestService_Kickoff(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		api := mock.NewMockSFNAPI(ctrl)
		service := stepFunctions.NewService(api, MACHINE)

		api.EXPECT().StartExecutionWithContext(ctx, &sfn.StartExecutionInput{
			Input:           aws.String("FOO"),
			Name:            aws.String("BAR"),
			StateMachineArn: aws.String(MACHINE),
		}).Return(nil, nil)

		err := service.Kickoff(ctx, "FOO", "BAR")
		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		api := mock.NewMockSFNAPI(ctrl)
		service := stepFunctions.NewService(api, MACHINE)
		mockErr := errors.New("boom")

		api.EXPECT().StartExecutionWithContext(ctx, &sfn.StartExecutionInput{
			Input:           aws.String("FOO"),
			Name:            aws.String("BAR"),
			StateMachineArn: aws.String(MACHINE),
		}).Return(nil, mockErr)

		err := service.Kickoff(ctx, "FOO", "BAR")
		assert.Equal(t, mockErr, err)
	})
}

func TestService_NotifySuccess(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		api := mock.NewMockSFNAPI(ctrl)
		service := stepFunctions.NewService(api, MACHINE)
		mockErr := errors.New("boom")

		api.EXPECT().SendTaskSuccessWithContext(ctx, &sfn.SendTaskSuccessInput{
			Output:    aws.String("FOO"),
			TaskToken: aws.String("BAR"),
		}).Return(nil, mockErr)

		err := service.NotifySuccess(ctx, "FOO", "BAR")
		assert.Equal(t, err, mockErr)
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		api := mock.NewMockSFNAPI(ctrl)
		service := stepFunctions.NewService(api, MACHINE)

		api.EXPECT().SendTaskSuccessWithContext(ctx, &sfn.SendTaskSuccessInput{
			Output:    aws.String("FOO"),
			TaskToken: aws.String("BAR"),
		}).Return(nil, nil)

		err := service.NotifySuccess(ctx, "FOO", "BAR")
		assert.NoError(t, err)
	})
}

func TestService_NotifyFailure(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		api := mock.NewMockSFNAPI(ctrl)
		service := stepFunctions.NewService(api, MACHINE)
		mockErr := errors.New("boom")

		api.EXPECT().SendTaskFailureWithContext(ctx, &sfn.SendTaskFailureInput{
			Cause:     aws.String("CAUSE"),
			Error:     aws.String("ERR"),
			TaskToken: aws.String("BAR"),
		}).Return(nil, mockErr)

		err := service.NotifyFailure(ctx, "CAUSE", "ERR", "BAR")
		assert.Equal(t, err, mockErr)
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		api := mock.NewMockSFNAPI(ctrl)
		service := stepFunctions.NewService(api, MACHINE)

		api.EXPECT().SendTaskFailureWithContext(ctx, &sfn.SendTaskFailureInput{
			Cause:     aws.String("CAUSE"),
			Error:     aws.String("ERR"),
			TaskToken: aws.String("BAR"),
		}).Return(nil, nil)

		err := service.NotifyFailure(ctx, "CAUSE", "ERR", "BAR")
		assert.NoError(t, err)
	})
}
