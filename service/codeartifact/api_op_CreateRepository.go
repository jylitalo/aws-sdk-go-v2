// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a repository.
func (c *Client) CreateRepository(ctx context.Context, params *CreateRepositoryInput, optFns ...func(*Options)) (*CreateRepositoryOutput, error) {
	if params == nil {
		params = &CreateRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRepository", params, optFns, addOperationCreateRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRepositoryInput struct {

	// The domain that contains the created repository.
	//
	// This member is required.
	Domain *string

	// The name of the repository to create.
	//
	// This member is required.
	Repository *string

	// A description of the created repository.
	Description *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// One or more tag key-value pairs for the repository.
	Tags []*types.Tag

	// A list of upstream repositories to associate with the repository. The order of
	// the upstream repositories in the list determines their priority order when AWS
	// CodeArtifact looks for a requested package version. For more information, see
	// Working with upstream repositories
	// (https://docs.aws.amazon.com/codeartifact/latest/ug/repos-upstream.html).
	Upstreams []*types.UpstreamRepository
}

type CreateRepositoryOutput struct {

	// Information about the created repository after processing the request.
	Repository *types.RepositoryDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRepository{}, middleware.After)
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
	addOpCreateRepositoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRepository(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateRepository(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "CreateRepository",
	}
}
