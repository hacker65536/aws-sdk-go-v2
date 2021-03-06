// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the results of the Security Hub insight specified by the insight ARN.
func (c *Client) GetInsightResults(ctx context.Context, params *GetInsightResultsInput, optFns ...func(*Options)) (*GetInsightResultsOutput, error) {
	if params == nil {
		params = &GetInsightResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInsightResults", params, optFns, addOperationGetInsightResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInsightResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInsightResultsInput struct {

	// The ARN of the insight for which to return results.
	//
	// This member is required.
	InsightArn *string
}

type GetInsightResultsOutput struct {

	// The insight results returned by the operation.
	//
	// This member is required.
	InsightResults *types.InsightResults

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetInsightResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetInsightResults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInsightResults{}, middleware.After)
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
	addOpGetInsightResultsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInsightResults(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetInsightResults(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "GetInsightResults",
	}
}
