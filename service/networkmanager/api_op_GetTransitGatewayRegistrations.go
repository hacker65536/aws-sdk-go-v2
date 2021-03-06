// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the transit gateway registrations in a specified global
// network.
func (c *Client) GetTransitGatewayRegistrations(ctx context.Context, params *GetTransitGatewayRegistrationsInput, optFns ...func(*Options)) (*GetTransitGatewayRegistrationsOutput, error) {
	if params == nil {
		params = &GetTransitGatewayRegistrationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTransitGatewayRegistrations", params, optFns, addOperationGetTransitGatewayRegistrationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTransitGatewayRegistrationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayRegistrationsInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The Amazon Resource Names (ARNs) of one or more transit gateways. The maximum is
	// 10.
	TransitGatewayArns []*string
}

type GetTransitGatewayRegistrationsOutput struct {

	// The token for the next page of results.
	NextToken *string

	// The transit gateway registrations.
	TransitGatewayRegistrations []*types.TransitGatewayRegistration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTransitGatewayRegistrationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTransitGatewayRegistrations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTransitGatewayRegistrations{}, middleware.After)
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
	addOpGetTransitGatewayRegistrationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransitGatewayRegistrations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetTransitGatewayRegistrations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "GetTransitGatewayRegistrations",
	}
}
