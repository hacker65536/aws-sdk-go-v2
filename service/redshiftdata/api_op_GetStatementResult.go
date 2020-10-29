// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshiftdata/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Fetches the temporarily cached result of an SQL statement. A token is returned
// to page through the statement results.
func (c *Client) GetStatementResult(ctx context.Context, params *GetStatementResultInput, optFns ...func(*Options)) (*GetStatementResultOutput, error) {
	if params == nil {
		params = &GetStatementResultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStatementResult", params, optFns, addOperationGetStatementResultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStatementResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetStatementResultInput struct {

	// The identifier of the SQL statement whose results are to be fetched. This value
	// is a universally unique identifier (UUID) generated by Amazon Redshift Data API.
	// This identifier is returned by ExecuteStatment and ListStatements.
	//
	// This member is required.
	Id *string

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string
}

type GetStatementResultOutput struct {

	// The results of the SQL statement.
	//
	// This member is required.
	Records [][]types.Field

	// The properties (metadata) of a column.
	ColumnMetadata []*types.ColumnMetadata

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// The total number of rows in the result set returned from a query. You can use
	// this number to estimate the number of calls to the GetStatementResult operation
	// needed to page through the results.
	TotalNumRows *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetStatementResultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetStatementResult{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetStatementResult{}, middleware.After)
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
	addOpGetStatementResultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetStatementResult(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetStatementResult(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-data",
		OperationName: "GetStatementResult",
	}
}