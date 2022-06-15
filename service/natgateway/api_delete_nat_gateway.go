// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package natgateway

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDeleteNatGatewayCommon = "DeleteNatGateway"

// DeleteNatGatewayCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DeleteNatGatewayCommon operation. The "output" return
// value will be populated with the DeleteNatGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteNatGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteNatGatewayCommon Send returns without error.
//
// See DeleteNatGatewayCommon for more information on using the DeleteNatGatewayCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteNatGatewayCommonRequest method.
//    req, resp := client.DeleteNatGatewayCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NATGATEWAY) DeleteNatGatewayCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteNatGatewayCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteNatGatewayCommon API operation for NATGATEWAY.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for NATGATEWAY's
// API operation DeleteNatGatewayCommon for usage and error information.
func (c *NATGATEWAY) DeleteNatGatewayCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteNatGatewayCommonRequest(input)
	return out, req.Send()
}

// DeleteNatGatewayCommonWithContext is the same as DeleteNatGatewayCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteNatGatewayCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) DeleteNatGatewayCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteNatGatewayCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteNatGateway = "DeleteNatGateway"

// DeleteNatGatewayRequest generates a "volcstack/request.Request" representing the
// client's request for the DeleteNatGateway operation. The "output" return
// value will be populated with the DeleteNatGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteNatGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteNatGatewayCommon Send returns without error.
//
// See DeleteNatGateway for more information on using the DeleteNatGateway
// API call, and error handling.
//
//    // Example sending a request using the DeleteNatGatewayRequest method.
//    req, resp := client.DeleteNatGatewayRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NATGATEWAY) DeleteNatGatewayRequest(input *DeleteNatGatewayInput) (req *request.Request, output *DeleteNatGatewayOutput) {
	op := &request.Operation{
		Name:       opDeleteNatGateway,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNatGatewayInput{}
	}

	output = &DeleteNatGatewayOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteNatGateway API operation for NATGATEWAY.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for NATGATEWAY's
// API operation DeleteNatGateway for usage and error information.
func (c *NATGATEWAY) DeleteNatGateway(input *DeleteNatGatewayInput) (*DeleteNatGatewayOutput, error) {
	req, out := c.DeleteNatGatewayRequest(input)
	return out, req.Send()
}

// DeleteNatGatewayWithContext is the same as DeleteNatGateway with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteNatGateway for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) DeleteNatGatewayWithContext(ctx volcstack.Context, input *DeleteNatGatewayInput, opts ...request.Option) (*DeleteNatGatewayOutput, error) {
	req, out := c.DeleteNatGatewayRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteNatGatewayInput struct {
	_ struct{} `type:"structure"`

	// NatGatewayId is a required field
	NatGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteNatGatewayInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteNatGatewayInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNatGatewayInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteNatGatewayInput"}
	if s.NatGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("NatGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetNatGatewayId sets the NatGatewayId field's value.
func (s *DeleteNatGatewayInput) SetNatGatewayId(v string) *DeleteNatGatewayInput {
	s.NatGatewayId = &v
	return s
}

type DeleteNatGatewayOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteNatGatewayOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteNatGatewayOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteNatGatewayOutput) SetRequestId(v string) *DeleteNatGatewayOutput {
	s.RequestId = &v
	return s
}