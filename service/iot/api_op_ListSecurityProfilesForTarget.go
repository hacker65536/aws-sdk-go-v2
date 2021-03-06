// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the Device Defender security profiles attached to a target (thing group).
func (c *Client) ListSecurityProfilesForTarget(ctx context.Context, params *ListSecurityProfilesForTargetInput, optFns ...func(*Options)) (*ListSecurityProfilesForTargetOutput, error) {
	if params == nil {
		params = &ListSecurityProfilesForTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecurityProfilesForTarget", params, optFns, addOperationListSecurityProfilesForTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecurityProfilesForTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityProfilesForTargetInput struct {

	// The ARN of the target (thing group) whose attached security profiles you want to
	// get.
	//
	// This member is required.
	SecurityProfileTargetArn *string

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// If true, return child groups too.
	Recursive *bool
}

type ListSecurityProfilesForTargetOutput struct {

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// A list of security profiles and their associated targets.
	SecurityProfileTargetMappings []*types.SecurityProfileTargetMapping

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSecurityProfilesForTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSecurityProfilesForTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSecurityProfilesForTarget{}, middleware.After)
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
	addOpListSecurityProfilesForTargetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityProfilesForTarget(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListSecurityProfilesForTarget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListSecurityProfilesForTarget",
	}
}
