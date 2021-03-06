// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Restores a directory using an existing directory snapshot. When you restore a
// directory from a snapshot, any changes made to the directory after the snapshot
// date are overwritten. This action returns as soon as the restore operation is
// initiated. You can monitor the progress of the restore operation by calling the
// DescribeDirectories operation with the directory identifier. When the
// DirectoryDescription.Stage value changes to Active, the restore operation is
// complete.
func (c *Client) RestoreFromSnapshot(ctx context.Context, params *RestoreFromSnapshotInput, optFns ...func(*Options)) (*RestoreFromSnapshotOutput, error) {
	if params == nil {
		params = &RestoreFromSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreFromSnapshot", params, optFns, addOperationRestoreFromSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// An object representing the inputs for the RestoreFromSnapshot operation.
type RestoreFromSnapshotInput struct {

	// The identifier of the snapshot to restore from.
	//
	// This member is required.
	SnapshotId *string
}

// Contains the results of the RestoreFromSnapshot operation.
type RestoreFromSnapshotOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRestoreFromSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRestoreFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRestoreFromSnapshot{}, middleware.After)
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
	addOpRestoreFromSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreFromSnapshot(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRestoreFromSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "RestoreFromSnapshot",
	}
}
