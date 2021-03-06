// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an event destination in a configuration set.
func (c *Client) DeleteConfigurationSetEventDestination(ctx context.Context, params *DeleteConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*DeleteConfigurationSetEventDestinationOutput, error) {
	if params == nil {
		params = &DeleteConfigurationSetEventDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConfigurationSetEventDestination", params, optFns, addOperationDeleteConfigurationSetEventDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConfigurationSetEventDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteConfigurationSetEventDestinationInput struct {

	// ConfigurationSetName
	//
	// This member is required.
	ConfigurationSetName *string

	// EventDestinationName
	//
	// This member is required.
	EventDestinationName *string
}

// An empty object that indicates that the event destination was deleted
// successfully.
type DeleteConfigurationSetEventDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteConfigurationSetEventDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteConfigurationSetEventDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteConfigurationSetEventDestination{}, middleware.After)
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
	addOpDeleteConfigurationSetEventDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConfigurationSetEventDestination(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteConfigurationSetEventDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DeleteConfigurationSetEventDestination",
	}
}
