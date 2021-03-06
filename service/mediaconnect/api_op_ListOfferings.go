// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays a list of all offerings that are available to this account in the
// current AWS Region. If you have an active reservation (which means you've
// purchased an offering that has already started and hasn't expired yet), your
// account isn't eligible for other offerings.
func (c *Client) ListOfferings(ctx context.Context, params *ListOfferingsInput, optFns ...func(*Options)) (*ListOfferingsOutput, error) {
	if params == nil {
		params = &ListOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOfferings", params, optFns, addOperationListOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOfferingsInput struct {

	// The maximum number of results to return per API request. For example, you submit
	// a ListOfferings request with MaxResults set at 5. Although 20 items match your
	// request, the service returns no more than the first 5 items. (The service also
	// returns a NextToken value that you can use to fetch the next batch of results.)
	// The service might return fewer results than the MaxResults value. If MaxResults
	// is not included in the request, the service defaults to pagination with a
	// maximum of 10 results per page.
	MaxResults *int32

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListOfferings request with MaxResults set at 5. The
	// service returns the first batch of results (up to 5) and a NextToken value. To
	// see the next batch of results, you can submit the ListOfferings request a second
	// time and specify the NextToken value.
	NextToken *string
}

type ListOfferingsOutput struct {

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListOfferings request with MaxResults set at 5. The
	// service returns the first batch of results (up to 5) and a NextToken value. To
	// see the next batch of results, you can submit the ListOfferings request a second
	// time and specify the NextToken value.
	NextToken *string

	// A list of offerings that are available to this account in the current AWS
	// Region.
	Offerings []*types.Offering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOfferings{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListOfferings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListOfferings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "ListOfferings",
	}
}
