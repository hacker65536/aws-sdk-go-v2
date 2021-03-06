// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Copies the specified DB parameter group.
func (c *Client) CopyDBParameterGroup(ctx context.Context, params *CopyDBParameterGroupInput, optFns ...func(*Options)) (*CopyDBParameterGroupOutput, error) {
	if params == nil {
		params = &CopyDBParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyDBParameterGroup", params, optFns, addOperationCopyDBParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyDBParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CopyDBParameterGroupInput struct {

	// The identifier or ARN for the source DB parameter group. For information about
	// creating an ARN, see  Constructing an ARN for Amazon RDS
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing)
	// in the Amazon RDS User Guide. Constraints:
	//
	//     * Must specify a valid DB
	// parameter group.
	//
	//     * Must specify a valid DB parameter group identifier, for
	// example my-db-param-group, or a valid ARN.
	//
	// This member is required.
	SourceDBParameterGroupIdentifier *string

	// A description for the copied DB parameter group.
	//
	// This member is required.
	TargetDBParameterGroupDescription *string

	// The identifier for the copied DB parameter group. Constraints:
	//
	//     * Can't be
	// null, empty, or blank
	//
	//     * Must contain from 1 to 255 letters, numbers, or
	// hyphens
	//
	//     * First character must be a letter
	//
	//     * Can't end with a hyphen
	// or contain two consecutive hyphens
	//
	// Example: my-db-parameter-group
	//
	// This member is required.
	TargetDBParameterGroupIdentifier *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []*types.Tag
}

type CopyDBParameterGroupOutput struct {

	// Contains the details of an Amazon RDS DB parameter group. This data type is used
	// as a response element in the DescribeDBParameterGroups action.
	DBParameterGroup *types.DBParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCopyDBParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopyDBParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyDBParameterGroup{}, middleware.After)
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
	addOpCopyDBParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDBParameterGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCopyDBParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CopyDBParameterGroup",
	}
}
