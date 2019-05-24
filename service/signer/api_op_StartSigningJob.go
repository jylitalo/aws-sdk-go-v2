// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package signer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/StartSigningJobRequest
type StartSigningJobInput struct {
	_ struct{} `type:"structure"`

	// String that identifies the signing request. All calls after the first that
	// use this token return the same response as the first call.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `locationName:"clientRequestToken" type:"string" required:"true" idempotencyToken:"true"`

	// The S3 bucket in which to save your signed object. The destination contains
	// the name of your bucket and an optional prefix.
	//
	// Destination is a required field
	Destination *Destination `locationName:"destination" type:"structure" required:"true"`

	// The name of the signing profile.
	ProfileName *string `locationName:"profileName" min:"2" type:"string"`

	// The S3 bucket that contains the object to sign or a BLOB that contains your
	// raw code.
	//
	// Source is a required field
	Source *Source `locationName:"source" type:"structure" required:"true"`
}

// String returns the string representation
func (s StartSigningJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSigningJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartSigningJobInput"}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}

	if s.Destination == nil {
		invalidParams.Add(aws.NewErrParamRequired("Destination"))
	}
	if s.ProfileName != nil && len(*s.ProfileName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfileName", 2))
	}

	if s.Source == nil {
		invalidParams.Add(aws.NewErrParamRequired("Source"))
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			invalidParams.AddNested("Source", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartSigningJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	var ClientRequestToken string
	if s.ClientRequestToken != nil {
		ClientRequestToken = *s.ClientRequestToken
	} else {
		ClientRequestToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Destination != nil {
		v := s.Destination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "destination", v, metadata)
	}
	if s.ProfileName != nil {
		v := *s.ProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "profileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Source != nil {
		v := s.Source

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "source", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/StartSigningJobResponse
type StartSigningJobOutput struct {
	_ struct{} `type:"structure"`

	// The ID of your signing job.
	JobId *string `locationName:"jobId" type:"string"`
}

// String returns the string representation
func (s StartSigningJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartSigningJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opStartSigningJob = "StartSigningJob"

// StartSigningJobRequest returns a request value for making API operation for
// AWS Signer.
//
// Initiates a signing job to be performed on the code provided. Signing jobs
// are viewable by the ListSigningJobs operation for two years after they are
// performed. Note the following requirements:
//
//    * You must create an Amazon S3 source bucket. For more information, see
//    Create a Bucket (http://docs.aws.amazon.com/AmazonS3/latest/gsg/CreatingABucket.html)
//    in the Amazon S3 Getting Started Guide.
//
//    * Your S3 source bucket must be version enabled.
//
//    * You must create an S3 destination bucket. AWS Signer uses your S3 destination
//    bucket to write your signed code.
//
//    * You specify the name of the source and destination buckets when calling
//    the StartSigningJob operation.
//
//    * You must also specify a request token that identifies your request to
//    AWS Signer.
//
// You can call the DescribeSigningJob and the ListSigningJobs actions after
// you call StartSigningJob.
//
// For a Java example that shows how to use this action, see http://docs.aws.amazon.com/acm/latest/userguide/
// (http://docs.aws.amazon.com/acm/latest/userguide/)
//
//    // Example sending a request using StartSigningJobRequest.
//    req := client.StartSigningJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/StartSigningJob
func (c *Client) StartSigningJobRequest(input *StartSigningJobInput) StartSigningJobRequest {
	op := &aws.Operation{
		Name:       opStartSigningJob,
		HTTPMethod: "POST",
		HTTPPath:   "/signing-jobs",
	}

	if input == nil {
		input = &StartSigningJobInput{}
	}

	req := c.newRequest(op, input, &StartSigningJobOutput{})
	return StartSigningJobRequest{Request: req, Input: input, Copy: c.StartSigningJobRequest}
}

// StartSigningJobRequest is the request type for the
// StartSigningJob API operation.
type StartSigningJobRequest struct {
	*aws.Request
	Input *StartSigningJobInput
	Copy  func(*StartSigningJobInput) StartSigningJobRequest
}

// Send marshals and sends the StartSigningJob API request.
func (r StartSigningJobRequest) Send(ctx context.Context) (*StartSigningJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartSigningJobResponse{
		StartSigningJobOutput: r.Request.Data.(*StartSigningJobOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartSigningJobResponse is the response type for the
// StartSigningJob API operation.
type StartSigningJobResponse struct {
	*StartSigningJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartSigningJob request.
func (r *StartSigningJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
