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

// Represents a collection of BasePathMapping resources.
func (c *Client) GetBasePathMappings(ctx context.Context, params *GetBasePathMappingsInput, optFns ...func(*Options)) (*GetBasePathMappingsOutput, error) {
	if params == nil {
		params = &GetBasePathMappingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBasePathMappings", params, optFns, addOperationGetBasePathMappingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBasePathMappingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get information about a collection of BasePathMapping resources.
type GetBasePathMappingsInput struct {

	// [Required] The domain name of a BasePathMapping resource.
	//
	// This member is required.
	DomainName *string

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	// The current pagination position in the paged result set.
	Position *string
}

// Represents a collection of BasePathMapping resources. Use Custom Domain Names
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type GetBasePathMappingsOutput struct {

	// The current page of elements from this collection.
	Items []*types.BasePathMapping

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBasePathMappingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBasePathMappings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBasePathMappings{}, middleware.After)
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
	addOpGetBasePathMappingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBasePathMappings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBasePathMappings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetBasePathMappings",
	}
}
