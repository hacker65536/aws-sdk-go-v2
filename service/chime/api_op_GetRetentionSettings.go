// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets the retention settings for the specified Amazon Chime Enterprise account.
// For more information about retention settings, see Managing Chat Retention
// Policies (https://docs.aws.amazon.com/chime/latest/ag/chat-retention.html) in
// the Amazon Chime Administration Guide.
func (c *Client) GetRetentionSettings(ctx context.Context, params *GetRetentionSettingsInput, optFns ...func(*Options)) (*GetRetentionSettingsOutput, error) {
	if params == nil {
		params = &GetRetentionSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRetentionSettings", params, optFns, addOperationGetRetentionSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRetentionSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRetentionSettingsInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string
}

type GetRetentionSettingsOutput struct {

	// The timestamp representing the time at which the specified items are permanently
	// deleted, in ISO 8601 format.
	InitiateDeletionTimestamp *time.Time

	// The retention settings.
	RetentionSettings *types.RetentionSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRetentionSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRetentionSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRetentionSettings{}, middleware.After)
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
	addOpGetRetentionSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRetentionSettings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetRetentionSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetRetentionSettings",
	}
}
