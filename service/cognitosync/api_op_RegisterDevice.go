// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a device to receive push sync notifications. This API can only be
// called with temporary credentials provided by Cognito Identity. You cannot call
// this API with developer credentials.
func (c *Client) RegisterDevice(ctx context.Context, params *RegisterDeviceInput, optFns ...func(*Options)) (*RegisterDeviceOutput, error) {
	if params == nil {
		params = &RegisterDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterDevice", params, optFns, addOperationRegisterDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to RegisterDevice.
type RegisterDeviceInput struct {

	// The unique ID for this identity.
	//
	// This member is required.
	IdentityId *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. Here, the ID of the pool that the identity belongs
	// to.
	//
	// This member is required.
	IdentityPoolId *string

	// The SNS platform type (e.g. GCM, SDM, APNS, APNS_SANDBOX).
	//
	// This member is required.
	Platform types.Platform

	// The push token.
	//
	// This member is required.
	Token *string
}

// Response to a RegisterDevice request.
type RegisterDeviceOutput struct {

	// The unique ID generated for this device by Cognito.
	DeviceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterDevice{}, middleware.After)
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
	addOpRegisterDeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterDevice(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterDevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "RegisterDevice",
	}
}
