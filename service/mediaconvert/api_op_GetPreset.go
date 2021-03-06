// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve the JSON for a specific preset.
func (c *Client) GetPreset(ctx context.Context, params *GetPresetInput, optFns ...func(*Options)) (*GetPresetOutput, error) {
	if params == nil {
		params = &GetPresetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPreset", params, optFns, addOperationGetPresetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPresetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPresetInput struct {

	// The name of the preset.
	//
	// This member is required.
	Name *string
}

type GetPresetOutput struct {

	// A preset is a collection of preconfigured media conversion settings that you
	// want MediaConvert to apply to the output during the conversion process.
	Preset *types.Preset

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPresetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPreset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPreset{}, middleware.After)
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
	addOpGetPresetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPreset(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetPreset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "GetPreset",
	}
}
