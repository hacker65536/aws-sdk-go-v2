// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists resources managed using Systems Manager inventory.
func (c *Client) ListResourceInventory(ctx context.Context, params *ListResourceInventoryInput, optFns ...func(*Options)) (*ListResourceInventoryOutput, error) {
	if params == nil {
		params = &ListResourceInventoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceInventory", params, optFns, addOperationListResourceInventoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceInventoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceInventoryInput struct {

	// Filters to scope the results. The following filters and logical operators are
	// supported:
	//
	//     * account_id - The ID of the AWS account that owns the resource.
	// Logical operators are EQUALS | NOT_EQUALS.
	//
	//     * application_name - The name of
	// the application. Logical operators are EQUALS | BEGINS_WITH.
	//
	//     *
	// license_included - The type of license included. Logical operators are EQUALS |
	// NOT_EQUALS. Possible values are sql-server-enterprise | sql-server-standard |
	// sql-server-web | windows-server-datacenter.
	//
	//     * platform - The platform of
	// the resource. Logical operators are EQUALS | BEGINS_WITH.
	//
	//     * resource_id -
	// The ID of the resource. Logical operators are EQUALS | NOT_EQUALS.
	Filters []*types.InventoryFilter

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string
}

type ListResourceInventoryOutput struct {

	// Token for the next set of results.
	NextToken *string

	// Information about the resources.
	ResourceInventoryList []*types.ResourceInventory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListResourceInventoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourceInventory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourceInventory{}, middleware.After)
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
	addOpListResourceInventoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceInventory(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListResourceInventory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "ListResourceInventory",
	}
}
