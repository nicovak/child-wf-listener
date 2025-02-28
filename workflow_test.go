package main

import (
	"log"
	"testing"

	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"

	"github.com/stretchr/testify/suite"
)

type TestMyWorkflowSuite struct {
	suite.Suite

	Env *testsuite.TestWorkflowEnvironment
	testsuite.WorkflowTestSuite
}

func (s *TestMyWorkflowSuite) SetupTest() {
	s.Env = s.NewTestWorkflowEnvironment()
}

func (s *TestMyWorkflowSuite) Test_MyWorkflow_Success() {
	s.Env.RegisterWorkflow(MyWorkflow)
	s.Env.ExecuteWorkflow(MyWorkflow, MyWorkflowInput{})

	s.Require().True(s.Env.IsWorkflowCompleted())
	s.Require().NoError(s.Env.GetWorkflowError())

	res := MyWorkflowOutput{}
	expectedRes := MyWorkflowOutput{}

	err := s.Env.GetWorkflowResult(&res)
	s.Require().NoError(err)

	s.Require().Equal(expectedRes, res)
}

func (s *TestMyWorkflowSuite) Test_MyWorkflow_Child_Error() {
	s.Env.SetOnChildWorkflowCompletedListener(func(workflowInfo *workflow.Info, result converter.EncodedValue, err error) {
		// This code never runs
		log.Panic("Child workflow should not complete")
	})

	s.Env.RegisterWorkflow(MyWorkflow)
	s.Env.ExecuteWorkflow(MyWorkflow, MyWorkflowInput{})
}

func TestTestMyWorkflowSuite(t *testing.T) {
	suite.Run(t, &TestMyWorkflowSuite{})
}
