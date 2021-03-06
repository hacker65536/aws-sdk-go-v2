// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a package for use with Amazon ES domains.
func (c *Client) CreatePackage(ctx context.Context, params *CreatePackageInput, optFns ...func(*Options)) (*CreatePackageOutput, error) {
	if params == nil {
		params = &CreatePackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePackage", params, optFns, addOperationCreatePackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to CreatePackage operation.
type CreatePackageInput struct {

	// Unique identifier for the package.
	//
	// This member is required.
	PackageName *string

	// The customer S3 location PackageSource for importing the package.
	//
	// This member is required.
	PackageSource *types.PackageSource

	// Type of package. Currently supports only TXT-DICTIONARY.
	//
	// This member is required.
	PackageType types.PackageType

	// Description of the package.
	PackageDescription *string
}

// Container for response returned by CreatePackage operation.
type CreatePackageOutput struct {

	// Information about the package PackageDetails.
	PackageDetails *types.PackageDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePackage{}, middleware.After)
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
	addOpCreatePackageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePackage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreatePackage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "CreatePackage",
	}
}
