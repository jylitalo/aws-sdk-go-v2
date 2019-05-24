// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the GetBulkPublishDetails operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/GetBulkPublishDetailsRequest
type GetBulkPublishDetailsInput struct {
	_ struct{} `type:"structure"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `location:"uri" locationName:"IdentityPoolId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBulkPublishDetailsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBulkPublishDetailsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBulkPublishDetailsInput"}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBulkPublishDetailsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.IdentityPoolId != nil {
		v := *s.IdentityPoolId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IdentityPoolId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output for the GetBulkPublishDetails operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/GetBulkPublishDetailsResponse
type GetBulkPublishDetailsOutput struct {
	_ struct{} `type:"structure"`

	// If BulkPublishStatus is SUCCEEDED, the time the last bulk publish operation
	// completed.
	BulkPublishCompleteTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The date/time at which the last bulk publish was initiated.
	BulkPublishStartTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Status of the last bulk publish operation, valid values are:
	// NOT_STARTED - No bulk publish has been requested for this identity pool
	//
	// IN_PROGRESS - Data is being published to the configured stream
	//
	// SUCCEEDED - All data for the identity pool has been published to the configured
	// stream
	//
	// FAILED - Some portion of the data has failed to publish, check FailureMessage
	// for the cause.
	BulkPublishStatus BulkPublishStatus `type:"string" enum:"true"`

	// If BulkPublishStatus is FAILED this field will contain the error message
	// that caused the bulk publish to fail.
	FailureMessage *string `type:"string"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetBulkPublishDetailsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBulkPublishDetailsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BulkPublishCompleteTime != nil {
		v := *s.BulkPublishCompleteTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BulkPublishCompleteTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.BulkPublishStartTime != nil {
		v := *s.BulkPublishStartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BulkPublishStartTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if len(s.BulkPublishStatus) > 0 {
		v := s.BulkPublishStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BulkPublishStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.FailureMessage != nil {
		v := *s.FailureMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FailureMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentityPoolId != nil {
		v := *s.IdentityPoolId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IdentityPoolId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetBulkPublishDetails = "GetBulkPublishDetails"

// GetBulkPublishDetailsRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Get the status of the last BulkPublish operation for an identity pool.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
//
//    // Example sending a request using GetBulkPublishDetailsRequest.
//    req := client.GetBulkPublishDetailsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/GetBulkPublishDetails
func (c *Client) GetBulkPublishDetailsRequest(input *GetBulkPublishDetailsInput) GetBulkPublishDetailsRequest {
	op := &aws.Operation{
		Name:       opGetBulkPublishDetails,
		HTTPMethod: "POST",
		HTTPPath:   "/identitypools/{IdentityPoolId}/getBulkPublishDetails",
	}

	if input == nil {
		input = &GetBulkPublishDetailsInput{}
	}

	req := c.newRequest(op, input, &GetBulkPublishDetailsOutput{})
	return GetBulkPublishDetailsRequest{Request: req, Input: input, Copy: c.GetBulkPublishDetailsRequest}
}

// GetBulkPublishDetailsRequest is the request type for the
// GetBulkPublishDetails API operation.
type GetBulkPublishDetailsRequest struct {
	*aws.Request
	Input *GetBulkPublishDetailsInput
	Copy  func(*GetBulkPublishDetailsInput) GetBulkPublishDetailsRequest
}

// Send marshals and sends the GetBulkPublishDetails API request.
func (r GetBulkPublishDetailsRequest) Send(ctx context.Context) (*GetBulkPublishDetailsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBulkPublishDetailsResponse{
		GetBulkPublishDetailsOutput: r.Request.Data.(*GetBulkPublishDetailsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBulkPublishDetailsResponse is the response type for the
// GetBulkPublishDetails API operation.
type GetBulkPublishDetailsResponse struct {
	*GetBulkPublishDetailsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBulkPublishDetails request.
func (r *GetBulkPublishDetailsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
