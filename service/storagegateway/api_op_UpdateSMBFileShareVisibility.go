// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Controls whether the shares on a gateway are visible in a net view or browse
// list.
func (c *Client) UpdateSMBFileShareVisibility(ctx context.Context, params *UpdateSMBFileShareVisibilityInput, optFns ...func(*Options)) (*UpdateSMBFileShareVisibilityOutput, error) {
	if params == nil {
		params = &UpdateSMBFileShareVisibilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSMBFileShareVisibility", params, optFns, addOperationUpdateSMBFileShareVisibilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSMBFileShareVisibilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSMBFileShareVisibilityInput struct {

	// The shares on this gateway appear when listing shares.
	//
	// This member is required.
	FileSharesVisible *bool

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

type UpdateSMBFileShareVisibilityOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateSMBFileShareVisibilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSMBFileShareVisibility{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSMBFileShareVisibility{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateSMBFileShareVisibilityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSMBFileShareVisibility(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateSMBFileShareVisibility(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateSMBFileShareVisibility",
	}
}
