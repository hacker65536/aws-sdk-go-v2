// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Deletes the lifecycle policy associated with the specified repository.
func (c *Client) DeleteLifecyclePolicy(ctx context.Context, params *DeleteLifecyclePolicyInput, optFns ...func(*Options)) (*DeleteLifecyclePolicyOutput, error) {
	if params == nil {
		params = &DeleteLifecyclePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteLifecyclePolicy", params, optFns, addOperationDeleteLifecyclePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteLifecyclePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLifecyclePolicyInput struct {

	// The name of the repository.
	//
	// This member is required.
	RepositoryName *string

	// The AWS account ID associated with the registry that contains the repository. If
	// you do not specify a registry, the default registry is assumed.
	RegistryId *string
}

type DeleteLifecyclePolicyOutput struct {

	// The time stamp of the last time that the lifecycle policy was run.
	LastEvaluatedAt *time.Time

	// The JSON lifecycle policy text.
	LifecyclePolicyText *string

	// The registry ID associated with the request.
	RegistryId *string

	// The repository name associated with the request.
	RepositoryName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteLifecyclePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteLifecyclePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteLifecyclePolicy{}, middleware.After)
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
	addOpDeleteLifecyclePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLifecyclePolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteLifecyclePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "DeleteLifecyclePolicy",
	}
}
