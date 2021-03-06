// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more of your linked EC2-Classic instances. This request only
// returns information about EC2-Classic instances linked to a VPC through
// ClassicLink. You cannot use this request to return information about other
// instances.
func (c *Client) DescribeClassicLinkInstances(ctx context.Context, params *DescribeClassicLinkInstancesInput, optFns ...func(*Options)) (*DescribeClassicLinkInstancesOutput, error) {
	if params == nil {
		params = &DescribeClassicLinkInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClassicLinkInstances", params, optFns, addOperationDescribeClassicLinkInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClassicLinkInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClassicLinkInstancesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	//     * group-id - The ID of a VPC security group that's
	// associated with the instance.
	//
	//     * instance-id - The ID of the instance.
	//
	//
	// * tag: - The key/value combination of a tag assigned to the resource. Use the
	// tag key in the filter name and the tag value as the filter value. For example,
	// to find all resources that have a tag with the key Owner and the value TeamA,
	// specify tag:Owner for the filter name and TeamA for the filter value.
	//
	//     *
	// tag-key - The key of a tag assigned to the resource. Use this filter to find all
	// resources assigned a tag with a specific key, regardless of the tag value.
	//
	//
	// * vpc-id - The ID of the VPC to which the instance is linked. vpc-id - The ID of
	// the VPC that the instance is linked to.
	Filters []*types.Filter

	// One or more instance IDs. Must be instances linked to a VPC through ClassicLink.
	InstanceIds []*string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	// Constraint: If the value is greater than 1000, we return only 1000 items.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type DescribeClassicLinkInstancesOutput struct {

	// Information about one or more linked EC2-Classic instances.
	Instances []*types.ClassicLinkInstance

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeClassicLinkInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeClassicLinkInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeClassicLinkInstances{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClassicLinkInstances(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeClassicLinkInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeClassicLinkInstances",
	}
}
