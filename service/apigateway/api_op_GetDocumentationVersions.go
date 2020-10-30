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

func (c *Client) GetDocumentationVersions(ctx context.Context, params *GetDocumentationVersionsInput, optFns ...func(*Options)) (*GetDocumentationVersionsOutput, error) {
	if params == nil {
		params = &GetDocumentationVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDocumentationVersions", params, optFns, addOperationGetDocumentationVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDocumentationVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets the documentation versions of an API.
type GetDocumentationVersionsInput struct {

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	// The current pagination position in the paged result set.
	Position *string
}

// The collection of documentation snapshots of an API. Use the
// DocumentationVersions to manage documentation snapshots associated with various
// API stages. Documenting an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// DocumentationPart, DocumentationVersion
type GetDocumentationVersionsOutput struct {

	// The current page of elements from this collection.
	Items []*types.DocumentationVersion

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDocumentationVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDocumentationVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDocumentationVersions{}, middleware.After)
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
	addOpGetDocumentationVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocumentationVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDocumentationVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetDocumentationVersions",
	}
}
