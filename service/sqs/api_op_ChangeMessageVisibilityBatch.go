// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the visibility timeout of multiple messages. This is a batch version of
// ChangeMessageVisibility. The result of the action on each message is reported
// individually in the response. You can send up to 10 ChangeMessageVisibility
// requests with each ChangeMessageVisibilityBatch action. Because the batch
// request can result in a combination of successful and unsuccessful actions, you
// should check for batch errors even when the call returns an HTTP status code of
// 200. Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example, a
// parameter list with two elements looks like this: &AttributeName.1=first
//
// &AttributeName.2=second
func (c *Client) ChangeMessageVisibilityBatch(ctx context.Context, params *ChangeMessageVisibilityBatchInput, optFns ...func(*Options)) (*ChangeMessageVisibilityBatchOutput, error) {
	if params == nil {
		params = &ChangeMessageVisibilityBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChangeMessageVisibilityBatch", params, optFns, addOperationChangeMessageVisibilityBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ChangeMessageVisibilityBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ChangeMessageVisibilityBatchInput struct {

	// A list of receipt handles of the messages for which the visibility timeout must
	// be changed.
	//
	// This member is required.
	Entries []*types.ChangeMessageVisibilityBatchRequestEntry

	// The URL of the Amazon SQS queue whose messages' visibility is changed. Queue
	// URLs and names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string
}

// For each message in the batch, the response contains a
// ChangeMessageVisibilityBatchResultEntry tag if the message succeeds or a
// BatchResultErrorEntry tag if the message fails.
type ChangeMessageVisibilityBatchOutput struct {

	// A list of BatchResultErrorEntry items.
	//
	// This member is required.
	Failed []*types.BatchResultErrorEntry

	// A list of ChangeMessageVisibilityBatchResultEntry items.
	//
	// This member is required.
	Successful []*types.ChangeMessageVisibilityBatchResultEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationChangeMessageVisibilityBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpChangeMessageVisibilityBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpChangeMessageVisibilityBatch{}, middleware.After)
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
	addOpChangeMessageVisibilityBatchValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opChangeMessageVisibilityBatch(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opChangeMessageVisibilityBatch(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "ChangeMessageVisibilityBatch",
	}
}
