// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new BasePathMapping resource.
func (c *Client) CreateBasePathMapping(ctx context.Context, params *CreateBasePathMappingInput, optFns ...func(*Options)) (*CreateBasePathMappingOutput, error) {
	if params == nil {
		params = &CreateBasePathMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBasePathMapping", params, optFns, addOperationCreateBasePathMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBasePathMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to create a new BasePathMapping resource.
type CreateBasePathMappingInput struct {

	// [Required] The domain name of the BasePathMapping resource to create.
	//
	// This member is required.
	DomainName *string

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// The base path name that callers of the API must provide as part of the URL after
	// the domain name. This value must be unique for all of the mappings across a
	// single API. Specify '(none)' if you do not want callers to specify a base path
	// name after the domain name.
	BasePath *string

	// The name of the API's stage that you want to use for this mapping. Specify
	// '(none)' if you want callers to explicitly specify the stage name after any base
	// path name.
	Stage *string
}

// Represents the base path that callers of the API must provide as part of the URL
// after the domain name. A custom domain name plus a BasePathMapping specification
// identifies a deployed RestApi in a given stage of the owner Account. Use Custom
// Domain Names
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type CreateBasePathMappingOutput struct {

	// The base path name that callers of the API must provide as part of the URL after
	// the domain name.
	BasePath *string

	// The string identifier of the associated RestApi.
	RestApiId *string

	// The name of the associated stage.
	Stage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateBasePathMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBasePathMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBasePathMapping{}, middleware.After)
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
	addOpCreateBasePathMappingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBasePathMapping(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateBasePathMapping(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateBasePathMapping",
	}
}
