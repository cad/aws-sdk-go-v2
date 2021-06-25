// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used by administrators to choose which groups in the directory should have
// access to upload and download files over the enabled protocols using Amazon Web
// Services Transfer Family. For example, a Microsoft Active Directory might
// contain 50,000 users, but only a small fraction might need the ability to
// transfer files to the server. An administrator can use CreateAccess to limit the
// access to the correct set of users who need this ability.
func (c *Client) CreateAccess(ctx context.Context, params *CreateAccessInput, optFns ...func(*Options)) (*CreateAccessOutput, error) {
	if params == nil {
		params = &CreateAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccess", params, optFns, c.addOperationCreateAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessInput struct {

	// A unique identifier that is required to identify specific groups within your
	// directory. The users of the group that you associate have access to your Amazon
	// S3 or Amazon EFS resources over the enabled protocols using Amazon Web Services
	// Transfer Family. If you know the group name, you can view the SID values by
	// running the following command using Windows PowerShell. Get-ADGroup -Filter
	// {samAccountName -like "YourGroupName*"} -Properties * | Select
	// SamAccountName,ObjectSid In that command, replace YourGroupName with the name of
	// your Active Directory group. The regex used to validate this parameter is a
	// string of characters consisting of uppercase and lowercase alphanumeric
	// characters with no spaces. You can also include underscores or any of the
	// following characters: =,.@:/-
	//
	// This member is required.
	ExternalId *string

	// Specifies the Amazon Resource Name (ARN) of the IAM role that controls your
	// users' access to your Amazon S3 bucket or EFS file system. The policies attached
	// to this role determine the level of access that you want to provide your users
	// when transferring files into and out of your Amazon S3 bucket or EFS file
	// system. The IAM role should also contain a trust relationship that allows the
	// server to access your resources when servicing your users' transfer requests.
	//
	// This member is required.
	Role *string

	// A system-assigned unique identifier for a server instance. This is the specific
	// server that you added your user to.
	//
	// This member is required.
	ServerId *string

	// The landing directory (folder) for a user when they log in to the server using
	// the client. A HomeDirectory example is /bucket_name/home/mydirectory.
	HomeDirectory *string

	// Logical directory mappings that specify what Amazon S3 or Amazon EFS paths and
	// keys should be visible to your user and how you want to make them visible. You
	// must specify the Entry and Target pair, where Entry shows how the path is made
	// visible and Target is the actual Amazon S3 or Amazon EFS path. If you only
	// specify a target, it is displayed as is. You also must ensure that your Amazon
	// Web Services Identity and Access Management (IAM) role provides access to paths
	// in Target. This value can only be set when HomeDirectoryType is set to LOGICAL.
	// The following is an Entry and Target pair example. [ { "Entry":
	// "your-personal-report.pdf", "Target":
	// "/bucket3/customized-reports/${transfer:UserName}.pdf" } ] In most cases, you
	// can use this value instead of the scope-down policy to lock down your user to
	// the designated home directory ("chroot"). To do this, you can set Entry to / and
	// set Target to the HomeDirectory parameter value. The following is an Entry and
	// Target pair example for chroot. [ { "Entry:": "/", "Target":
	// "/bucket_name/home/mydirectory" } ] If the target of a logical directory entry
	// does not exist in Amazon S3 or EFS, the entry is ignored. As a workaround, you
	// can use the Amazon S3 API or EFS API to create 0 byte objects as place holders
	// for your directory. If using the CLI, use the s3api or efsapi call instead of s3
	// or efs so you can use the put-object operation. For example, you use the
	// following: aws s3api put-object --bucket bucketname --key path/to/folder/. Make
	// sure that the end of the key name ends in a / for it to be considered a folder.
	HomeDirectoryMappings []types.HomeDirectoryMapEntry

	// The type of landing directory (folder) you want your users' home directory to be
	// when they log into the server. If you set it to PATH, the user will see the
	// absolute Amazon S3 bucket or EFS paths as is in their file transfer protocol
	// clients. If you set it LOGICAL, you will need to provide mappings in the
	// HomeDirectoryMappings for how you want to make Amazon S3 or EFS paths visible to
	// your users.
	HomeDirectoryType types.HomeDirectoryType

	// A scope-down policy for your user so that you can use the same IAM role across
	// multiple users. This policy scopes down user access to portions of their Amazon
	// S3 bucket. Variables that you can use inside this policy include
	// ${Transfer:UserName}, ${Transfer:HomeDirectory}, and ${Transfer:HomeBucket}.
	// This only applies when domain of ServerId is S3. Amazon EFS does not use
	// scope-down policies. For scope-down policies, Amazon Web Services Transfer
	// Family stores the policy as a JSON blob, instead of the Amazon Resource Name
	// (ARN) of the policy. You save the policy as a JSON blob and pass it in the
	// Policy argument. For an example of a scope-down policy, see Example scope-down
	// policy
	// (https://docs.aws.amazon.com/transfer/latest/userguide/scope-down-policy.html).
	// For more information, see AssumeRole
	// (https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) in the
	// Amazon Web Services Security Token Service API Reference.
	Policy *string

	// The full POSIX identity, including user ID (Uid), group ID (Gid), and any
	// secondary groups IDs (SecondaryGids), that controls your users' access to your
	// Amazon EFS file systems. The POSIX permissions that are set on files and
	// directories in your file system determine the level of access your users get
	// when transferring files into and out of your Amazon EFS file systems.
	PosixProfile *types.PosixProfile
}

type CreateAccessOutput struct {

	// The external ID of the group whose users have access to your Amazon S3 or Amazon
	// EFS resources over the enabled protocols using Amazon Web Services Transfer
	// Family.
	//
	// This member is required.
	ExternalId *string

	// The ID of the server that the user is attached to.
	//
	// This member is required.
	ServerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAccess{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccess(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "CreateAccess",
	}
}
