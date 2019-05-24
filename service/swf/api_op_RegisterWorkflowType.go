// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type RegisterWorkflowTypeInput struct {
	_ struct{} `type:"structure"`

	// If set, specifies the default policy to use for the child workflow executions
	// when a workflow execution of this type is terminated, by calling the TerminateWorkflowExecution
	// action explicitly or due to an expired timeout. This default can be overridden
	// when starting a workflow execution using the StartWorkflowExecution action
	// or the StartChildWorkflowExecution Decision.
	//
	// The supported child policies are:
	//
	//    * TERMINATE – The child executions are terminated.
	//
	//    * REQUEST_CANCEL – A request to cancel is attempted for each child execution
	//    by recording a WorkflowExecutionCancelRequested event in its history.
	//    It is up to the decider to take appropriate actions when it receives an
	//    execution history with this event.
	//
	//    * ABANDON – No action is taken. The child executions continue to run.
	DefaultChildPolicy ChildPolicy `locationName:"defaultChildPolicy" type:"string" enum:"true"`

	// If set, specifies the default maximum duration for executions of this workflow
	// type. You can override this default when starting an execution through the
	// StartWorkflowExecution Action or StartChildWorkflowExecution Decision.
	//
	// The duration is specified in seconds; an integer greater than or equal to
	// 0. Unlike some of the other timeout parameters in Amazon SWF, you cannot
	// specify a value of "NONE" for defaultExecutionStartToCloseTimeout; there
	// is a one-year max limit on the time that a workflow execution can run. Exceeding
	// this limit always causes the workflow execution to time out.
	DefaultExecutionStartToCloseTimeout *string `locationName:"defaultExecutionStartToCloseTimeout" type:"string"`

	// The default IAM role attached to this workflow type.
	//
	// Executions of this workflow type need IAM roles to invoke Lambda functions.
	// If you don't specify an IAM role when you start this workflow type, the default
	// Lambda role is attached to the execution. For more information, see http://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html
	// (http://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html)
	// in the Amazon SWF Developer Guide.
	DefaultLambdaRole *string `locationName:"defaultLambdaRole" min:"1" type:"string"`

	// If set, specifies the default task list to use for scheduling decision tasks
	// for executions of this workflow type. This default is used only if a task
	// list isn't provided when starting the execution through the StartWorkflowExecution
	// Action or StartChildWorkflowExecution Decision.
	DefaultTaskList *TaskList `locationName:"defaultTaskList" type:"structure"`

	// The default task priority to assign to the workflow type. If not assigned,
	// then 0 is used. Valid values are integers that range from Java's Integer.MIN_VALUE
	// (-2147483648) to Integer.MAX_VALUE (2147483647). Higher numbers indicate
	// higher priority.
	//
	// For more information about setting task priority, see Setting Task Priority
	// (http://docs.aws.amazon.com/amazonswf/latest/developerguide/programming-priority.html)
	// in the Amazon SWF Developer Guide.
	DefaultTaskPriority *string `locationName:"defaultTaskPriority" type:"string"`

	// If set, specifies the default maximum duration of decision tasks for this
	// workflow type. This default can be overridden when starting a workflow execution
	// using the StartWorkflowExecution action or the StartChildWorkflowExecution
	// Decision.
	//
	// The duration is specified in seconds, an integer greater than or equal to
	// 0. You can use NONE to specify unlimited duration.
	DefaultTaskStartToCloseTimeout *string `locationName:"defaultTaskStartToCloseTimeout" type:"string"`

	// Textual description of the workflow type.
	Description *string `locationName:"description" type:"string"`

	// The name of the domain in which to register the workflow type.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// The name of the workflow type.
	//
	// The specified string must not start or end with whitespace. It must not contain
	// a : (colon), / (slash), | (vertical bar), or any control characters (\u0000-\u001f
	// | \u007f-\u009f). Also, it must not contain the literal string arn.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The version of the workflow type.
	//
	// The workflow type consists of the name and version, the combination of which
	// must be unique within the domain. To get a list of all currently registered
	// workflow types, use the ListWorkflowTypes action.
	//
	// The specified string must not start or end with whitespace. It must not contain
	// a : (colon), / (slash), | (vertical bar), or any control characters (\u0000-\u001f
	// | \u007f-\u009f). Also, it must not contain the literal string arn.
	//
	// Version is a required field
	Version *string `locationName:"version" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterWorkflowTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterWorkflowTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterWorkflowTypeInput"}
	if s.DefaultLambdaRole != nil && len(*s.DefaultLambdaRole) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DefaultLambdaRole", 1))
	}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}
	if s.DefaultTaskList != nil {
		if err := s.DefaultTaskList.Validate(); err != nil {
			invalidParams.AddNested("DefaultTaskList", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterWorkflowTypeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RegisterWorkflowTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterWorkflowType = "RegisterWorkflowType"

// RegisterWorkflowTypeRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Registers a new workflow type and its configuration settings in the specified
// domain.
//
// The retention period for the workflow history is set by the RegisterDomain
// action.
//
// If the type already exists, then a TypeAlreadyExists fault is returned. You
// cannot change the configuration settings of a workflow type once it is registered
// and it must be registered as a new version.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * Constrain the following parameters by using a Condition element with
//    the appropriate keys. defaultTaskList.name: String constraint. The key
//    is swf:defaultTaskList.name. name: String constraint. The key is swf:name.
//    version: String constraint. The key is swf:version.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (http://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using RegisterWorkflowTypeRequest.
//    req := client.RegisterWorkflowTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RegisterWorkflowTypeRequest(input *RegisterWorkflowTypeInput) RegisterWorkflowTypeRequest {
	op := &aws.Operation{
		Name:       opRegisterWorkflowType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterWorkflowTypeInput{}
	}

	req := c.newRequest(op, input, &RegisterWorkflowTypeOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RegisterWorkflowTypeRequest{Request: req, Input: input, Copy: c.RegisterWorkflowTypeRequest}
}

// RegisterWorkflowTypeRequest is the request type for the
// RegisterWorkflowType API operation.
type RegisterWorkflowTypeRequest struct {
	*aws.Request
	Input *RegisterWorkflowTypeInput
	Copy  func(*RegisterWorkflowTypeInput) RegisterWorkflowTypeRequest
}

// Send marshals and sends the RegisterWorkflowType API request.
func (r RegisterWorkflowTypeRequest) Send(ctx context.Context) (*RegisterWorkflowTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterWorkflowTypeResponse{
		RegisterWorkflowTypeOutput: r.Request.Data.(*RegisterWorkflowTypeOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterWorkflowTypeResponse is the response type for the
// RegisterWorkflowType API operation.
type RegisterWorkflowTypeResponse struct {
	*RegisterWorkflowTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterWorkflowType request.
func (r *RegisterWorkflowTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
