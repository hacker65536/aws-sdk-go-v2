// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Reboots the specified WorkSpaces. You cannot reboot a WorkSpace unless its state
// is AVAILABLE or UNHEALTHY. This operation is asynchronous and returns before the
// WorkSpaces have rebooted.
func (c *Client) RebootWorkspaces(ctx context.Context, params *RebootWorkspacesInput, optFns ...func(*Options)) (*RebootWorkspacesOutput, error) {
	if params == nil {
		params = &RebootWorkspacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RebootWorkspaces", params, optFns, addOperationRebootWorkspacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RebootWorkspacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RebootWorkspacesInput struct {

	// The WorkSpaces to reboot. You can specify up to 25 WorkSpaces.
	//
	// This member is required.
	RebootWorkspaceRequests []*types.RebootRequest
}

type RebootWorkspacesOutput struct {

	// Information about the WorkSpaces that could not be rebooted.
	FailedRequests []*types.FailedWorkspaceChangeRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRebootWorkspacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRebootWorkspaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRebootWorkspaces{}, middleware.After)
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
	addOpRebootWorkspacesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRebootWorkspaces(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRebootWorkspaces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "RebootWorkspaces",
	}
}
