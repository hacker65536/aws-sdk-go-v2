// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the named query if you have access to the workgroup in which the query
// was saved. For code samples using the AWS SDK for Java, see Examples and Code
// Samples (http://docs.aws.amazon.com/athena/latest/ug/code-samples.html) in the
// Amazon Athena User Guide.
func (c *Client) DeleteNamedQuery(ctx context.Context, params *DeleteNamedQueryInput, optFns ...func(*Options)) (*DeleteNamedQueryOutput, error) {
	if params == nil {
		params = &DeleteNamedQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteNamedQuery", params, optFns, addOperationDeleteNamedQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteNamedQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteNamedQueryInput struct {

	// The unique ID of the query to delete.
	//
	// This member is required.
	NamedQueryId *string
}

type DeleteNamedQueryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteNamedQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteNamedQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteNamedQuery{}, middleware.After)
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
	addIdempotencyToken_opDeleteNamedQueryMiddleware(stack, options)
	addOpDeleteNamedQueryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteNamedQuery(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpDeleteNamedQuery struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteNamedQuery) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteNamedQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteNamedQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteNamedQueryInput ")
	}

	if input.NamedQueryId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.NamedQueryId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteNamedQueryMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpDeleteNamedQuery{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteNamedQuery(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "DeleteNamedQuery",
	}
}
