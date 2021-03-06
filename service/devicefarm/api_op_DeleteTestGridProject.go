// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a Selenium testing project and all content generated under it. You
// cannot undo this operation. You cannot delete a project if it has active
// sessions.
func (c *Client) DeleteTestGridProject(ctx context.Context, params *DeleteTestGridProjectInput, optFns ...func(*Options)) (*DeleteTestGridProjectOutput, error) {
	if params == nil {
		params = &DeleteTestGridProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTestGridProject", params, optFns, addOperationDeleteTestGridProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTestGridProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTestGridProjectInput struct {

	// The ARN of the project to delete, from CreateTestGridProject or
	// ListTestGridProjects.
	//
	// This member is required.
	ProjectArn *string
}

type DeleteTestGridProjectOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTestGridProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTestGridProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTestGridProject{}, middleware.After)
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
	addOpDeleteTestGridProjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTestGridProject(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteTestGridProject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "DeleteTestGridProject",
	}
}
