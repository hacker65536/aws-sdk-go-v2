// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified organization config rule and all of its evaluation results
// from all member accounts in that organization. Only a master account and a
// delegated administrator account can delete an organization config rule. When
// calling this API with a delegated administrator, you must ensure AWS
// Organizations ListDelegatedAdministrator permissions are added. AWS Config sets
// the state of a rule to DELETE_IN_PROGRESS until the deletion is complete. You
// cannot update a rule while it is in this state.
func (c *Client) DeleteOrganizationConfigRule(ctx context.Context, params *DeleteOrganizationConfigRuleInput, optFns ...func(*Options)) (*DeleteOrganizationConfigRuleOutput, error) {
	if params == nil {
		params = &DeleteOrganizationConfigRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteOrganizationConfigRule", params, optFns, addOperationDeleteOrganizationConfigRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteOrganizationConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOrganizationConfigRuleInput struct {

	// The name of organization config rule that you want to delete.
	//
	// This member is required.
	OrganizationConfigRuleName *string
}

type DeleteOrganizationConfigRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteOrganizationConfigRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteOrganizationConfigRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteOrganizationConfigRule{}, middleware.After)
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
	addOpDeleteOrganizationConfigRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOrganizationConfigRule(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteOrganizationConfigRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteOrganizationConfigRule",
	}
}
