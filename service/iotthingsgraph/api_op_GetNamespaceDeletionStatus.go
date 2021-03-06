// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the status of a namespace deletion task.
func (c *Client) GetNamespaceDeletionStatus(ctx context.Context, params *GetNamespaceDeletionStatusInput, optFns ...func(*Options)) (*GetNamespaceDeletionStatusOutput, error) {
	if params == nil {
		params = &GetNamespaceDeletionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetNamespaceDeletionStatus", params, optFns, addOperationGetNamespaceDeletionStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetNamespaceDeletionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetNamespaceDeletionStatusInput struct {
}

type GetNamespaceDeletionStatusOutput struct {

	// An error code returned by the namespace deletion task.
	ErrorCode types.NamespaceDeletionStatusErrorCodes

	// An error code returned by the namespace deletion task.
	ErrorMessage *string

	// The ARN of the namespace that is being deleted.
	NamespaceArn *string

	// The name of the namespace that is being deleted.
	NamespaceName *string

	// The status of the deletion request.
	Status types.NamespaceDeletionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetNamespaceDeletionStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetNamespaceDeletionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetNamespaceDeletionStatus{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetNamespaceDeletionStatus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetNamespaceDeletionStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "GetNamespaceDeletionStatus",
	}
}
