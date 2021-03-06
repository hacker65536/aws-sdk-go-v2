// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates tags for the specified Auto Scaling group. When you specify a
// tag with a key that already exists, the operation overwrites the previous tag
// definition, and you do not get an error message. For more information, see
// Tagging Auto Scaling Groups and Instances
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-tagging.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) CreateOrUpdateTags(ctx context.Context, params *CreateOrUpdateTagsInput, optFns ...func(*Options)) (*CreateOrUpdateTagsOutput, error) {
	if params == nil {
		params = &CreateOrUpdateTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOrUpdateTags", params, optFns, addOperationCreateOrUpdateTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOrUpdateTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOrUpdateTagsInput struct {

	// One or more tags.
	//
	// This member is required.
	Tags []*types.Tag
}

type CreateOrUpdateTagsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateOrUpdateTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateOrUpdateTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateOrUpdateTags{}, middleware.After)
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
	addOpCreateOrUpdateTagsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOrUpdateTags(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateOrUpdateTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "CreateOrUpdateTags",
	}
}
