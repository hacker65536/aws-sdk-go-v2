// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the configuration settings for storing data classification results.
func (c *Client) GetClassificationExportConfiguration(ctx context.Context, params *GetClassificationExportConfigurationInput, optFns ...func(*Options)) (*GetClassificationExportConfigurationOutput, error) {
	if params == nil {
		params = &GetClassificationExportConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetClassificationExportConfiguration", params, optFns, addOperationGetClassificationExportConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetClassificationExportConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetClassificationExportConfigurationInput struct {
}

type GetClassificationExportConfigurationOutput struct {

	// The location where data classification results are stored, and the encryption
	// settings that are used when storing results in that location.
	Configuration *types.ClassificationExportConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetClassificationExportConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetClassificationExportConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetClassificationExportConfiguration{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetClassificationExportConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetClassificationExportConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetClassificationExportConfiguration",
	}
}
