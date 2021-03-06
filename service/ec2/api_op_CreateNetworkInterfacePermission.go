// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Grants an AWS-authorized account permission to attach the specified network
// interface to an instance in their account. You can grant permission to a single
// AWS account only, and only one account at a time.
func (c *Client) CreateNetworkInterfacePermission(ctx context.Context, params *CreateNetworkInterfacePermissionInput, optFns ...func(*Options)) (*CreateNetworkInterfacePermissionOutput, error) {
	if params == nil {
		params = &CreateNetworkInterfacePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNetworkInterfacePermission", params, optFns, addOperationCreateNetworkInterfacePermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNetworkInterfacePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateNetworkInterfacePermission.
type CreateNetworkInterfacePermissionInput struct {

	// The ID of the network interface.
	//
	// This member is required.
	NetworkInterfaceId *string

	// The type of permission to grant.
	//
	// This member is required.
	Permission types.InterfacePermissionType

	// The AWS account ID.
	AwsAccountId *string

	// The AWS service. Currently not supported.
	AwsService *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

// Contains the output of CreateNetworkInterfacePermission.
type CreateNetworkInterfacePermissionOutput struct {

	// Information about the permission for the network interface.
	InterfacePermission *types.NetworkInterfacePermission

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateNetworkInterfacePermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateNetworkInterfacePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateNetworkInterfacePermission{}, middleware.After)
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
	addOpCreateNetworkInterfacePermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNetworkInterfacePermission(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateNetworkInterfacePermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateNetworkInterfacePermission",
	}
}
