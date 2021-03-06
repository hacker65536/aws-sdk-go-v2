// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a registry.
func (c *Client) CreateRegistry(ctx context.Context, params *CreateRegistryInput, optFns ...func(*Options)) (*CreateRegistryOutput, error) {
	if params == nil {
		params = &CreateRegistryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRegistry", params, optFns, addOperationCreateRegistryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRegistryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRegistryInput struct {

	// The name of the registry.
	//
	// This member is required.
	RegistryName *string

	// A description of the registry to be created.
	Description *string

	// Tags to associate with the registry.
	Tags map[string]*string
}

type CreateRegistryOutput struct {

	// The description of the registry.
	Description *string

	// The ARN of the registry.
	RegistryArn *string

	// The name of the registry.
	RegistryName *string

	// Tags associated with the registry.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRegistryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRegistry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRegistry{}, middleware.After)
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
	addOpCreateRegistryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRegistry(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateRegistry(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "CreateRegistry",
	}
}
