// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/DeletePermissionRequest
type DeletePermissionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Number (ARN) of the private CA that issued the permissions.
	// You can find the CA's ARN by calling the ListCertificateAuthorities operation.
	// This must have the following form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012 .
	//
	// CertificateAuthorityArn is a required field
	CertificateAuthorityArn *string `min:"5" type:"string" required:"true"`

	// The AWS service or identity that will have its CA permissions revoked. At
	// this time, the only valid service principal is acm.amazonaws.com
	//
	// Principal is a required field
	Principal *string `type:"string" required:"true"`

	// The AWS account that calls this operation.
	SourceAccount *string `min:"12" type:"string"`
}

// String returns the string representation
func (s DeletePermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePermissionInput"}

	if s.CertificateAuthorityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if s.CertificateAuthorityArn != nil && len(*s.CertificateAuthorityArn) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateAuthorityArn", 5))
	}

	if s.Principal == nil {
		invalidParams.Add(aws.NewErrParamRequired("Principal"))
	}
	if s.SourceAccount != nil && len(*s.SourceAccount) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("SourceAccount", 12))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/DeletePermissionOutput
type DeletePermissionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePermissionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeletePermission = "DeletePermission"

// DeletePermissionRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Revokes permissions that a private CA assigned to a designated AWS service.
// Permissions can be created with the CreatePermission operation and listed
// with the ListPermissions operation.
//
//    // Example sending a request using DeletePermissionRequest.
//    req := client.DeletePermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/DeletePermission
func (c *Client) DeletePermissionRequest(input *DeletePermissionInput) DeletePermissionRequest {
	op := &aws.Operation{
		Name:       opDeletePermission,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeletePermissionInput{}
	}

	req := c.newRequest(op, input, &DeletePermissionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeletePermissionRequest{Request: req, Input: input, Copy: c.DeletePermissionRequest}
}

// DeletePermissionRequest is the request type for the
// DeletePermission API operation.
type DeletePermissionRequest struct {
	*aws.Request
	Input *DeletePermissionInput
	Copy  func(*DeletePermissionInput) DeletePermissionRequest
}

// Send marshals and sends the DeletePermission API request.
func (r DeletePermissionRequest) Send(ctx context.Context) (*DeletePermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePermissionResponse{
		DeletePermissionOutput: r.Request.Data.(*DeletePermissionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePermissionResponse is the response type for the
// DeletePermission API operation.
type DeletePermissionResponse struct {
	*DeletePermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePermission request.
func (r *DeletePermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
