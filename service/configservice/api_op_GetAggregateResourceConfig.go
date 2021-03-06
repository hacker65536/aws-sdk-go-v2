// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns configuration item that is aggregated for your specific resource in a
// specific source account and region.
func (c *Client) GetAggregateResourceConfig(ctx context.Context, params *GetAggregateResourceConfigInput, optFns ...func(*Options)) (*GetAggregateResourceConfigOutput, error) {
	if params == nil {
		params = &GetAggregateResourceConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAggregateResourceConfig", params, optFns, addOperationGetAggregateResourceConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAggregateResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAggregateResourceConfigInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// An object that identifies aggregate resource.
	//
	// This member is required.
	ResourceIdentifier *types.AggregateResourceIdentifier
}

type GetAggregateResourceConfigOutput struct {

	// Returns a ConfigurationItem object.
	ConfigurationItem *types.ConfigurationItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAggregateResourceConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAggregateResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAggregateResourceConfig{}, middleware.After)
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
	addOpGetAggregateResourceConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAggregateResourceConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAggregateResourceConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetAggregateResourceConfig",
	}
}
