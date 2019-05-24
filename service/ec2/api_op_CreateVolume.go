// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CreateVolume.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVolumeRequest
type CreateVolumeInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone in which to create the volume. Use DescribeAvailabilityZones
	// to list the Availability Zones that are currently available to you.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// Specifies whether the volume should be encrypted. Encrypted Amazon EBS volumes
	// may only be attached to instances that support Amazon EBS encryption. Volumes
	// that are created from encrypted snapshots are automatically encrypted. There
	// is no way to create an encrypted volume from an unencrypted snapshot or vice
	// versa. If your AMI uses encrypted volumes, you can only launch it on supported
	// instance types. For more information, see Amazon EBS Encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	Encrypted *bool `locationName:"encrypted" type:"boolean"`

	// The number of I/O operations per second (IOPS) to provision for the volume,
	// with a maximum ratio of 50 IOPS/GiB. Range is 100 to 64,000 IOPS for volumes
	// in most Regions. Maximum IOPS of 64,000 is guaranteed only on Nitro-based
	// instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances).
	// Other instance families guarantee performance up to 32,000 IOPS. For more
	// information, see Amazon EBS Volume Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// This parameter is valid only for Provisioned IOPS SSD (io1) volumes.
	Iops *int64 `type:"integer"`

	// An identifier for the AWS Key Management Service (AWS KMS) customer master
	// key (CMK) to use when creating the encrypted volume. This parameter is only
	// required if you want to use a non-default CMK; if this parameter is not specified,
	// the default CMK for EBS is used. If a KmsKeyId is specified, the Encrypted
	// flag must also be set.
	//
	// The CMK identifier may be provided in any of the following formats:
	//
	//    * Key ID
	//
	//    * Key alias. The alias ARN contains the arn:aws:kms namespace, followed
	//    by the region of the CMK, the AWS account ID of the CMK owner, the alias
	//    namespace, and then the CMK alias. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	//    * ARN using key ID. The ID ARN contains the arn:aws:kms namespace, followed
	//    by the region of the CMK, the AWS account ID of the CMK owner, the key
	//    namespace, and then the CMK ID. For example, arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	//    * ARN using key alias. The alias ARN contains the arn:aws:kms namespace,
	//    followed by the region of the CMK, the AWS account ID of the CMK owner,
	//    the alias namespace, and then the CMK alias. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// AWS parses KmsKeyId asynchronously, meaning that the action you call may
	// appear to complete even though you provided an invalid identifier. The action
	// will eventually fail.
	KmsKeyId *string `type:"string"`

	// The size of the volume, in GiBs.
	//
	// Constraints: 1-16,384 for gp2, 4-16,384 for io1, 500-16,384 for st1, 500-16,384
	// for sc1, and 1-1,024 for standard. If you specify a snapshot, the volume
	// size must be equal to or larger than the snapshot size.
	//
	// Default: If you're creating the volume from a snapshot and don't specify
	// a volume size, the default is the snapshot size.
	//
	// At least one of Size or SnapshotId are required.
	Size *int64 `type:"integer"`

	// The snapshot from which to create the volume.
	//
	// At least one of Size or SnapshotId are required.
	SnapshotId *string `type:"string"`

	// The tags to apply to the volume during creation.
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`

	// The volume type. This can be gp2 for General Purpose SSD, io1 for Provisioned
	// IOPS SSD, st1 for Throughput Optimized HDD, sc1 for Cold HDD, or standard
	// for Magnetic volumes.
	//
	// Defaults: If no volume type is specified, the default is standard in us-east-1,
	// eu-west-1, eu-central-1, us-west-2, us-west-1, sa-east-1, ap-northeast-1,
	// ap-northeast-2, ap-southeast-1, ap-southeast-2, ap-south-1, us-gov-west-1,
	// and cn-north-1. In all other Regions, EBS defaults to gp2.
	VolumeType VolumeType `type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVolumeInput"}

	if s.AvailabilityZone == nil {
		invalidParams.Add(aws.NewErrParamRequired("AvailabilityZone"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes a volume.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/Volume
type CreateVolumeOutput struct {
	_ struct{} `type:"structure"`

	// Information about the volume attachments.
	Attachments []VolumeAttachment `locationName:"attachmentSet" locationNameList:"item" type:"list"`

	// The Availability Zone for the volume.
	AvailabilityZone *string `locationName:"availabilityZone" type:"string"`

	// The time stamp when volume creation was initiated.
	CreateTime *time.Time `locationName:"createTime" type:"timestamp" timestampFormat:"iso8601"`

	// Indicates whether the volume will be encrypted.
	Encrypted *bool `locationName:"encrypted" type:"boolean"`

	// The number of I/O operations per second (IOPS) that the volume supports.
	// For Provisioned IOPS SSD volumes, this represents the number of IOPS that
	// are provisioned for the volume. For General Purpose SSD volumes, this represents
	// the baseline performance of the volume and the rate at which the volume accumulates
	// I/O credits for bursting. For more information, see Amazon EBS Volume Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// Constraints: Range is 100-16,000 IOPS for gp2 volumes and 100 to 64,000IOPS
	// for io1 volumes, in most Regions. The maximum IOPS for io1 of 64,000 is guaranteed
	// only on Nitro-based instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances).
	// Other instance families guarantee performance up to 32,000 IOPS.
	//
	// Condition: This parameter is required for requests to create io1 volumes;
	// it is not used in requests to create gp2, st1, sc1, or standard volumes.
	Iops *int64 `locationName:"iops" type:"integer"`

	// The full ARN of the AWS Key Management Service (AWS KMS) customer master
	// key (CMK) that was used to protect the volume encryption key for the volume.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// The size of the volume, in GiBs.
	Size *int64 `locationName:"size" type:"integer"`

	// The snapshot from which the volume was created, if applicable.
	SnapshotId *string `locationName:"snapshotId" type:"string"`

	// The volume state.
	State VolumeState `locationName:"status" type:"string" enum:"true"`

	// Any tags assigned to the volume.
	Tags []Tag `locationName:"tagSet" locationNameList:"item" type:"list"`

	// The ID of the volume.
	VolumeId *string `locationName:"volumeId" type:"string"`

	// The volume type. This can be gp2 for General Purpose SSD, io1 for Provisioned
	// IOPS SSD, st1 for Throughput Optimized HDD, sc1 for Cold HDD, or standard
	// for Magnetic volumes.
	VolumeType VolumeType `locationName:"volumeType" type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateVolumeOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVolume = "CreateVolume"

// CreateVolumeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates an EBS volume that can be attached to an instance in the same Availability
// Zone. The volume is created in the regional endpoint that you send the HTTP
// request to. For more information see Regions and Endpoints (https://docs.aws.amazon.com/general/latest/gr/rande.html).
//
// You can create a new empty volume or restore a volume from an EBS snapshot.
// Any AWS Marketplace product codes from the snapshot are propagated to the
// volume.
//
// You can create encrypted volumes with the Encrypted parameter. Encrypted
// volumes may only be attached to instances that support Amazon EBS encryption.
// Volumes that are created from encrypted snapshots are also automatically
// encrypted. For more information, see Amazon EBS Encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// You can tag your volumes during creation. For more information, see Tagging
// Your Amazon EC2 Resources (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// For more information, see Creating an Amazon EBS Volume (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-creating-volume.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using CreateVolumeRequest.
//    req := client.CreateVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVolume
func (c *Client) CreateVolumeRequest(input *CreateVolumeInput) CreateVolumeRequest {
	op := &aws.Operation{
		Name:       opCreateVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVolumeInput{}
	}

	req := c.newRequest(op, input, &CreateVolumeOutput{})
	return CreateVolumeRequest{Request: req, Input: input, Copy: c.CreateVolumeRequest}
}

// CreateVolumeRequest is the request type for the
// CreateVolume API operation.
type CreateVolumeRequest struct {
	*aws.Request
	Input *CreateVolumeInput
	Copy  func(*CreateVolumeInput) CreateVolumeRequest
}

// Send marshals and sends the CreateVolume API request.
func (r CreateVolumeRequest) Send(ctx context.Context) (*CreateVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVolumeResponse{
		CreateVolumeOutput: r.Request.Data.(*CreateVolumeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVolumeResponse is the response type for the
// CreateVolume API operation.
type CreateVolumeResponse struct {
	*CreateVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVolume request.
func (r *CreateVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
