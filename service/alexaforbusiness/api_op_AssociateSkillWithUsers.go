// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Makes a private skill available for enrolled users to enable on their devices.
func (c *Client) AssociateSkillWithUsers(ctx context.Context, params *AssociateSkillWithUsersInput, optFns ...func(*Options)) (*AssociateSkillWithUsersOutput, error) {
	if params == nil {
		params = &AssociateSkillWithUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateSkillWithUsers", params, optFns, addOperationAssociateSkillWithUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateSkillWithUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateSkillWithUsersInput struct {

	// The private skill ID you want to make available to enrolled users.
	//
	// This member is required.
	SkillId *string
}

type AssociateSkillWithUsersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateSkillWithUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateSkillWithUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateSkillWithUsers{}, middleware.After)
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
	addOpAssociateSkillWithUsersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateSkillWithUsers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAssociateSkillWithUsers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "AssociateSkillWithUsers",
	}
}
