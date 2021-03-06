// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a specific database from a stopped state in Amazon Lightsail. To restart
// a database, use the reboot relational database operation. The start relational
// database operation supports tag-based access control via resource tags applied
// to the resource identified by relationalDatabaseName. For more information, see
// the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) StartRelationalDatabase(ctx context.Context, params *StartRelationalDatabaseInput, optFns ...func(*Options)) (*StartRelationalDatabaseOutput, error) {
	if params == nil {
		params = &StartRelationalDatabaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartRelationalDatabase", params, optFns, addOperationStartRelationalDatabaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartRelationalDatabaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartRelationalDatabaseInput struct {

	// The name of your database to start.
	//
	// This member is required.
	RelationalDatabaseName *string
}

type StartRelationalDatabaseOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartRelationalDatabaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartRelationalDatabase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartRelationalDatabase{}, middleware.After)
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
	addOpStartRelationalDatabaseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartRelationalDatabase(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartRelationalDatabase(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "StartRelationalDatabase",
	}
}
