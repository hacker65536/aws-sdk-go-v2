// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the specified proxy session details for the specified Amazon Chime Voice
// Connector.
func (c *Client) GetProxySession(ctx context.Context, params *GetProxySessionInput, optFns ...func(*Options)) (*GetProxySessionOutput, error) {
	if params == nil {
		params = &GetProxySessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetProxySession", params, optFns, addOperationGetProxySessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetProxySessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetProxySessionInput struct {

	// The proxy session ID.
	//
	// This member is required.
	ProxySessionId *string

	// The Amazon Chime voice connector ID.
	//
	// This member is required.
	VoiceConnectorId *string
}

type GetProxySessionOutput struct {

	// The proxy session details.
	ProxySession *types.ProxySession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetProxySessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetProxySession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetProxySession{}, middleware.After)
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
	addOpGetProxySessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetProxySession(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetProxySession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetProxySession",
	}
}
