// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/CreateStackRequest
type CreateStackInput struct {
	_ struct{} `type:"structure"`

	// The default AWS OpsWorks Stacks agent version. You have the following options:
	//
	//    * Auto-update - Set this parameter to LATEST. AWS OpsWorks Stacks automatically
	//    installs new agent versions on the stack's instances as soon as they are
	//    available.
	//
	//    * Fixed version - Set this parameter to your preferred agent version.
	//    To update the agent version, you must edit the stack configuration and
	//    specify a new version. AWS OpsWorks Stacks then automatically installs
	//    that version on the stack's instances.
	//
	// The default setting is the most recent release of the agent. To specify an
	// agent version, you must use the complete version number, not the abbreviated
	// number shown on the console. For a list of available agent version numbers,
	// call DescribeAgentVersions. AgentVersion cannot be set to Chef 12.2.
	//
	// You can also specify an agent version when you create or update an instance,
	// which overrides the stack's default setting.
	AgentVersion *string `type:"string"`

	// One or more user-defined key-value pairs to be added to the stack attributes.
	Attributes map[string]string `type:"map"`

	// A ChefConfiguration object that specifies whether to enable Berkshelf and
	// the Berkshelf version on Chef 11.10 stacks. For more information, see Create
	// a New Stack (http://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-creating.html).
	ChefConfiguration *ChefConfiguration `type:"structure"`

	// The configuration manager. When you create a stack we recommend that you
	// use the configuration manager to specify the Chef version: 12, 11.10, or
	// 11.4 for Linux stacks, or 12.2 for Windows stacks. The default value for
	// Linux stacks is currently 12.
	ConfigurationManager *StackConfigurationManager `type:"structure"`

	// Contains the information required to retrieve an app or cookbook from a repository.
	// For more information, see Creating Apps (http://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html)
	// or Custom Recipes and Cookbooks (http://docs.aws.amazon.com/opsworks/latest/userguide/workingcookbook.html).
	CustomCookbooksSource *Source `type:"structure"`

	// A string that contains user-defined, custom JSON. It can be used to override
	// the corresponding default stack configuration attribute values or to pass
	// data to recipes. The string should be in the following format:
	//
	// "{\"key1\": \"value1\", \"key2\": \"value2\",...}"
	//
	// For more information about custom JSON, see Use Custom JSON to Modify the
	// Stack Configuration Attributes (http://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-json.html).
	CustomJson *string `type:"string"`

	// The stack's default Availability Zone, which must be in the specified region.
	// For more information, see Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html).
	// If you also specify a value for DefaultSubnetId, the subnet must be in the
	// same zone. For more information, see the VpcId parameter description.
	DefaultAvailabilityZone *string `type:"string"`

	// The Amazon Resource Name (ARN) of an IAM profile that is the default profile
	// for all of the stack's EC2 instances. For more information about IAM ARNs,
	// see Using Identifiers (http://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html).
	//
	// DefaultInstanceProfileArn is a required field
	DefaultInstanceProfileArn *string `type:"string" required:"true"`

	// The stack's default operating system, which is installed on every instance
	// unless you specify a different operating system when you create the instance.
	// You can specify one of the following.
	//
	//    * A supported Linux operating system: An Amazon Linux version, such as
	//    Amazon Linux 2017.09, Amazon Linux 2017.03, Amazon Linux 2016.09, Amazon
	//    Linux 2016.03, Amazon Linux 2015.09, or Amazon Linux 2015.03.
	//
	//    * A supported Ubuntu operating system, such as Ubuntu 16.04 LTS, Ubuntu
	//    14.04 LTS, or Ubuntu 12.04 LTS.
	//
	//    * CentOS Linux 7
	//
	//    * Red Hat Enterprise Linux 7
	//
	//    * A supported Windows operating system, such as Microsoft Windows Server
	//    2012 R2 Base, Microsoft Windows Server 2012 R2 with SQL Server Express,
	//    Microsoft Windows Server 2012 R2 with SQL Server Standard, or Microsoft
	//    Windows Server 2012 R2 with SQL Server Web.
	//
	//    * A custom AMI: Custom. You specify the custom AMI you want to use when
	//    you create instances. For more information, see Using Custom AMIs (http://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-custom-ami.html).
	//
	// The default option is the current Amazon Linux version. For more information
	// about supported operating systems, see AWS OpsWorks Stacks Operating Systems
	// (http://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-os.html).
	DefaultOs *string `type:"string"`

	// The default root device type. This value is the default for all instances
	// in the stack, but you can override it when you create an instance. The default
	// option is instance-store. For more information, see Storage for the Root
	// Device (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ComponentsAMIs.html#storage-for-the-root-device).
	DefaultRootDeviceType RootDeviceType `type:"string" enum:"true"`

	// A default Amazon EC2 key pair name. The default value is none. If you specify
	// a key pair name, AWS OpsWorks installs the public key on the instance and
	// you can use the private key with an SSH client to log in to the instance.
	// For more information, see Using SSH to Communicate with an Instance (http://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-ssh.html)
	// and Managing SSH Access (http://docs.aws.amazon.com/opsworks/latest/userguide/security-ssh-access.html).
	// You can override this setting by specifying a different key pair, or no key
	// pair, when you create an instance (http://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-add.html).
	DefaultSshKeyName *string `type:"string"`

	// The stack's default VPC subnet ID. This parameter is required if you specify
	// a value for the VpcId parameter. All instances are launched into this subnet
	// unless you specify otherwise when you create the instance. If you also specify
	// a value for DefaultAvailabilityZone, the subnet must be in that zone. For
	// information on default values and when this parameter is required, see the
	// VpcId parameter description.
	DefaultSubnetId *string `type:"string"`

	// The stack's host name theme, with spaces replaced by underscores. The theme
	// is used to generate host names for the stack's instances. By default, HostnameTheme
	// is set to Layer_Dependent, which creates host names by appending integers
	// to the layer's short name. The other themes are:
	//
	//    * Baked_Goods
	//
	//    * Clouds
	//
	//    * Europe_Cities
	//
	//    * Fruits
	//
	//    * Greek_Deities
	//
	//    * Legendary_creatures_from_Japan
	//
	//    * Planets_and_Moons
	//
	//    * Roman_Deities
	//
	//    * Scottish_Islands
	//
	//    * US_Cities
	//
	//    * Wild_Cats
	//
	// To obtain a generated host name, call GetHostNameSuggestion, which returns
	// a host name based on the current theme.
	HostnameTheme *string `type:"string"`

	// The stack name.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The stack's AWS region, such as ap-south-1. For more information about Amazon
	// regions, see Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html).
	//
	// In the AWS CLI, this API maps to the --stack-region parameter. If the --stack-region
	// parameter and the AWS CLI common parameter --region are set to the same value,
	// the stack uses a regional endpoint. If the --stack-region parameter is not
	// set, but the AWS CLI --region parameter is, this also results in a stack
	// with a regional endpoint. However, if the --region parameter is set to us-east-1,
	// and the --stack-region parameter is set to one of the following, then the
	// stack uses a legacy or classic region: us-west-1, us-west-2, sa-east-1, eu-central-1,
	// eu-west-1, ap-northeast-1, ap-southeast-1, ap-southeast-2. In this case,
	// the actual API endpoint of the stack is in us-east-1. Only the preceding
	// regions are supported as classic regions in the us-east-1 API endpoint. Because
	// it is a best practice to choose the regional endpoint that is closest to
	// where you manage AWS, we recommend that you use regional endpoints for new
	// stacks. The AWS CLI common --region parameter always specifies a regional
	// API endpoint; it cannot be used to specify a classic AWS OpsWorks Stacks
	// region.
	//
	// Region is a required field
	Region *string `type:"string" required:"true"`

	// The stack's AWS Identity and Access Management (IAM) role, which allows AWS
	// OpsWorks Stacks to work with AWS resources on your behalf. You must set this
	// parameter to the Amazon Resource Name (ARN) for an existing IAM role. For
	// more information about IAM ARNs, see Using Identifiers (http://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html).
	//
	// ServiceRoleArn is a required field
	ServiceRoleArn *string `type:"string" required:"true"`

	// Whether the stack uses custom cookbooks.
	UseCustomCookbooks *bool `type:"boolean"`

	// Whether to associate the AWS OpsWorks Stacks built-in security groups with
	// the stack's layers.
	//
	// AWS OpsWorks Stacks provides a standard set of built-in security groups,
	// one for each layer, which are associated with layers by default. With UseOpsworksSecurityGroups
	// you can instead provide your own custom security groups. UseOpsworksSecurityGroups
	// has the following settings:
	//
	//    * True - AWS OpsWorks Stacks automatically associates the appropriate
	//    built-in security group with each layer (default setting). You can associate
	//    additional security groups with a layer after you create it, but you cannot
	//    delete the built-in security group.
	//
	//    * False - AWS OpsWorks Stacks does not associate built-in security groups
	//    with layers. You must create appropriate EC2 security groups and associate
	//    a security group with each layer that you create. However, you can still
	//    manually associate a built-in security group with a layer on creation;
	//    custom security groups are required only for those layers that need custom
	//    settings.
	//
	// For more information, see Create a New Stack (http://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-creating.html).
	UseOpsworksSecurityGroups *bool `type:"boolean"`

	// The ID of the VPC that the stack is to be launched into. The VPC must be
	// in the stack's region. All instances are launched into this VPC. You cannot
	// change the ID later.
	//
	//    * If your account supports EC2-Classic, the default value is no VPC.
	//
	//    * If your account does not support EC2-Classic, the default value is the
	//    default VPC for the specified region.
	//
	// If the VPC ID corresponds to a default VPC and you have specified either
	// the DefaultAvailabilityZone or the DefaultSubnetId parameter only, AWS OpsWorks
	// Stacks infers the value of the other parameter. If you specify neither parameter,
	// AWS OpsWorks Stacks sets these parameters to the first valid Availability
	// Zone for the specified region and the corresponding default VPC subnet ID,
	// respectively.
	//
	// If you specify a nondefault VPC ID, note the following:
	//
	//    * It must belong to a VPC in your account that is in the specified region.
	//
	//    * You must specify a value for DefaultSubnetId.
	//
	// For more information about how to use AWS OpsWorks Stacks with a VPC, see
	// Running a Stack in a VPC (http://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-vpc.html).
	// For more information about default VPC and EC2-Classic, see Supported Platforms
	// (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html).
	VpcId *string `type:"string"`
}

