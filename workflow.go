package main

import (
	"errors"

	"go.temporal.io/sdk/workflow"
)

type MyWorkflowInput struct {
	ShouldRaiseError bool
}

type MyWorkflowOutput struct{}

func MyWorkflow(ctx workflow.Context, input MyWorkflowInput) (MyWorkflowOutput, error) {
	if input.ShouldRaiseError {
		return MyWorkflowOutput{}, errors.New("error from MyWorkflow")
	}

	childWorkflowFuture := workflow.ExecuteChildWorkflow(ctx, MyWorkflow, MyWorkflowInput{ShouldRaiseError: true})

	var childWE workflow.Execution

	errChild := childWorkflowFuture.GetChildWorkflowExecution().Get(ctx, &childWE)

	if errChild != nil {
		return MyWorkflowOutput{}, errChild
	}

	return MyWorkflowOutput{}, nil
}
