// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates supported fields of the specified job.
func (c *Client) UpdateJob(ctx context.Context, params *UpdateJobInput, optFns ...func(*Options)) (*UpdateJobOutput, error) {
	if params == nil {
		params = &UpdateJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateJob", params, optFns, addOperationUpdateJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateJobInput struct {

	// The ID of the job to be updated.
	//
	// This member is required.
	JobId *string

	// Allows you to create criteria to abort a job.
	AbortConfig *types.AbortConfig

	// A short text description of the job.
	Description *string

	// Allows you to create a staged rollout of the job.
	JobExecutionsRolloutConfig *types.JobExecutionsRolloutConfig

	// The namespace used to indicate that a job is a customer-managed job. When you
	// specify a value for this parameter, AWS IoT Core sends jobs notifications to
	// MQTT topics that contain the value in the following format.
	// $aws/things/THING_NAME/jobs/JOB_ID/notify-namespace-NAMESPACE_ID/ The
	// namespaceId feature is in public preview.
	NamespaceId *string

	// Configuration information for pre-signed S3 URLs.
	PresignedUrlConfig *types.PresignedUrlConfig

	// Specifies the amount of time each device has to finish its execution of the job.
	// The timer is started when the job execution status is set to IN_PROGRESS. If the
	// job execution status is not set to another terminal state before the time
	// expires, it will be automatically set to TIMED_OUT.
	TimeoutConfig *types.TimeoutConfig
}

type UpdateJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateJob{}, middleware.After)
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
	addOpUpdateJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateJob",
	}
}
