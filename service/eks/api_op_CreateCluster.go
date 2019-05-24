// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/CreateClusterRequest
type CreateClusterInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string `locationName:"clientRequestToken" type:"string" idempotencyToken:"true"`

	// Enable or disable exporting the Kubernetes control plane logs for your cluster
	// to CloudWatch Logs. By default, cluster control plane logs are not exported
	// to CloudWatch Logs. For more information, see Amazon EKS Cluster Control
	// Plane Logs (https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
	// in the Amazon EKS User Guide .
	//
	// CloudWatch Logs ingestion, archive storage, and data scanning rates apply
	// to exported control plane logs. For more information, see Amazon CloudWatch
	// Pricing (http://aws.amazon.com/cloudwatch/pricing/).
	Logging *Logging `locationName:"logging" type:"structure"`

	// The unique name to give to your cluster.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The VPC configuration used by the cluster control plane. Amazon EKS VPC resources
	// have specific requirements to work properly with Kubernetes. For more information,
	// see Cluster VPC Considerations (https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html)
	// and Cluster Security Group Considerations (https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html)
	// in the Amazon EKS User Guide. You must specify at least two subnets. You
	// may specify up to five security groups, but we recommend that you use a dedicated
	// security group for your cluster control plane.
	//
	// ResourcesVpcConfig is a required field
	ResourcesVpcConfig *VpcConfigRequest `locationName:"resourcesVpcConfig" type:"structure" required:"true"`

	// The Amazon Resource Name (ARN) of the IAM role that provides permissions
	// for Amazon EKS to make calls to other AWS API operations on your behalf.
	// For more information, see Amazon EKS Service IAM Role (https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html)
	// in the Amazon EKS User Guide .
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" type:"string" required:"true"`

	// The desired Kubernetes version for your cluster. If you do not specify a
	// value here, the latest version available in Amazon EKS is used.
	Version *string `locationName:"version" type:"string"`
}

// String returns the string representation
func (s CreateClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClusterInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.ResourcesVpcConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourcesVpcConfig"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateClusterInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.Logging != nil {
		v := s.Logging

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "logging", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourcesVpcConfig != nil {
		v := s.ResourcesVpcConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "resourcesVpcConfig", v, metadata)
	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/CreateClusterResponse
type CreateClusterOutput struct {
	_ struct{} `type:"structure"`

	// The full description of your new cluster.
	Cluster *Cluster `locationName:"cluster" type:"structure"`
}

// String returns the string representation
func (s CreateClusterOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateClusterOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Cluster != nil {
		v := s.Cluster

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "cluster", v, metadata)
	}
	return nil
}

const opCreateCluster = "CreateCluster"

// CreateClusterRequest returns a request value for making API operation for
// Amazon Elastic Container Service for Kubernetes.
//
// Creates an Amazon EKS control plane.
//
// The Amazon EKS control plane consists of control plane instances that run
// the Kubernetes software, like etcd and the API server. The control plane
// runs in an account managed by AWS, and the Kubernetes API is exposed via
// the Amazon EKS API server endpoint. Each Amazon EKS cluster control plane
// is single-tenant and unique, and runs on its own set of Amazon EC2 instances.
//
// The cluster control plane is provisioned across multiple Availability Zones
// and fronted by an Elastic Load Balancing Network Load Balancer. Amazon EKS
// also provisions elastic network interfaces in your VPC subnets to provide
// connectivity from the control plane instances to the worker nodes (for example,
// to support kubectl exec, logs, and proxy data flows).
//
// Amazon EKS worker nodes run in your AWS account and connect to your cluster's
// control plane via the Kubernetes API server endpoint and a certificate file
// that is created for your cluster.
//
// You can use the endpointPublicAccess and endpointPrivateAccess parameters
// to enable or disable public and private access to your cluster's Kubernetes
// API server endpoint. By default, public access is enabled and private access
// is disabled. For more information, see Amazon EKS Cluster Endpoint Access
// Control (https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html)
// in the Amazon EKS User Guide .
//
// You can use the logging parameter to enable or disable exporting the Kubernetes
// control plane logs for your cluster to CloudWatch Logs. By default, cluster
// control plane logs are not exported to CloudWatch Logs. For more information,
// see Amazon EKS Cluster Control Plane Logs (https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
// in the Amazon EKS User Guide .
//
// CloudWatch Logs ingestion, archive storage, and data scanning rates apply
// to exported control plane logs. For more information, see Amazon CloudWatch
// Pricing (http://aws.amazon.com/cloudwatch/pricing/).
//
// Cluster creation typically takes between 10 and 15 minutes. After you create
// an Amazon EKS cluster, you must configure your Kubernetes tooling to communicate
// with the API server and launch worker nodes into your cluster. For more information,
// see Managing Cluster Authentication (https://docs.aws.amazon.com/eks/latest/userguide/managing-auth.html)
// and Launching Amazon EKS Worker Nodes (https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html)
// in the Amazon EKS User Guide.
//
//    // Example sending a request using CreateClusterRequest.
//    req := client.CreateClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/CreateCluster
func (c *Client) CreateClusterRequest(input *CreateClusterInput) CreateClusterRequest {
	op := &aws.Operation{
		Name:       opCreateCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/clusters",
	}

	if input == nil {
		input = &CreateClusterInput{}
	}

	req := c.newRequest(op, input, &CreateClusterOutput{})
	return CreateClusterRequest{Request: req, Input: input, Copy: c.CreateClusterRequest}
}

// CreateClusterRequest is the request type for the
// CreateCluster API operation.
type CreateClusterRequest struct {
	*aws.Request
	Input *CreateClusterInput
	Copy  func(*CreateClusterInput) CreateClusterRequest
}

// Send marshals and sends the CreateCluster API request.
func (r CreateClusterRequest) Send(ctx context.Context) (*CreateClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClusterResponse{
		CreateClusterOutput: r.Request.Data.(*CreateClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClusterResponse is the response type for the
// CreateCluster API operation.
type CreateClusterResponse struct {
	*CreateClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCluster request.
func (r *CreateClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
