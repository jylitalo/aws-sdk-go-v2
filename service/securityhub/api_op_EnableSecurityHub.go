// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/EnableSecurityHubRequest
type EnableSecurityHubInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableSecurityHubInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EnableSecurityHubInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/EnableSecurityHubResponse
type EnableSecurityHubOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableSecurityHubOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EnableSecurityHubOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opEnableSecurityHub = "EnableSecurityHub"

// EnableSecurityHubRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Enables the AWS Security Hub service.
//
//    // Example sending a request using EnableSecurityHubRequest.
//    req := client.EnableSecurityHubRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/EnableSecurityHub
func (c *Client) EnableSecurityHubRequest(input *EnableSecurityHubInput) EnableSecurityHubRequest {
	op := &aws.Operation{
		Name:       opEnableSecurityHub,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts",
	}

	if input == nil {
		input = &EnableSecurityHubInput{}
	}

	req := c.newRequest(op, input, &EnableSecurityHubOutput{})
	return EnableSecurityHubRequest{Request: req, Input: input, Copy: c.EnableSecurityHubRequest}
}

// EnableSecurityHubRequest is the request type for the
// EnableSecurityHub API operation.
type EnableSecurityHubRequest struct {
	*aws.Request
	Input *EnableSecurityHubInput
	Copy  func(*EnableSecurityHubInput) EnableSecurityHubRequest
}

// Send marshals and sends the EnableSecurityHub API request.
func (r EnableSecurityHubRequest) Send(ctx context.Context) (*EnableSecurityHubResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableSecurityHubResponse{
		EnableSecurityHubOutput: r.Request.Data.(*EnableSecurityHubOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableSecurityHubResponse is the response type for the
// EnableSecurityHub API operation.
type EnableSecurityHubResponse struct {
	*EnableSecurityHubOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableSecurityHub request.
func (r *EnableSecurityHubResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}