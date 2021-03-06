// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a service mesh. A service mesh is a logical boundary for network traffic
// between services that are represented by resources within the mesh. After you
// create your service mesh, you can create virtual services, virtual nodes,
// virtual routers, and routes to distribute traffic between the applications in
// your mesh. For more information about service meshes, see Service meshes
// (https://docs.aws.amazon.com/app-mesh/latest/userguide/meshes.html).
func (c *Client) CreateMesh(ctx context.Context, params *CreateMeshInput, optFns ...func(*Options)) (*CreateMeshOutput, error) {
	if params == nil {
		params = &CreateMeshInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMesh", params, optFns, addOperationCreateMeshMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMeshOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateMeshInput struct {

	// The name to use for the service mesh.
	//
	// This member is required.
	MeshName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string

	// The service mesh specification to apply.
	Spec *types.MeshSpec

	// Optional metadata that you can apply to the service mesh to assist with
	// categorization and organization. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length of
	// 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []*types.TagRef
}

//
type CreateMeshOutput struct {

	// The full description of your service mesh following the create call.
	//
	// This member is required.
	Mesh *types.MeshData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateMeshMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMesh{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMesh{}, middleware.After)
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
	addOpCreateMeshValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
