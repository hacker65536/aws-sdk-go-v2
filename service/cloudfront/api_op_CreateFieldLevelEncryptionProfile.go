// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a field-level encryption profile.
func (c *Client) CreateFieldLevelEncryptionProfile(ctx context.Context, params *CreateFieldLevelEncryptionProfileInput, optFns ...func(*Options)) (*CreateFieldLevelEncryptionProfileOutput, error) {
	if params == nil {
		params = &CreateFieldLevelEncryptionProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFieldLevelEncryptionProfile", params, optFns, addOperationCreateFieldLevelEncryptionProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFieldLevelEncryptionProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFieldLevelEncryptionProfileInput struct {

	// The request to create a field-level encryption profile.
	//
	// This member is required.
	FieldLevelEncryptionProfileConfig *types.FieldLevelEncryptionProfileConfig
}

type CreateFieldLevelEncryptionProfileOutput struct {

	// The current version of the field level encryption profile. For example:
	// E2QWRUHAPOMQZL.
	ETag *string

	// Returned when you create a new field-level encryption profile.
	FieldLevelEncryptionProfile *types.FieldLevelEncryptionProfile

	// The fully qualified URI of the new profile resource just created.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFieldLevelEncryptionProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateFieldLevelEncryptionProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateFieldLevelEncryptionProfile{}, middleware.After)
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
	addOpCreateFieldLevelEncryptionProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFieldLevelEncryptionProfile(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateFieldLevelEncryptionProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateFieldLevelEncryptionProfile",
	}
}
