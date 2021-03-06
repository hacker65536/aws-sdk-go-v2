// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a billing group.
func (c *Client) DescribeBillingGroup(ctx context.Context, params *DescribeBillingGroupInput, optFns ...func(*Options)) (*DescribeBillingGroupOutput, error) {
	if params == nil {
		params = &DescribeBillingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBillingGroup", params, optFns, addOperationDescribeBillingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBillingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBillingGroupInput struct {

	// The name of the billing group.
	//
	// This member is required.
	BillingGroupName *string
}

type DescribeBillingGroupOutput struct {

	// The ARN of the billing group.
	BillingGroupArn *string

	// The ID of the billing group.
	BillingGroupId *string

	// Additional information about the billing group.
	BillingGroupMetadata *types.BillingGroupMetadata

	// The name of the billing group.
	BillingGroupName *string

	// The properties of the billing group.
	BillingGroupProperties *types.BillingGroupProperties

	// The version of the billing group.
	Version *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBillingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBillingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBillingGroup{}, middleware.After)
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
	addOpDescribeBillingGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBillingGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeBillingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeBillingGroup",
	}
}
