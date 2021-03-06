// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified delivery stream and its status. For example, after your
// delivery stream is created, call DescribeDeliveryStream to see whether the
// delivery stream is ACTIVE and therefore ready for data to be sent to it. If the
// status of a delivery stream is CREATING_FAILED, this status doesn't change, and
// you can't invoke CreateDeliveryStream again on it. However, you can invoke the
// DeleteDeliveryStream operation to delete it. If the status is DELETING_FAILED,
// you can force deletion by invoking DeleteDeliveryStream again but with
// DeleteDeliveryStreamInput$AllowForceDelete set to true.
func (c *Client) DescribeDeliveryStream(ctx context.Context, params *DescribeDeliveryStreamInput, optFns ...func(*Options)) (*DescribeDeliveryStreamOutput, error) {
	if params == nil {
		params = &DescribeDeliveryStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDeliveryStream", params, optFns, addOperationDescribeDeliveryStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDeliveryStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDeliveryStreamInput struct {

	// The name of the delivery stream.
	//
	// This member is required.
	DeliveryStreamName *string

	// The ID of the destination to start returning the destination information.
	// Kinesis Data Firehose supports one destination per delivery stream.
	ExclusiveStartDestinationId *string

	// The limit on the number of destinations to return. You can have one destination
	// per delivery stream.
	Limit *int32
}

type DescribeDeliveryStreamOutput struct {

	// Information about the delivery stream.
	//
	// This member is required.
	DeliveryStreamDescription *types.DeliveryStreamDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDeliveryStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDeliveryStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDeliveryStream{}, middleware.After)
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
	addOpDescribeDeliveryStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDeliveryStream(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDeliveryStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "DescribeDeliveryStream",
	}
}
