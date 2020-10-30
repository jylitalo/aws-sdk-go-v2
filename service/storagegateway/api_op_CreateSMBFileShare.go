// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Server Message Block (SMB) file share on an existing file gateway. In
// Storage Gateway, a file share is a file system mount point backed by Amazon S3
// cloud storage. Storage Gateway exposes file shares using an SMB interface. This
// operation is only supported for file gateways. File gateways require AWS
// Security Token Service (AWS STS) to be activated to enable you to create a file
// share. Make sure that AWS STS is activated in the AWS Region you are creating
// your file gateway in. If AWS STS is not activated in this AWS Region, activate
// it. For information about how to activate AWS STS, see Activating and
// deactivating AWS STS in an AWS Region
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
// in the AWS Identity and Access Management User Guide. File gateways don't
// support creating hard or symbolic links on a file share.
func (c *Client) CreateSMBFileShare(ctx context.Context, params *CreateSMBFileShareInput, optFns ...func(*Options)) (*CreateSMBFileShareOutput, error) {
	if params == nil {
		params = &CreateSMBFileShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSMBFileShare", params, optFns, addOperationCreateSMBFileShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSMBFileShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateSMBFileShareInput
type CreateSMBFileShareInput struct {

	// A unique string value that you supply that is used by file gateway to ensure
	// idempotent file share creation.
	//
	// This member is required.
	ClientToken *string

	// The ARN of the file gateway on which you want to create a file share.
	//
	// This member is required.
	GatewayARN *string

	// The ARN of the backend storage used for storing file data. A prefix name can be
	// added to the S3 bucket name. It must end with a "/".
	//
	// This member is required.
	LocationARN *string

	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway
	// assumes when it accesses the underlying storage.
	//
	// This member is required.
	Role *string

	// The files and folders on this share will only be visible to users with read
	// access.
	AccessBasedEnumeration *bool

	// A list of users or groups in the Active Directory that will be granted
	// administrator privileges on the file share. These users can do all file
	// operations as the super-user. Acceptable formats include: DOMAIN\User1, user1,
	// @group1, and @DOMAIN\group1. Use this option very carefully, because any user in
	// this list can do anything they like on the file share, regardless of file
	// permissions.
	AdminUserList []*string

	// The Amazon Resource Name (ARN) of the storage used for the audit logs.
	AuditDestinationARN *string

	// The authentication method that users use to access the file share. The default
	// is ActiveDirectory. Valid Values: ActiveDirectory | GuestAccess
	Authentication *string

	// Refresh cache information.
	CacheAttributes *types.CacheAttributes

	// The case of an object name in an Amazon S3 bucket. For ClientSpecified, the
	// client determines the case sensitivity. For CaseSensitive, the gateway
	// determines the case sensitivity. The default value is ClientSpecified.
	CaseSensitivity types.CaseSensitivity

	// The default storage class for objects put into an Amazon S3 bucket by the file
	// gateway. The default value is S3_INTELLIGENT_TIERING. Optional. Valid Values:
	// S3_STANDARD | S3_INTELLIGENT_TIERING | S3_STANDARD_IA | S3_ONEZONE_IA
	DefaultStorageClass *string

	// The name of the file share. Optional. FileShareName must be set if an S3 prefix
	// name is set in LocationARN.
	FileShareName *string

	// A value that enables guessing of the MIME type for uploaded objects based on
	// file extensions. Set this value to true to enable MIME type guessing, otherwise
	// set to false. The default value is true. Valid Values: true | false
	GuessMIMETypeEnabled *bool

	// A list of users or groups in the Active Directory that are not allowed to access
	// the file share. A group must be prefixed with the @ character. Acceptable
	// formats include: DOMAIN\User1, user1, @group1, and @DOMAIN\group1. Can only be
	// set if Authentication is set to ActiveDirectory.
	InvalidUserList []*string

	// Set to true to use Amazon S3 server-side encryption with your own AWS KMS key,
	// or false to use a key managed by Amazon S3. Optional. Valid Values: true | false
	KMSEncrypted *bool

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for
	// Amazon S3 server-side encryption. Storage Gateway does not support asymmetric
	// CMKs. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string

	// The notification policy of the file share.
	NotificationPolicy *string

	// A value that sets the access control list (ACL) permission for objects in the S3
	// bucket that a file gateway puts objects into. The default value is private.
	ObjectACL types.ObjectACL

	// A value that sets the write status of a file share. Set this value to true to
	// set the write status to read-only, otherwise set to false. Valid Values: true |
	// false
	ReadOnly *bool

	// A value that sets who pays the cost of the request and the cost associated with
	// data download from the S3 bucket. If this value is set to true, the requester
	// pays the costs; otherwise, the S3 bucket owner pays. However, the S3 bucket
	// owner always pays the cost of storing data. RequesterPays is a configuration for
	// the S3 bucket that backs the file share, so make sure that the configuration on
	// the file share is the same as the S3 bucket configuration. Valid Values: true |
	// false
	RequesterPays *bool

	// Set this value to true to enable access control list (ACL) on the SMB file
	// share. Set it to false to map file and directory permissions to the POSIX
	// permissions. For more information, see Using Microsoft Windows ACLs to control
	// access to an SMB file share
	// (https://docs.aws.amazon.com/storagegateway/latest/userguide/smb-acl.html) in
	// the AWS Storage Gateway User Guide. Valid Values: true | false
	SMBACLEnabled *bool

	// A list of up to 50 tags that can be assigned to the NFS file share. Each tag is
	// a key-value pair. Valid characters for key and value are letters, spaces, and
	// numbers representable in UTF-8 format, and the following special characters: + -
	// = . _ : / @. The maximum length of a tag's key is 128 characters, and the
	// maximum length for a tag's value is 256.
	Tags []*types.Tag

	// A list of users or groups in the Active Directory that are allowed to access the
	// file  share. A group must be prefixed with the @ character. Acceptable formats
	// include: DOMAIN\User1, user1, @group1, and @DOMAIN\group1. Can only be set if
	// Authentication is set to ActiveDirectory.
	ValidUserList []*string
}

// CreateSMBFileShareOutput
type CreateSMBFileShareOutput struct {

	// The Amazon Resource Name (ARN) of the newly created file share.
	FileShareARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSMBFileShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSMBFileShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSMBFileShare{}, middleware.After)
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
	addOpCreateSMBFileShareValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSMBFileShare(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateSMBFileShare(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CreateSMBFileShare",
	}
}
