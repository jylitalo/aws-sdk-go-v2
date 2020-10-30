// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new Stage resource that references a pre-existing Deployment for the
// API.
func (c *Client) CreateStage(ctx context.Context, params *CreateStageInput, optFns ...func(*Options)) (*CreateStageOutput, error) {
	if params == nil {
		params = &CreateStageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStage", params, optFns, addOperationCreateStageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to create a Stage resource.
type CreateStageInput struct {

	// [Required] The identifier of the Deployment resource for the Stage resource.
	//
	// This member is required.
	DeploymentId *string

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// [Required] The name for the Stage resource. Stage names can only contain
	// alphanumeric characters, hyphens, and underscores. Maximum length is 128
	// characters.
	//
	// This member is required.
	StageName *string

	// Whether cache clustering is enabled for the stage.
	CacheClusterEnabled *bool

	// The stage's cache cluster size.
	CacheClusterSize types.CacheClusterSize

	// The canary deployment settings of this stage.
	CanarySettings *types.CanarySettings

	// The description of the Stage resource.
	Description *string

	// The version of the associated API documentation.
	DocumentationVersion *string

	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The
	// tag key can be up to 128 characters and must not start with aws:. The tag value
	// can be up to 256 characters.
	Tags map[string]*string

	// Specifies whether active tracing with X-ray is enabled for the Stage.
	TracingEnabled *bool

	// A map that defines the stage variables for the new Stage resource. Variable
	// names can have alphanumeric and underscore characters, and the values must match
	// [A-Za-z0-9-._~:/?#&=,]+.
	Variables map[string]*string
}

// Represents a unique identifier for a version of a deployed RestApi that is
// callable by users. Deploy an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-deploy-api.html)
type CreateStageOutput struct {

	// Settings for logging access in this stage.
	AccessLogSettings *types.AccessLogSettings

	// Specifies whether a cache cluster is enabled for the stage.
	CacheClusterEnabled *bool

	// The size of the cache cluster for the stage, if enabled.
	CacheClusterSize types.CacheClusterSize

	// The status of the cache cluster for the stage, if enabled.
	CacheClusterStatus types.CacheClusterStatus

	// Settings for the canary deployment in this stage.
	CanarySettings *types.CanarySettings

	// The identifier of a client certificate for an API stage.
	ClientCertificateId *string

	// The timestamp when the stage was created.
	CreatedDate *time.Time

	// The identifier of the Deployment that the stage points to.
	DeploymentId *string

	// The stage's description.
	Description *string

	// The version of the associated API documentation.
	DocumentationVersion *string

	// The timestamp when the stage last updated.
	LastUpdatedDate *time.Time

	// A map that defines the method settings for a Stage resource. Keys (designated as
	// /{method_setting_key below) are method paths defined as
	// {resource_path}/{http_method} for an individual method override, or /\*/\* for
	// overriding all methods in the stage.
	MethodSettings map[string]*types.MethodSetting

	// The name of the stage is the first path segment in the Uniform Resource
	// Identifier (URI) of a call to API Gateway. Stage names can only contain
	// alphanumeric characters, hyphens, and underscores. Maximum length is 128
	// characters.
	StageName *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// Specifies whether active tracing with X-ray is enabled for the Stage.
	TracingEnabled *bool

	// A map that defines the stage variables for a Stage resource. Variable names can
	// have alphanumeric and underscore characters, and the values must match
	// [A-Za-z0-9-._~:/?#&=,]+.
	Variables map[string]*string

	// The ARN of the WebAcl associated with the Stage.
	WebAclArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateStageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateStage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateStage{}, middleware.After)
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
	addOpCreateStageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateStage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateStage",
	}
}
