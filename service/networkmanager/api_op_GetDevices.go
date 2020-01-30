// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetDevicesInput struct {
	_ struct{} `type:"structure"`

	// One or more device IDs. The maximum is 10.
	DeviceIds []string `location:"querystring" locationName:"deviceIds" type:"list"`

	// The ID of the global network.
	//
	// GlobalNetworkId is a required field
	GlobalNetworkId *string `location:"uri" locationName:"globalNetworkId" type:"string" required:"true"`

	// The maximum number of results to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The ID of the site.
	SiteId *string `location:"querystring" locationName:"siteId" type:"string"`
}

// String returns the string representation
func (s GetDevicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDevicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDevicesInput"}

	if s.GlobalNetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalNetworkId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDevicesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GlobalNetworkId != nil {
		v := *s.GlobalNetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "globalNetworkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeviceIds != nil {
		v := s.DeviceIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "deviceIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SiteId != nil {
		v := *s.SiteId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "siteId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetDevicesOutput struct {
	_ struct{} `type:"structure"`

	// The devices.
	Devices []Device `type:"list"`

	// The token for the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetDevicesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDevicesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Devices != nil {
		v := s.Devices

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Devices", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetDevices = "GetDevices"

// GetDevicesRequest returns a request value for making API operation for
// AWS Network Manager.
//
// Gets information about one or more of your devices in a global network.
//
//    // Example sending a request using GetDevicesRequest.
//    req := client.GetDevicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/networkmanager-2019-07-05/GetDevices
func (c *Client) GetDevicesRequest(input *GetDevicesInput) GetDevicesRequest {
	op := &aws.Operation{
		Name:       opGetDevices,
		HTTPMethod: "GET",
		HTTPPath:   "/global-networks/{globalNetworkId}/devices",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetDevicesInput{}
	}

	req := c.newRequest(op, input, &GetDevicesOutput{})
	return GetDevicesRequest{Request: req, Input: input, Copy: c.GetDevicesRequest}
}

// GetDevicesRequest is the request type for the
// GetDevices API operation.
type GetDevicesRequest struct {
	*aws.Request
	Input *GetDevicesInput
	Copy  func(*GetDevicesInput) GetDevicesRequest
}

// Send marshals and sends the GetDevices API request.
func (r GetDevicesRequest) Send(ctx context.Context) (*GetDevicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDevicesResponse{
		GetDevicesOutput: r.Request.Data.(*GetDevicesOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetDevicesRequestPaginator returns a paginator for GetDevices.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetDevicesRequest(input)
//   p := networkmanager.NewGetDevicesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetDevicesPaginator(req GetDevicesRequest) GetDevicesPaginator {
	return GetDevicesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetDevicesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetDevicesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetDevicesPaginator struct {
	aws.Pager
}

func (p *GetDevicesPaginator) CurrentPage() *GetDevicesOutput {
	return p.Pager.CurrentPage().(*GetDevicesOutput)
}

// GetDevicesResponse is the response type for the
// GetDevices API operation.
type GetDevicesResponse struct {
	*GetDevicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDevices request.
func (r *GetDevicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}