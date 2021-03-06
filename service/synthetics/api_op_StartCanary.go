// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use this operation to run a canary that has already been created. The frequency
// of the canary runs is determined by the value of the canary's Schedule. To see a
// canary's schedule, use GetCanary
// (https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_GetCanary.html).
func (c *Client) StartCanary(ctx context.Context, params *StartCanaryInput, optFns ...func(*Options)) (*StartCanaryOutput, error) {
	if params == nil {
		params = &StartCanaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartCanary", params, optFns, addOperationStartCanaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartCanaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartCanaryInput struct {

	// The name of the canary that you want to run. To find canary names, use
	// DescribeCanaries
	// (https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DescribeCanaries.html).
	//
	// This member is required.
	Name *string
}

type StartCanaryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartCanaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartCanary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartCanary{}, middleware.After)
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
	addOpStartCanaryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartCanary(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartCanary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "StartCanary",
	}
}
