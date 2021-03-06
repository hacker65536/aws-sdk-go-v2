// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the IP address ranges that were specified in calls to ProvisionByoipCidr
// (https://docs.aws.amazon.com/global-accelerator/latest/api/ProvisionByoipCidr.html),
// including the current state and a history of state changes. To see an AWS CLI
// example of listing BYOIP CIDR addresses, scroll down to Example.
func (c *Client) ListByoipCidrs(ctx context.Context, params *ListByoipCidrsInput, optFns ...func(*Options)) (*ListByoipCidrsOutput, error) {
	if params == nil {
		params = &ListByoipCidrsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListByoipCidrs", params, optFns, addOperationListByoipCidrsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListByoipCidrsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListByoipCidrsInput struct {

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type ListByoipCidrsOutput struct {

	// Information about your address ranges.
	ByoipCidrs []*types.ByoipCidr

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListByoipCidrsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListByoipCidrs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListByoipCidrs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListByoipCidrs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListByoipCidrs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "ListByoipCidrs",
	}
}
