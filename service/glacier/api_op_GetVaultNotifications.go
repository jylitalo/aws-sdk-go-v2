// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Provides options for retrieving the notification configuration set on an
// Amazon Glacier vault.
type GetVaultNotificationsInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetVaultNotificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVaultNotificationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetVaultNotificationsInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVaultNotificationsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Contains the Amazon Glacier response to your request.
type GetVaultNotificationsOutput struct {
	_ struct{} `type:"structure" payload:"VaultNotificationConfig"`

	// Returns the notification configuration set on the vault.
	VaultNotificationConfig *VaultNotificationConfig `locationName:"vaultNotificationConfig" type:"structure"`
}

// String returns the string representation
func (s GetVaultNotificationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVaultNotificationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VaultNotificationConfig != nil {
		v := s.VaultNotificationConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "vaultNotificationConfig", v, metadata)
	}
	return nil
}

const opGetVaultNotifications = "GetVaultNotifications"

// GetVaultNotificationsRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation retrieves the notification-configuration subresource of the
// specified vault.
//
// For information about setting a notification configuration on a vault, see
// SetVaultNotifications. If a notification configuration for a vault is not
// set, the operation returns a 404 Not Found error. For more information about
// vault notifications, see Configuring Vault Notifications in Amazon Glacier
// (http://docs.aws.amazon.com/amazonglacier/latest/dev/configuring-notifications.html).
//
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions
// by default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access Management
// (IAM) (http://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
//
// For conceptual information and underlying REST API, see Configuring Vault
// Notifications in Amazon Glacier (http://docs.aws.amazon.com/amazonglacier/latest/dev/configuring-notifications.html)
// and Get Vault Notification Configuration (http://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-get.html)
// in the Amazon Glacier Developer Guide.
//
//    // Example sending a request using GetVaultNotificationsRequest.
//    req := client.GetVaultNotificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetVaultNotificationsRequest(input *GetVaultNotificationsInput) GetVaultNotificationsRequest {
	op := &aws.Operation{
		Name:       opGetVaultNotifications,
		HTTPMethod: "GET",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/notification-configuration",
	}

	if input == nil {
		input = &GetVaultNotificationsInput{}
	}

	req := c.newRequest(op, input, &GetVaultNotificationsOutput{})
	return GetVaultNotificationsRequest{Request: req, Input: input, Copy: c.GetVaultNotificationsRequest}
}

// GetVaultNotificationsRequest is the request type for the
// GetVaultNotifications API operation.
type GetVaultNotificationsRequest struct {
	*aws.Request
	Input *GetVaultNotificationsInput
	Copy  func(*GetVaultNotificationsInput) GetVaultNotificationsRequest
}

// Send marshals and sends the GetVaultNotifications API request.
func (r GetVaultNotificationsRequest) Send(ctx context.Context) (*GetVaultNotificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetVaultNotificationsResponse{
		GetVaultNotificationsOutput: r.Request.Data.(*GetVaultNotificationsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetVaultNotificationsResponse is the response type for the
// GetVaultNotifications API operation.
type GetVaultNotificationsResponse struct {
	*GetVaultNotificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetVaultNotifications request.
func (r *GetVaultNotificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
