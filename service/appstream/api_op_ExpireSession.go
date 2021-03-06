// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Immediately stops the specified streaming session.
func (c *Client) ExpireSession(ctx context.Context, params *ExpireSessionInput, optFns ...func(*Options)) (*ExpireSessionOutput, error) {
	if params == nil {
		params = &ExpireSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExpireSession", params, optFns, addOperationExpireSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExpireSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExpireSessionInput struct {

	// The identifier of the streaming session.
	//
	// This member is required.
	SessionId *string
}

type ExpireSessionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationExpireSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpExpireSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpExpireSession{}, middleware.After)
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
	addOpExpireSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExpireSession(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opExpireSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "ExpireSession",
	}
}
