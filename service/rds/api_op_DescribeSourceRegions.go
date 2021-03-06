// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the source AWS Regions where the current AWS Region can create
// a read replica or copy a DB snapshot from. This API action supports pagination.
func (c *Client) DescribeSourceRegions(ctx context.Context, params *DescribeSourceRegionsInput, optFns ...func(*Options)) (*DescribeSourceRegionsOutput, error) {
	if params == nil {
		params = &DescribeSourceRegionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSourceRegions", params, optFns, addOperationDescribeSourceRegionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSourceRegionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeSourceRegionsInput struct {

	// This parameter isn't currently supported.
	Filters []*types.Filter

	// An optional pagination token provided by a previous DescribeSourceRegions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The source AWS Region name. For example, us-east-1. Constraints:
	//
	//     * Must
	// specify a valid AWS Region name.
	RegionName *string
}

// Contains the result of a successful invocation of the DescribeSourceRegions
// action.
type DescribeSourceRegionsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// A list of SourceRegion instances that contains each source AWS Region that the
	// current AWS Region can get a read replica or a DB snapshot from.
	SourceRegions []*types.SourceRegion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSourceRegionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeSourceRegions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeSourceRegions{}, middleware.After)
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
	addOpDescribeSourceRegionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSourceRegions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeSourceRegions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeSourceRegions",
	}
}
