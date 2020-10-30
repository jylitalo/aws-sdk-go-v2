// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the VpcLinks collection under the caller's account in a selected region.
func (c *Client) GetVpcLinks(ctx context.Context, params *GetVpcLinksInput, optFns ...func(*Options)) (*GetVpcLinksOutput, error) {
	if params == nil {
		params = &GetVpcLinksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVpcLinks", params, optFns, addOperationGetVpcLinksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVpcLinksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets the VpcLinks collection under the caller's account in a selected region.
type GetVpcLinksInput struct {

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	// The current pagination position in the paged result set.
	Position *string
}

// The collection of VPC links under the caller's account in a region. Getting
// Started with Private Integrations
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-with-private-integration.html),
// Set up Private Integrations
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-private-integration.html)
type GetVpcLinksOutput struct {

	// The current page of elements from this collection.
	Items []*types.VpcLink

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetVpcLinksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetVpcLinks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVpcLinks{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetVpcLinks(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetVpcLinks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetVpcLinks",
	}
}
