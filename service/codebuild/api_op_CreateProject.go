// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/CreateProjectInput
type CreateProjectInput struct {
	_ struct{} `type:"structure"`

	// Information about the build output artifacts for the build project.
	//
	// Artifacts is a required field
	Artifacts *ProjectArtifacts `locationName:"artifacts" type:"structure" required:"true"`

	// Set this to true to generate a publicly accessible URL for your project's
	// build badge.
	BadgeEnabled *bool `locationName:"badgeEnabled" type:"boolean"`

	// Stores recently used information so that it can be quickly accessed at a
	// later time.
	Cache *ProjectCache `locationName:"cache" type:"structure"`

	// A description that makes the build project easy to identify.
	Description *string `locationName:"description" type:"string"`

	// The AWS Key Management Service (AWS KMS) customer master key (CMK) to be
	// used for encrypting the build output artifacts.
	//
	// You can use a cross-account KMS key to encrypt the build output artifacts
	// if your service role has permission to that key.
	//
	// You can specify either the Amazon Resource Name (ARN) of the CMK or, if available,
	// the CMK's alias (using the format alias/alias-name ).
	EncryptionKey *string `locationName:"encryptionKey" min:"1" type:"string"`

	// Information about the build environment for the build project.
	//
	// Environment is a required field
	Environment *ProjectEnvironment `locationName:"environment" type:"structure" required:"true"`

	// Information about logs for the build project. These can be logs in Amazon
	// CloudWatch Logs, logs uploaded to a specified S3 bucket, or both.
	LogsConfig *LogsConfig `locationName:"logsConfig" type:"structure"`

	// The name of the build project.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"2" type:"string" required:"true"`

	// The number of minutes a build is allowed to be queued before it times out.
	QueuedTimeoutInMinutes *int64 `locationName:"queuedTimeoutInMinutes" min:"5" type:"integer"`

	// An array of ProjectArtifacts objects.
	SecondaryArtifacts []ProjectArtifacts `locationName:"secondaryArtifacts" type:"list"`

	// An array of ProjectSource objects.
	SecondarySources []ProjectSource `locationName:"secondarySources" type:"list"`

	// The ARN of the AWS Identity and Access Management (IAM) role that enables
	// AWS CodeBuild to interact with dependent AWS services on behalf of the AWS
	// account.
	//
	// ServiceRole is a required field
	ServiceRole *string `locationName:"serviceRole" min:"1" type:"string" required:"true"`

	// Information about the build input source code for the build project.
	//
	// Source is a required field
	Source *ProjectSource `locationName:"source" type:"structure" required:"true"`

	// A set of tags for this build project.
	//
	// These tags are available for use by AWS services that support AWS CodeBuild
	// build project tags.
	Tags []Tag `locationName:"tags" type:"list"`

	// How long, in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait
	// before it times out any build that has not been marked as completed. The
	// default is 60 minutes.
	TimeoutInMinutes *int64 `locationName:"timeoutInMinutes" min:"5" type:"integer"`

	// VpcConfig enables AWS CodeBuild to access resources in an Amazon VPC.
	VpcConfig *VpcConfig `locationName:"vpcConfig" type:"structure"`
}

// String returns the string representation
func (s CreateProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProjectInput"}

	if s.Artifacts == nil {
		invalidParams.Add(aws.NewErrParamRequired("Artifacts"))
	}
	if s.EncryptionKey != nil && len(*s.EncryptionKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EncryptionKey", 1))
	}

	if s.Environment == nil {
		invalidParams.Add(aws.NewErrParamRequired("Environment"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 2))
	}
	if s.QueuedTimeoutInMinutes != nil && *s.QueuedTimeoutInMinutes < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("QueuedTimeoutInMinutes", 5))
	}

	if s.ServiceRole == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceRole"))
	}
	if s.ServiceRole != nil && len(*s.ServiceRole) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceRole", 1))
	}

	if s.Source == nil {
		invalidParams.Add(aws.NewErrParamRequired("Source"))
	}
	if s.TimeoutInMinutes != nil && *s.TimeoutInMinutes < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("TimeoutInMinutes", 5))
	}
	if s.Artifacts != nil {
		if err := s.Artifacts.Validate(); err != nil {
			invalidParams.AddNested("Artifacts", err.(aws.ErrInvalidParams))
		}
	}
	if s.Cache != nil {
		if err := s.Cache.Validate(); err != nil {
			invalidParams.AddNested("Cache", err.(aws.ErrInvalidParams))
		}
	}
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			invalidParams.AddNested("Environment", err.(aws.ErrInvalidParams))
		}
	}
	if s.LogsConfig != nil {
		if err := s.LogsConfig.Validate(); err != nil {
			invalidParams.AddNested("LogsConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.SecondaryArtifacts != nil {
		for i, v := range s.SecondaryArtifacts {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SecondaryArtifacts", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SecondarySources != nil {
		for i, v := range s.SecondarySources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SecondarySources", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			invalidParams.AddNested("Source", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			invalidParams.AddNested("VpcConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/CreateProjectOutput
type CreateProjectOutput struct {
	_ struct{} `type:"structure"`

	// Information about the build project that was created.
	Project *Project `locationName:"project" type:"structure"`
}

// String returns the string representation
func (s CreateProjectOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateProject = "CreateProject"

// CreateProjectRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Creates a build project.
//
//    // Example sending a request using CreateProjectRequest.
//    req := client.CreateProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/CreateProject
func (c *Client) CreateProjectRequest(input *CreateProjectInput) CreateProjectRequest {
	op := &aws.Operation{
		Name:       opCreateProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateProjectInput{}
	}

	req := c.newRequest(op, input, &CreateProjectOutput{})
	return CreateProjectRequest{Request: req, Input: input, Copy: c.CreateProjectRequest}
}

// CreateProjectRequest is the request type for the
// CreateProject API operation.
type CreateProjectRequest struct {
	*aws.Request
	Input *CreateProjectInput
	Copy  func(*CreateProjectInput) CreateProjectRequest
}

// Send marshals and sends the CreateProject API request.
func (r CreateProjectRequest) Send(ctx context.Context) (*CreateProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProjectResponse{
		CreateProjectOutput: r.Request.Data.(*CreateProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProjectResponse is the response type for the
// CreateProject API operation.
type CreateProjectResponse struct {
	*CreateProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProject request.
func (r *CreateProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
