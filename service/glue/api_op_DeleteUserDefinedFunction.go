// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an existing function definition from the Data Catalog.
func (c *Client) DeleteUserDefinedFunction(ctx context.Context, params *DeleteUserDefinedFunctionInput, optFns ...func(*Options)) (*DeleteUserDefinedFunctionOutput, error) {
	if params == nil {
		params = &DeleteUserDefinedFunctionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteUserDefinedFunction", params, optFns, addOperationDeleteUserDefinedFunctionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteUserDefinedFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUserDefinedFunctionInput struct {

	// The name of the catalog database where the function is located.
	//
	// This member is required.
	DatabaseName *string

	// The name of the function definition to be deleted.
	//
	// This member is required.
	FunctionName *string

	// The ID of the Data Catalog where the function to be deleted is located. If none
	// is supplied, the AWS account ID is used by default.
	CatalogId *string
}

type DeleteUserDefinedFunctionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteUserDefinedFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteUserDefinedFunction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteUserDefinedFunction{}, middleware.After)
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
	addOpDeleteUserDefinedFunctionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUserDefinedFunction(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteUserDefinedFunction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteUserDefinedFunction",
	}
}
