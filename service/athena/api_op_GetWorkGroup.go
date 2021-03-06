// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the workgroup with the specified name.
func (c *Client) GetWorkGroup(ctx context.Context, params *GetWorkGroupInput, optFns ...func(*Options)) (*GetWorkGroupOutput, error) {
	if params == nil {
		params = &GetWorkGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkGroup", params, optFns, addOperationGetWorkGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkGroupInput struct {

	// The name of the workgroup.
	//
	// This member is required.
	WorkGroup *string
}

type GetWorkGroupOutput struct {

	// Information about the workgroup.
	WorkGroup *types.WorkGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetWorkGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetWorkGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetWorkGroup{}, middleware.After)
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
	addOpGetWorkGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetWorkGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "GetWorkGroup",
	}
}
