// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the number of minutes used by account.
func (c *Client) GetMinuteUsage(ctx context.Context, params *GetMinuteUsageInput, optFns ...func(*Options)) (*GetMinuteUsageOutput, error) {
	if params == nil {
		params = &GetMinuteUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMinuteUsage", params, optFns, addOperationGetMinuteUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMinuteUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type GetMinuteUsageInput struct {

	// The month being requested, with a value of 1-12.
	//
	// This member is required.
	Month *int32

	// The year being requested, in the format of YYYY.
	//
	// This member is required.
	Year *int32
}

//
type GetMinuteUsageOutput struct {

	// Estimated number of minutes remaining for an account, specific to the month
	// being requested.
	EstimatedMinutesRemaining *int32

	// Returns whether or not an account has signed up for the reserved minutes pricing
	// plan, specific to the month being requested.
	IsReservedMinutesCustomer *bool

	// Total number of reserved minutes allocated, specific to the month being
	// requested.
	TotalReservedMinuteAllocation *int32

	// Total scheduled minutes for an account, specific to the month being requested.
	TotalScheduledMinutes *int32

	// Upcoming minutes scheduled for an account, specific to the month being
	// requested.
	UpcomingMinutesScheduled *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMinuteUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMinuteUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMinuteUsage{}, middleware.After)
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
	addOpGetMinuteUsageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMinuteUsage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetMinuteUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "",
		OperationName: "GetMinuteUsage",
	}
}
