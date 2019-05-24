// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CountClosedWorkflowExecutionsInput struct {
	_ struct{} `type:"structure"`

	// If specified, only workflow executions that match this close status are counted.
	// This filter has an affect only if executionStatus is specified as CLOSED.
	//
	// closeStatusFilter, executionFilter, typeFilter and tagFilter are mutually
	// exclusive. You can specify at most one of these in a request.
	CloseStatusFilter *CloseStatusFilter `locationName:"closeStatusFilter" type:"structure"`

	// If specified, only workflow executions that meet the close time criteria
	// of the filter are counted.
	//
	// startTimeFilter and closeTimeFilter are mutually exclusive. You must specify
	// one of these in a request but not both.
	CloseTimeFilter *ExecutionTimeFilter `locationName:"closeTimeFilter" type:"structure"`

	// The name of the domain containing the workflow executions to count.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// If specified, only workflow executions matching the WorkflowId in the filter
	// are counted.
	//
	// closeStatusFilter, executionFilter, typeFilter and tagFilter are mutually
	// exclusive. You can specify at most one of these in a request.
	ExecutionFilter *WorkflowExecutionFilter `locationName:"executionFilter" type:"structure"`

	// If specified, only workflow executions that meet the start time criteria
	// of the filter are counted.
	//
	// startTimeFilter and closeTimeFilter are mutually exclusive. You must specify
	// one of these in a request but not both.
	StartTimeFilter *ExecutionTimeFilter `locationName:"startTimeFilter" type:"structure"`

	// If specified, only executions that have a tag that matches the filter are
	// counted.
	//
	// closeStatusFilter, executionFilter, typeFilter and tagFilter are mutually
	// exclusive. You can specify at most one of these in a request.
	TagFilter *TagFilter `locationName:"tagFilter" type:"structure"`

	// If specified, indicates the type of the workflow executions to be counted.
	//
	// closeStatusFilter, executionFilter, typeFilter and tagFilter are mutually
	// exclusive. You can specify at most one of these in a request.
	TypeFilter *WorkflowTypeFilter `locationName:"typeFilter" type:"structure"`
}

// String returns the string representation
func (s CountClosedWorkflowExecutionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CountClosedWorkflowExecutionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CountClosedWorkflowExecutionsInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}
	if s.CloseStatusFilter != nil {
		if err := s.CloseStatusFilter.Validate(); err != nil {
			invalidParams.AddNested("CloseStatusFilter", err.(aws.ErrInvalidParams))
		}
	}
	if s.CloseTimeFilter != nil {
		if err := s.CloseTimeFilter.Validate(); err != nil {
			invalidParams.AddNested("CloseTimeFilter", err.(aws.ErrInvalidParams))
		}
	}
	if s.ExecutionFilter != nil {
		if err := s.ExecutionFilter.Validate(); err != nil {
			invalidParams.AddNested("ExecutionFilter", err.(aws.ErrInvalidParams))
		}
	}
	if s.StartTimeFilter != nil {
		if err := s.StartTimeFilter.Validate(); err != nil {
			invalidParams.AddNested("StartTimeFilter", err.(aws.ErrInvalidParams))
		}
	}
	if s.TagFilter != nil {
		if err := s.TagFilter.Validate(); err != nil {
			invalidParams.AddNested("TagFilter", err.(aws.ErrInvalidParams))
		}
	}
	if s.TypeFilter != nil {
		if err := s.TypeFilter.Validate(); err != nil {
			invalidParams.AddNested("TypeFilter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the count of workflow executions returned from CountOpenWorkflowExecutions
// or CountClosedWorkflowExecutions
type CountClosedWorkflowExecutionsOutput struct {
	_ struct{} `type:"structure"`

	// The number of workflow executions.
	//
	// Count is a required field
	Count *int64 `locationName:"count" type:"integer" required:"true"`

	// If set to true, indicates that the actual count was more than the maximum
	// supported by this API and the count returned is the truncated value.
	Truncated *bool `locationName:"truncated" type:"boolean"`
}

// String returns the string representation
func (s CountClosedWorkflowExecutionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCountClosedWorkflowExecutions = "CountClosedWorkflowExecutions"

// CountClosedWorkflowExecutionsRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Returns the number of closed workflow executions within the given domain
// that meet the specified filtering criteria.
//
// This operation is eventually consistent. The results are best effort and
// may not exactly reflect recent updates and changes.
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
//    the appropriate keys. tagFilter.tag: String constraint. The key is swf:tagFilter.tag.
//    typeFilter.name: String constraint. The key is swf:typeFilter.name. typeFilter.version:
//    String constraint. The key is swf:typeFilter.version.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (http://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using CountClosedWorkflowExecutionsRequest.
//    req := client.CountClosedWorkflowExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CountClosedWorkflowExecutionsRequest(input *CountClosedWorkflowExecutionsInput) CountClosedWorkflowExecutionsRequest {
	op := &aws.Operation{
		Name:       opCountClosedWorkflowExecutions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CountClosedWorkflowExecutionsInput{}
	}

	req := c.newRequest(op, input, &CountClosedWorkflowExecutionsOutput{})
	return CountClosedWorkflowExecutionsRequest{Request: req, Input: input, Copy: c.CountClosedWorkflowExecutionsRequest}
}

// CountClosedWorkflowExecutionsRequest is the request type for the
// CountClosedWorkflowExecutions API operation.
type CountClosedWorkflowExecutionsRequest struct {
	*aws.Request
	Input *CountClosedWorkflowExecutionsInput
	Copy  func(*CountClosedWorkflowExecutionsInput) CountClosedWorkflowExecutionsRequest
}

// Send marshals and sends the CountClosedWorkflowExecutions API request.
func (r CountClosedWorkflowExecutionsRequest) Send(ctx context.Context) (*CountClosedWorkflowExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CountClosedWorkflowExecutionsResponse{
		CountClosedWorkflowExecutionsOutput: r.Request.Data.(*CountClosedWorkflowExecutionsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CountClosedWorkflowExecutionsResponse is the response type for the
// CountClosedWorkflowExecutions API operation.
type CountClosedWorkflowExecutionsResponse struct {
	*CountClosedWorkflowExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CountClosedWorkflowExecutions request.
func (r *CountClosedWorkflowExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
