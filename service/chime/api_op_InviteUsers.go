// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends email to a maximum of 50 users, inviting them to the specified Amazon
// Chime Team account. Only Team account types are currently supported for this
// action.
func (c *Client) InviteUsers(ctx context.Context, params *InviteUsersInput, optFns ...func(*Options)) (*InviteUsersOutput, error) {
	if params == nil {
		params = &InviteUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InviteUsers", params, optFns, addOperationInviteUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InviteUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InviteUsersInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The user email addresses to which to send the email invitation.
	//
	// This member is required.
	UserEmailList []*string

	// The user type.
	UserType types.UserType
}

type InviteUsersOutput struct {

	// The email invitation details.
	Invites []*types.Invite

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationInviteUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInviteUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInviteUsers{}, middleware.After)
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
	addOpInviteUsersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInviteUsers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opInviteUsers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "InviteUsers",
	}
}
