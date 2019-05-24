// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DescribeSimulationJobRequest
type DescribeSimulationJobInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the simulation job to be described.
	//
	// Job is a required field
	Job *string `locationName:"job" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSimulationJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSimulationJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSimulationJobInput"}

	if s.Job == nil {
		invalidParams.Add(aws.NewErrParamRequired("Job"))
	}
	if s.Job != nil && len(*s.Job) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Job", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSimulationJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Job != nil {
		v := *s.Job

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "job", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DescribeSimulationJobResponse
type DescribeSimulationJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the simulation job.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string"`

	// The failure behavior for the simulation job.
	FailureBehavior FailureBehavior `locationName:"failureBehavior" type:"string" enum:"true"`

	// The failure code of the simulation job if it failed:
	//
	// InternalServiceError
	//
	// Internal service error.
	//
	// RobotApplicationCrash
	//
	// Robot application exited abnormally.
	//
	// SimulationApplicationCrash
	//
	// Simulation application exited abnormally.
	//
	// BadPermissionsRobotApplication
	//
	// Robot application bundle could not be downloaded.
	//
	// BadPermissionsSimulationApplication
	//
	// Simulation application bundle could not be downloaded.
	//
	// BadPermissionsS3Output
	//
	// Unable to publish outputs to customer-provided S3 bucket.
	//
	// BadPermissionsCloudwatchLogs
	//
	// Unable to publish logs to customer-provided CloudWatch Logs resource.
	//
	// SubnetIpLimitExceeded
	//
	// Subnet IP limit exceeded.
	//
	// ENILimitExceeded
	//
	// ENI limit exceeded.
	//
	// BadPermissionsUserCredentials
	//
	// Unable to use the Role provided.
	//
	// InvalidBundleRobotApplication
	//
	// Robot bundle cannot be extracted (invalid format, bundling error, or other
	// issue).
	//
	// InvalidBundleSimulationApplication
	//
	// Simulation bundle cannot be extracted (invalid format, bundling error, or
	// other issue).
	//
	// RobotApplicationVersionMismatchedEtag
	//
	// Etag for RobotApplication does not match value during version creation.
	//
	// SimulationApplicationVersionMismatchedEtag
	//
	// Etag for SimulationApplication does not match value during version creation.
	FailureCode SimulationJobErrorCode `locationName:"failureCode" type:"string" enum:"true"`

	// Details about why the simulation job failed. For more information about troubleshooting,
	// see Troubleshooting (https://docs.aws.amazon.com/robomaker/latest/dg/troubleshooting.html).
	FailureReason *string `locationName:"failureReason" type:"string"`

	// The IAM role that allows the simulation instance to call the AWS APIs that
	// are specified in its associated policies on your behalf.
	IamRole *string `locationName:"iamRole" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the simulation job was last
	// updated.
	LastUpdatedAt *time.Time `locationName:"lastUpdatedAt" type:"timestamp" timestampFormat:"unix"`

	// The maximum job duration in seconds. The value must be 8 days (691,200 seconds)
	// or less.
	MaxJobDurationInSeconds *int64 `locationName:"maxJobDurationInSeconds" type:"long"`

	// The name of the simulation job.
	Name *string `locationName:"name" min:"1" type:"string"`

	// Location for output files generated by the simulation job.
	OutputLocation *OutputLocation `locationName:"outputLocation" type:"structure"`

	// A list of robot applications.
	RobotApplications []RobotApplicationConfig `locationName:"robotApplications" min:"1" type:"list"`

	// A list of simulation applications.
	SimulationApplications []SimulationApplicationConfig `locationName:"simulationApplications" min:"1" type:"list"`

	// The simulation job execution duration in milliseconds.
	SimulationTimeMillis *int64 `locationName:"simulationTimeMillis" type:"long"`

	// The status of the simulation job.
	Status SimulationJobStatus `locationName:"status" type:"string" enum:"true"`

	// The list of all tags added to the specified simulation job.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The VPC configuration.
	VpcConfig *VPCConfigResponse `locationName:"vpcConfig" type:"structure"`
}

// String returns the string representation
func (s DescribeSimulationJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSimulationJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClientRequestToken != nil {
		v := *s.ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.FailureBehavior) > 0 {
		v := s.FailureBehavior

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "failureBehavior", protocol.QuotedValue{ValueMarshaler: v}, metadata)
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
	if s.IamRole != nil {
		v := *s.IamRole

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "iamRole", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedAt != nil {
		v := *s.LastUpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedAt", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.MaxJobDurationInSeconds != nil {
		v := *s.MaxJobDurationInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxJobDurationInSeconds", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OutputLocation != nil {
		v := s.OutputLocation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "outputLocation", v, metadata)
	}
	if len(s.RobotApplications) > 0 {
		v := s.RobotApplications

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "robotApplications", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.SimulationApplications) > 0 {
		v := s.SimulationApplications

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "simulationApplications", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SimulationTimeMillis != nil {
		v := *s.SimulationTimeMillis

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "simulationTimeMillis", protocol.Int64Value(v), metadata)
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
	if s.VpcConfig != nil {
		v := s.VpcConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "vpcConfig", v, metadata)
	}
	return nil
}

const opDescribeSimulationJob = "DescribeSimulationJob"

// DescribeSimulationJobRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Describes a simulation job.
//
//    // Example sending a request using DescribeSimulationJobRequest.
//    req := client.DescribeSimulationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DescribeSimulationJob
func (c *Client) DescribeSimulationJobRequest(input *DescribeSimulationJobInput) DescribeSimulationJobRequest {
	op := &aws.Operation{
		Name:       opDescribeSimulationJob,
		HTTPMethod: "POST",
		HTTPPath:   "/describeSimulationJob",
	}

	if input == nil {
		input = &DescribeSimulationJobInput{}
	}

	req := c.newRequest(op, input, &DescribeSimulationJobOutput{})
	return DescribeSimulationJobRequest{Request: req, Input: input, Copy: c.DescribeSimulationJobRequest}
}

// DescribeSimulationJobRequest is the request type for the
// DescribeSimulationJob API operation.
type DescribeSimulationJobRequest struct {
	*aws.Request
	Input *DescribeSimulationJobInput
	Copy  func(*DescribeSimulationJobInput) DescribeSimulationJobRequest
}

// Send marshals and sends the DescribeSimulationJob API request.
func (r DescribeSimulationJobRequest) Send(ctx context.Context) (*DescribeSimulationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSimulationJobResponse{
		DescribeSimulationJobOutput: r.Request.Data.(*DescribeSimulationJobOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSimulationJobResponse is the response type for the
// DescribeSimulationJob API operation.
type DescribeSimulationJobResponse struct {
	*DescribeSimulationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSimulationJob request.
func (r *DescribeSimulationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
