// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/RegisterInstanceRequest
type RegisterInstanceInput struct {
	_ struct{} `type:"structure"`

	// The instance's hostname.
	Hostname *string `type:"string"`

	// An InstanceIdentity object that contains the instance's identity.
	InstanceIdentity *InstanceIdentity `type:"structure"`

	// The instance's private IP address.
	PrivateIp *string `type:"string"`

	// The instance's public IP address.
	PublicIp *string `type:"string"`

	// The instances public RSA key. This key is used to encrypt communication between
	// the instance and the service.
	RsaPublicKey *string `type:"string"`

	// The instances public RSA key fingerprint.
	RsaPublicKeyFingerprint *string `type:"string"`

	// The ID of the stack that the instance is to be registered with.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterInstanceInput"}

	if s.StackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a RegisterInstanceResult request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/RegisterInstanceResult
type RegisterInstanceOutput struct {
	_ struct{} `type:"structure"`

	// The registered instance's AWS OpsWorks Stacks ID.
	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s RegisterInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterInstance = "RegisterInstance"

// RegisterInstanceRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Registers instances that were created outside of AWS OpsWorks Stacks with
// a specified stack.
//
// We do not recommend using this action to register instances. The complete
// registration operation includes two tasks: installing the AWS OpsWorks Stacks
// agent on the instance, and registering the instance with the stack. RegisterInstance
// handles only the second step. You should instead use the AWS CLI register
// command, which performs the entire registration operation. For more information,
// see Registering an Instance with an AWS OpsWorks Stacks Stack (http://docs.aws.amazon.com/opsworks/latest/userguide/registered-instances-register.html).
//
// Registered instances have the same requirements as instances that are created
// by using the CreateInstance API. For example, registered instances must be
// running a supported Linux-based operating system, and they must have a supported
// instance type. For more information about requirements for instances that
// you want to register, see Preparing the Instance (http://docs.aws.amazon.com/opsworks/latest/userguide/registered-instances-register-registering-preparer.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (http://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using RegisterInstanceRequest.
//    req := client.RegisterInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/RegisterInstance
func (c *Client) RegisterInstanceRequest(input *RegisterInstanceInput) RegisterInstanceRequest {
	op := &aws.Operation{
		Name:       opRegisterInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterInstanceInput{}
	}

	req := c.newRequest(op, input, &RegisterInstanceOutput{})
	return RegisterInstanceRequest{Request: req, Input: input, Copy: c.RegisterInstanceRequest}
}

// RegisterInstanceRequest is the request type for the
// RegisterInstance API operation.
type RegisterInstanceRequest struct {
	*aws.Request
	Input *RegisterInstanceInput
	Copy  func(*RegisterInstanceInput) RegisterInstanceRequest
}

// Send marshals and sends the RegisterInstance API request.
func (r RegisterInstanceRequest) Send(ctx context.Context) (*RegisterInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterInstanceResponse{
		RegisterInstanceOutput: r.Request.Data.(*RegisterInstanceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterInstanceResponse is the response type for the
// RegisterInstance API operation.
type RegisterInstanceResponse struct {
	*RegisterInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterInstance request.
func (r *RegisterInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
