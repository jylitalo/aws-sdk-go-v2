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

// Updates a RequestValidator of a given RestApi.
func (c *Client) UpdateRequestValidator(ctx context.Context, params *UpdateRequestValidatorInput, optFns ...func(*Options)) (*UpdateRequestValidatorOutput, error) {
	if params == nil {
		params = &UpdateRequestValidatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRequestValidator", params, optFns, addOperationUpdateRequestValidatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRequestValidatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates a RequestValidator of a given RestApi.
type UpdateRequestValidatorInput struct {

	// [Required] The identifier of RequestValidator to be updated.
	//
	// This member is required.
	RequestValidatorId *string

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []*types.PatchOperation
}

// A set of validation rules for incoming Method requests. In OpenAPI, a
// RequestValidator of an API is defined by the
// x-amazon-apigateway-request-validators.requestValidator
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions.html#api-gateway-swagger-extensions-request-validators.requestValidator.html)
// object. It the referenced using the x-amazon-apigateway-request-validator
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions.html#api-gateway-swagger-extensions-request-validator)
// property. Enable Basic Request Validation in API Gateway
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-request-validation.html)
type UpdateRequestValidatorOutput struct {

	// The identifier of this RequestValidator.
	Id *string

	// The name of this RequestValidator
	Name *string

	// A Boolean flag to indicate whether to validate a request body according to the
	// configured Model schema.
	ValidateRequestBody *bool

	// A Boolean flag to indicate whether to validate request parameters (true) or not
	// (false).
	ValidateRequestParameters *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateRequestValidatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRequestValidator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRequestValidator{}, middleware.After)
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
	addOpUpdateRequestValidatorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRequestValidator(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateRequestValidator(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateRequestValidator",
	}
}
