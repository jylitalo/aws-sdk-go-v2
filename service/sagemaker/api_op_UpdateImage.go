// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the properties of a SageMaker image. To change the image's tags, use the
// AddTags and DeleteTags APIs.
func (c *Client) UpdateImage(ctx context.Context, params *UpdateImageInput, optFns ...func(*Options)) (*UpdateImageOutput, error) {
	if params == nil {
		params = &UpdateImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateImage", params, optFns, addOperationUpdateImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateImageInput struct {

	// The name of the image to update.
	//
	// This member is required.
	ImageName *string

	// A list of properties to delete. Only the Description and DisplayName properties
	// can be deleted.
	DeleteProperties []*string

	// The new description for the image.
	Description *string

	// The new display name for the image.
	DisplayName *string

	// The new Amazon Resource Name (ARN) for the IAM role that enables Amazon
	// SageMaker to perform tasks on your behalf.
	RoleArn *string
}

type UpdateImageOutput struct {

	// The Amazon Resource Name (ARN) of the image.
	ImageArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateImage{}, middleware.After)
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
	addOpUpdateImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateImage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateImage",
	}
}
