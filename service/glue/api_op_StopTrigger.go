// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a specified trigger.
func (c *Client) StopTrigger(ctx context.Context, params *StopTriggerInput, optFns ...func(*Options)) (*StopTriggerOutput, error) {
	if params == nil {
		params = &StopTriggerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopTrigger", params, optFns, addOperationStopTriggerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopTriggerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopTriggerInput struct {

	// The name of the trigger to stop.
	//
	// This member is required.
	Name *string
}

type StopTriggerOutput struct {

	// The name of the trigger that was stopped.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopTriggerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopTrigger{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopTrigger{}, middleware.After)
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
	addOpStopTriggerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopTrigger(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopTrigger(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StopTrigger",
	}
}
