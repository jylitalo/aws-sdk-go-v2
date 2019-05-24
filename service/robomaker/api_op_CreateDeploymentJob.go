// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateDeploymentJobRequest
type CreateDeploymentJobInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The deployment application configuration.
	//
	// DeploymentApplicationConfigs is a required field
	DeploymentApplicationConfigs []DeploymentApplicationConfig `locationName:"deploymentApplicationConfigs" min:"1" type:"list" required:"true"`

	// The requested deployment configuration.
	DeploymentConfig *DeploymentConfig `locationName:"deploymentConfig" type:"structure"`

	// The Amazon Resource Name (ARN) of the fleet to deploy.
	//
	// Fleet is a required field
	Fleet *string `locationName:"fleet" min:"1" type:"string" required:"true"`

	// A map that contains tag keys and tag values that are attached to the deployment
	// job.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateDeploymentJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeploymentJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeploymentJobInput"}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.DeploymentApplicationConfigs == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentApplicationConfigs"))
	}
	if s.DeploymentApplicationConfigs != nil && len(s.DeploymentApplicationConfigs) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeploymentApplicationConfigs", 1))
	}

	if s.Fleet == nil {
		invalidParams.Add(aws.NewErrParamRequired("Fleet"))
	}
	if s.Fleet != nil && len(*s.Fleet) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fleet", 1))
	}
	if s.DeploymentApplicationConfigs != nil {
		for i, v := range s.DeploymentApplicationConfigs {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "DeploymentApplicationConfigs", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.DeploymentConfig != nil {
		if err := s.DeploymentConfig.Validate(); err != nil {
			invalidParams.AddNested("DeploymentConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDeploymentJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	var ClientRequestToken string
	if s.ClientRequestToken != nil {
		ClientRequestToken = *s.ClientRequestToken
	} else {
		ClientRequestToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DeploymentApplicationConfigs) > 0 {
		v := s.DeploymentApplicationConfigs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "deploymentApplicationConfigs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.DeploymentConfig != nil {
		v := s.DeploymentConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "deploymentConfig", v, metadata)
	}
	if s.Fleet != nil {
		v := *s.Fleet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fleet", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Tags) > 0 {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateDeploymentJobResponse
type CreateDeploymentJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the deployment job.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	// The deployment application configuration.
	DeploymentApplicationConfigs []DeploymentApplicationConfig `locationName:"deploymentApplicationConfigs" min:"1" type:"list"`

	// The deployment configuration.
	DeploymentConfig *DeploymentConfig `locationName:"deploymentConfig" type:"structure"`

	// The failure code of the simulation job if it failed:
	//
	// BadPermissionError
	//
	// AWS Greengrass requires a service-level role permission to access other services.
	// The role must include the AWSGreengrassResourceAccessRolePolicy managed policy
	// (https://console.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSGreengrassResourceAccessRolePolicy$jsonEditor).
	//
	// ExtractingBundleFailure
	//
	// The robot application could not be extracted from the bundle.
	//
	// FailureThresholdBreached
	//
	// The percentage of robots that could not be updated exceeded the percentage
	// set for the deployment.
	//
	// GreengrassDeploymentFailed
	//
	// The robot application could not be deployed to the robot.
	//
	// GreengrassGroupVersionDoesNotExist
	//
	// The AWS Greengrass group or version associated with a robot is missing.
	//
	// InternalServerError
	//
	// An internal error has occurred. Retry your request, but if the problem persists,
	// contact us with details.
	//
	// MissingRobotApplicationArchitecture
	//
	// The robot application does not have a source that matches the architecture
	// of the robot.
	//
	// MissingRobotDeploymentResource
	//
	// One or more of the resources specified for the robot application are missing.
	// For example, does the robot application have the correct launch package and
	// launch file?
	//
	// PostLaunchFileFailure
	//
	// The post-launch script failed.
	//
	// PreLaunchFileFailure
	//
	// The pre-launch script failed.
	//
	// ResourceNotFound
	//
	// One or more deployment resources are missing. For example, do robot application
	// source bundles still exist?
	//
	// RobotDeploymentNoResponse
	//
	// There is no response from the robot. It might not be powered on or connected
	// to the internet.
	FailureCode DeploymentJobErrorCode `locationName:"failureCode" type:"string" enum:"true"`

	// The failure reason of the deployment job if it failed.
	FailureReason *string `locationName:"failureReason" type:"string"`

	// The target fleet for the deployment job.
	Fleet *string `locationName:"fleet" min:"1" type:"string"`

	// The status of the deployment job.
	Status DeploymentStatus `locationName:"status" type:"string" enum:"true"`

	// The list of all tags added to the deployment job.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateDeploymentJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDeploymentJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdAt", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if len(s.DeploymentApplicationConfigs) > 0 {
		v := s.DeploymentApplicationConfigs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "deploymentApplicationConfigs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.DeploymentConfig != nil {
		v := s.DeploymentConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "deploymentConfig", v, metadata)
	}
	if len(s.FailureCode) > 0 {
		v := s.FailureCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "failureCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.FailureReason != nil {
		v := *s.FailureReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "failureReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Fleet != nil {
		v := *s.Fleet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fleet", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.Tags) > 0 {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opCreateDeploymentJob = "CreateDeploymentJob"

// CreateDeploymentJobRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Deploys a specific version of a robot application to robots in a fleet.
//
// The robot application must have a numbered applicationVersion for consistency
// reasons. To create a new version, use CreateRobotApplicationVersion or see
// Creating a Robot Application Version (https://docs.aws.amazon.com/robomaker/latest/dg/create-robot-application-version.html).
//
// After 90 days, deployment jobs expire and will be deleted. They will no longer
// be accessible.
//
//    // Example sending a request using CreateDeploymentJobRequest.
//    req := client.CreateDeploymentJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateDeploymentJob
func (c *Client) CreateDeploymentJobRequest(input *CreateDeploymentJobInput) CreateDeploymentJobRequest {
	op := &aws.Operation{
		Name:       opCreateDeploymentJob,
		HTTPMethod: "POST",
		HTTPPath:   "/createDeploymentJob",
	}

	if input == nil {
		input = &CreateDeploymentJobInput{}
	}

	req := c.newRequest(op, input, &CreateDeploymentJobOutput{})
	return CreateDeploymentJobRequest{Request: req, Input: input, Copy: c.CreateDeploymentJobRequest}
}

// CreateDeploymentJobRequest is the request type for the
// CreateDeploymentJob API operation.
type CreateDeploymentJobRequest struct {
	*aws.Request
	Input *CreateDeploymentJobInput
	Copy  func(*CreateDeploymentJobInput) CreateDeploymentJobRequest
}

// Send marshals and sends the CreateDeploymentJob API request.
func (r CreateDeploymentJobRequest) Send(ctx context.Context) (*CreateDeploymentJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeploymentJobResponse{
		CreateDeploymentJobOutput: r.Request.Data.(*CreateDeploymentJobOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeploymentJobResponse is the response type for the
// CreateDeploymentJob API operation.
type CreateDeploymentJobResponse struct {
	*CreateDeploymentJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeploymentJob request.
func (r *CreateDeploymentJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
