// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/LeaveOrganizationInput
type LeaveOrganizationInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s LeaveOrganizationInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/LeaveOrganizationOutput
type LeaveOrganizationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s LeaveOrganizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opLeaveOrganization = "LeaveOrganization"

// LeaveOrganizationRequest returns a request value for making API operation for
// AWS Organizations.
//
// Removes a member account from its parent organization. This version of the
// operation is performed by the account that wants to leave. To remove a member
// account as a user in the master account, use RemoveAccountFromOrganization
// instead.
//
// This operation can be called only from a member account in the organization.
//
//    * The master account in an organization with all features enabled can
//    set service control policies (SCPs) that can restrict what administrators
//    of member accounts can do, including preventing them from successfully
//    calling LeaveOrganization and leaving the organization.
//
//    * You can leave an organization as a member account only if the account
//    is configured with the information required to operate as a standalone
//    account. When you create an account in an organization using the AWS Organizations
//    console, API, or CLI commands, the information required of standalone
//    accounts is not automatically collected. For each account that you want
//    to make standalone, you must accept the End User License Agreement (EULA),
//    choose a support plan, provide and verify the required contact information,
//    and provide a current payment method. AWS uses the payment method to charge
//    for any billable (not free tier) AWS activity that occurs while the account
//    is not attached to an organization. Follow the steps at To leave an organization
//    when all required account information has not yet been provided (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_remove.html#leave-without-all-info)
//    in the AWS Organizations User Guide.
//
//    * You can leave an organization only after you enable IAM user access
//    to billing in your account. For more information, see Activating Access
//    to the Billing and Cost Management Console (http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/grantaccess.html#ControllingAccessWebsite-Activate)
//    in the AWS Billing and Cost Management User Guide.
//
//    // Example sending a request using LeaveOrganizationRequest.
//    req := client.LeaveOrganizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/LeaveOrganization
func (c *Client) LeaveOrganizationRequest(input *LeaveOrganizationInput) LeaveOrganizationRequest {
	op := &aws.Operation{
		Name:       opLeaveOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &LeaveOrganizationInput{}
	}

	req := c.newRequest(op, input, &LeaveOrganizationOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return LeaveOrganizationRequest{Request: req, Input: input, Copy: c.LeaveOrganizationRequest}
}

// LeaveOrganizationRequest is the request type for the
// LeaveOrganization API operation.
type LeaveOrganizationRequest struct {
	*aws.Request
	Input *LeaveOrganizationInput
	Copy  func(*LeaveOrganizationInput) LeaveOrganizationRequest
}

// Send marshals and sends the LeaveOrganization API request.
func (r LeaveOrganizationRequest) Send(ctx context.Context) (*LeaveOrganizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &LeaveOrganizationResponse{
		LeaveOrganizationOutput: r.Request.Data.(*LeaveOrganizationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// LeaveOrganizationResponse is the response type for the
// LeaveOrganization API operation.
type LeaveOrganizationResponse struct {
	*LeaveOrganizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// LeaveOrganization request.
func (r *LeaveOrganizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
