package stepFunctions

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/sfn/sfniface"
)

// Service represents StepFunctions service
type Service struct {
	client  sfniface.SFNAPI
	machine string
}

// NewService creates a new StepFunctions service
func NewService(client sfniface.SFNAPI, machine string) Service {
	return Service{
		client:  client,
		machine: machine,
	}

}

// Kickoff starts execution of a given state machine
func (s Service) Kickoff(ctx context.Context, input, name string) error {
	_, err := s.client.StartExecutionWithContext(ctx, &sfn.StartExecutionInput{
		Input:           aws.String(input),
		Name:            aws.String(name),
		StateMachineArn: aws.String(s.machine),
	})

	return err
}

// NotifySuccess sends information to a state machine that given task succeeded
func (s Service) NotifySuccess(ctx context.Context, output, taskToken string) error {
	_, err := s.client.SendTaskSuccessWithContext(ctx, &sfn.SendTaskSuccessInput{
		Output:    aws.String(output),
		TaskToken: aws.String(taskToken),
	})

	return err
}

// NotifyFailure sends information to a state machine that given task failed
func (s Service) NotifyFailure(ctx context.Context, cause, error, taskToken string) error {
	_, err := s.client.SendTaskFailureWithContext(ctx, &sfn.SendTaskFailureInput{
		Cause:     aws.String(cause),
		Error:     aws.String(error),
		TaskToken: aws.String(taskToken),
	})

	return err
}