// String returns the string representation
func (s CreateStackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateStackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateStackInput"}

	if s.DefaultInstanceProfileArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DefaultInstanceProfileArn"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Region == nil {
		invalidParams.Add(aws.NewErrParamRequired("Region"))
	}

	if s.ServiceRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceRoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a CreateStack request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/CreateStackResult
type CreateStackOutput struct {
	_ struct{} `type:"structure"`

	// The stack ID, which is an opaque string that you use to identify the stack
	// when performing actions such as DescribeStacks.
	StackId *string `type:"string"`
}

// String returns the string representation
func (s CreateStackOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateStack = "CreateStack"

// CreateStackRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Creates a new stack. For more information, see Create a New Stack (http://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-edit.html).
//
// Required Permissions: To use this action, an IAM user must have an attached
// policy that explicitly grants permissions. For more information about user
// permissions, see Managing User Permissions (http://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using CreateStackRequest.
//    req := client.CreateStackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/CreateStack
func (c *Client) CreateStackRequest(input *CreateStackInput) CreateStackRequest {
	op := &aws.Operation{
		Name:       opCreateStack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateStackInput{}
	}

	req := c.newRequest(op, input, &CreateStackOutput{})
	return CreateStackRequest{Request: req, Input: input, Copy: c.CreateStackRequest}
}

// CreateStackRequest is the request type for the
// CreateStack API operation.
type CreateStackRequest struct {
	*aws.Request
	Input *CreateStackInput
	Copy  func(*CreateStackInput) CreateStackRequest
}

// Send marshals and sends the CreateStack API request.
func (r CreateStackRequest) Send(ctx context.Context) (*CreateStackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStackResponse{
		CreateStackOutput: r.Request.Data.(*CreateStackOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStackResponse is the response type for the
// CreateStack API operation.
type CreateStackResponse struct {
	*CreateStackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStack request.
func (r *CreateStackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
