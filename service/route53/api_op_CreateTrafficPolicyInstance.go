// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates resource record sets in a specified hosted zone based on the settings in
// a specified traffic policy version. In addition, CreateTrafficPolicyInstance
// associates the resource record sets with a specified domain name (such as
// example.com) or subdomain name (such as www.example.com). Amazon Route 53
// responds to DNS queries for the domain or subdomain name by using the resource
// record sets that CreateTrafficPolicyInstance created.
func (c *Client) CreateTrafficPolicyInstance(ctx context.Context, params *CreateTrafficPolicyInstanceInput, optFns ...func(*Options)) (*CreateTrafficPolicyInstanceOutput, error) {
	if params == nil {
		params = &CreateTrafficPolicyInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrafficPolicyInstance", params, optFns, addOperationCreateTrafficPolicyInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrafficPolicyInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the resource record sets that you
// want to create based on a specified traffic policy.
type CreateTrafficPolicyInstanceInput struct {

	// The ID of the hosted zone that you want Amazon Route 53 to create resource
	// record sets in by using the configuration in a traffic policy.
	//
	// This member is required.
	HostedZoneId *string

	// The domain name (such as example.com) or subdomain name (such as
	// www.example.com) for which Amazon Route 53 responds to DNS queries by using the
	// resource record sets that Route 53 creates for this traffic policy instance.
	//
	// This member is required.
	Name *string

	// (Optional) The TTL that you want Amazon Route 53 to assign to all of the
	// resource record sets that it creates in the specified hosted zone.
	//
	// This member is required.
	TTL *int64

	// The ID of the traffic policy that you want to use to create resource record sets
	// in the specified hosted zone.
	//
	// This member is required.
	TrafficPolicyId *string

	// The version of the traffic policy that you want to use to create resource record
	// sets in the specified hosted zone.
	//
	// This member is required.
	TrafficPolicyVersion *int32
}

// A complex type that contains the response information for the
// CreateTrafficPolicyInstance request.
type CreateTrafficPolicyInstanceOutput struct {

	// A unique URL that represents a new traffic policy instance.
	//
	// This member is required.
	Location *string

	// A complex type that contains settings for the new traffic policy instance.
	//
	// This member is required.
	TrafficPolicyInstance *types.TrafficPolicyInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTrafficPolicyInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateTrafficPolicyInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateTrafficPolicyInstance{}, middleware.After)
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
	addOpCreateTrafficPolicyInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrafficPolicyInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addSanitizeURLMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateTrafficPolicyInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "CreateTrafficPolicyInstance",
	}
}
