// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists exports as specified by ID. All continuous exports associated with your
// user account can be listed if you call DescribeContinuousExports as is without
// passing any parameters.
func (c *Client) DescribeContinuousExports(ctx context.Context, params *DescribeContinuousExportsInput, optFns ...func(*Options)) (*DescribeContinuousExportsOutput, error) {
	if params == nil {
		params = &DescribeContinuousExportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeContinuousExports", params, optFns, addOperationDescribeContinuousExportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeContinuousExportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeContinuousExportsInput struct {

	// The unique IDs assigned to the exports.
	ExportIds []*string

	// A number between 1 and 100 specifying the maximum number of continuous export
	// descriptions returned.
	MaxResults *int32

	// The token from the previous call to DescribeExportTasks.
	NextToken *string
}

type DescribeContinuousExportsOutput struct {

	// A list of continuous export descriptions.
	Descriptions []*types.ContinuousExportDescription

	// The token from the previous call to DescribeExportTasks.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeContinuousExportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeContinuousExports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeContinuousExports{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeContinuousExports(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeContinuousExports(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "DescribeContinuousExports",
	}
}
