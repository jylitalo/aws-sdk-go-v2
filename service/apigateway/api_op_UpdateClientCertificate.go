// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to change information about an ClientCertificate resource.
type UpdateClientCertificateInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the ClientCertificate resource to be updated.
	//
	// ClientCertificateId is a required field
	ClientCertificateId *string `location:"uri" locationName:"clientcertificate_id" type:"string" required:"true"`

	// A list of update operations to be applied to the specified resource and in
	// the order specified in this list.
	PatchOperations []PatchOperation `locationName:"patchOperations" type:"list"`
}

// String returns the string representation
func (s UpdateClientCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateClientCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateClientCertificateInput"}

	if s.ClientCertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientCertificateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateClientCertificateInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.PatchOperations) > 0 {
		v := s.PatchOperations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "patchOperations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ClientCertificateId != nil {
		v := *s.ClientCertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "clientcertificate_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents a client certificate used to configure client-side SSL authentication
// while sending requests to the integration endpoint.
//
// Client certificates are used to authenticate an API by the backend server.
// To authenticate an API client (or user), use IAM roles and policies, a custom
// Authorizer or an Amazon Cognito user pool.
//
// Use Client-Side Certificate (https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-client-side-ssl-authentication.html)
type UpdateClientCertificateOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the client certificate.
	ClientCertificateId *string `locationName:"clientCertificateId" type:"string"`

	// The timestamp when the client certificate was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix"`

	// The description of the client certificate.
	Description *string `locationName:"description" type:"string"`

	// The timestamp when the client certificate will expire.
	ExpirationDate *time.Time `locationName:"expirationDate" type:"timestamp" timestampFormat:"unix"`

	// The PEM-encoded public key of the client certificate, which can be used to
	// configure certificate authentication in the integration endpoint .
	PemEncodedCertificate *string `locationName:"pemEncodedCertificate" type:"string"`
}

// String returns the string representation
func (s UpdateClientCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateClientCertificateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientCertificateId != nil {
		v := *s.ClientCertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientCertificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExpirationDate != nil {
		v := *s.ExpirationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expirationDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.PemEncodedCertificate != nil {
		v := *s.PemEncodedCertificate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "pemEncodedCertificate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateClientCertificate = "UpdateClientCertificate"

// UpdateClientCertificateRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Changes information about an ClientCertificate resource.
//
//    // Example sending a request using UpdateClientCertificateRequest.
//    req := client.UpdateClientCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateClientCertificateRequest(input *UpdateClientCertificateInput) UpdateClientCertificateRequest {
	op := &aws.Operation{
		Name:       opUpdateClientCertificate,
		HTTPMethod: "PATCH",
		HTTPPath:   "/clientcertificates/{clientcertificate_id}",
	}

	if input == nil {
		input = &UpdateClientCertificateInput{}
	}

	req := c.newRequest(op, input, &UpdateClientCertificateOutput{})
	return UpdateClientCertificateRequest{Request: req, Input: input, Copy: c.UpdateClientCertificateRequest}
}

// UpdateClientCertificateRequest is the request type for the
// UpdateClientCertificate API operation.
type UpdateClientCertificateRequest struct {
	*aws.Request
	Input *UpdateClientCertificateInput
	Copy  func(*UpdateClientCertificateInput) UpdateClientCertificateRequest
}

// Send marshals and sends the UpdateClientCertificate API request.
func (r UpdateClientCertificateRequest) Send(ctx context.Context) (*UpdateClientCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateClientCertificateResponse{
		UpdateClientCertificateOutput: r.Request.Data.(*UpdateClientCertificateOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateClientCertificateResponse is the response type for the
// UpdateClientCertificate API operation.
type UpdateClientCertificateResponse struct {
	*UpdateClientCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateClientCertificate request.
func (r *UpdateClientCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
