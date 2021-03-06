// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the default retention policy details for the specified organization.
func (c *Client) GetDefaultRetentionPolicy(ctx context.Context, params *GetDefaultRetentionPolicyInput, optFns ...func(*Options)) (*GetDefaultRetentionPolicyOutput, error) {
	if params == nil {
		params = &GetDefaultRetentionPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDefaultRetentionPolicy", params, optFns, addOperationGetDefaultRetentionPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDefaultRetentionPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDefaultRetentionPolicyInput struct {

	// The organization ID.
	//
	// This member is required.
	OrganizationId *string
}

type GetDefaultRetentionPolicyOutput struct {

	// The retention policy description.
	Description *string

	// The retention policy folder configurations.
	FolderConfigurations []*types.FolderConfiguration

	// The retention policy ID.
	Id *string

	// The retention policy name.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDefaultRetentionPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDefaultRetentionPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDefaultRetentionPolicy{}, middleware.After)
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
	addOpGetDefaultRetentionPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDefaultRetentionPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDefaultRetentionPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "GetDefaultRetentionPolicy",
	}
}
