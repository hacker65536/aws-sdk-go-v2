// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a build. This operation permanently deletes the build resource and any
// uploaded build files. Deleting a build does not affect the status of any active
// fleets using the build, but you can no longer create new fleets with the deleted
// build. To delete a build, specify the build ID. Learn more  Upload a Custom
// Server Build
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-intro.html)
// Related operations
//
//     * CreateBuild
//
//     * ListBuilds
//
//     * DescribeBuild
//
//
// * UpdateBuild
//
//     * DeleteBuild
func (c *Client) DeleteBuild(ctx context.Context, params *DeleteBuildInput, optFns ...func(*Options)) (*DeleteBuildOutput, error) {
	if params == nil {
		params = &DeleteBuildInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBuild", params, optFns, addOperationDeleteBuildMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBuildOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DeleteBuildInput struct {

	// A unique identifier for a build to delete. You can use either the build ID or
	// ARN value.
	//
	// This member is required.
	BuildId *string
}

type DeleteBuildOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBuildMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBuild{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBuild{}, middleware.After)
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
	addOpDeleteBuildValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBuild(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteBuild(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteBuild",
	}
}
