// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opModifyDirectConnectVirtualInterfaceAttributesCommon = "ModifyDirectConnectVirtualInterfaceAttributes"

// ModifyDirectConnectVirtualInterfaceAttributesCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the ModifyDirectConnectVirtualInterfaceAttributesCommon operation. The "output" return
// value will be populated with the ModifyDirectConnectVirtualInterfaceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDirectConnectVirtualInterfaceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDirectConnectVirtualInterfaceAttributesCommon Send returns without error.
//
// See ModifyDirectConnectVirtualInterfaceAttributesCommon for more information on using the ModifyDirectConnectVirtualInterfaceAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDirectConnectVirtualInterfaceAttributesCommonRequest method.
//    req, resp := client.ModifyDirectConnectVirtualInterfaceAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) ModifyDirectConnectVirtualInterfaceAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDirectConnectVirtualInterfaceAttributesCommon,
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

// ModifyDirectConnectVirtualInterfaceAttributesCommon API operation for DIRECTCONNECT.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation ModifyDirectConnectVirtualInterfaceAttributesCommon for usage and error information.
func (c *DIRECTCONNECT) ModifyDirectConnectVirtualInterfaceAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDirectConnectVirtualInterfaceAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyDirectConnectVirtualInterfaceAttributesCommonWithContext is the same as ModifyDirectConnectVirtualInterfaceAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDirectConnectVirtualInterfaceAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) ModifyDirectConnectVirtualInterfaceAttributesCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDirectConnectVirtualInterfaceAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDirectConnectVirtualInterfaceAttributes = "ModifyDirectConnectVirtualInterfaceAttributes"

// ModifyDirectConnectVirtualInterfaceAttributesRequest generates a "volcstack/request.Request" representing the
// client's request for the ModifyDirectConnectVirtualInterfaceAttributes operation. The "output" return
// value will be populated with the ModifyDirectConnectVirtualInterfaceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDirectConnectVirtualInterfaceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDirectConnectVirtualInterfaceAttributesCommon Send returns without error.
//
// See ModifyDirectConnectVirtualInterfaceAttributes for more information on using the ModifyDirectConnectVirtualInterfaceAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyDirectConnectVirtualInterfaceAttributesRequest method.
//    req, resp := client.ModifyDirectConnectVirtualInterfaceAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) ModifyDirectConnectVirtualInterfaceAttributesRequest(input *ModifyDirectConnectVirtualInterfaceAttributesInput) (req *request.Request, output *ModifyDirectConnectVirtualInterfaceAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyDirectConnectVirtualInterfaceAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDirectConnectVirtualInterfaceAttributesInput{}
	}

	output = &ModifyDirectConnectVirtualInterfaceAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyDirectConnectVirtualInterfaceAttributes API operation for DIRECTCONNECT.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation ModifyDirectConnectVirtualInterfaceAttributes for usage and error information.
func (c *DIRECTCONNECT) ModifyDirectConnectVirtualInterfaceAttributes(input *ModifyDirectConnectVirtualInterfaceAttributesInput) (*ModifyDirectConnectVirtualInterfaceAttributesOutput, error) {
	req, out := c.ModifyDirectConnectVirtualInterfaceAttributesRequest(input)
	return out, req.Send()
}

// ModifyDirectConnectVirtualInterfaceAttributesWithContext is the same as ModifyDirectConnectVirtualInterfaceAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDirectConnectVirtualInterfaceAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) ModifyDirectConnectVirtualInterfaceAttributesWithContext(ctx volcstack.Context, input *ModifyDirectConnectVirtualInterfaceAttributesInput, opts ...request.Option) (*ModifyDirectConnectVirtualInterfaceAttributesOutput, error) {
	req, out := c.ModifyDirectConnectVirtualInterfaceAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDirectConnectVirtualInterfaceAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `min:"1" max:"255" type:"string"`

	// VirtualInterfaceId is a required field
	VirtualInterfaceId *string `type:"string" required:"true"`

	VirtualInterfaceName *string `min:"1" max:"128" type:"string"`
}

// String returns the string representation
func (s ModifyDirectConnectVirtualInterfaceAttributesInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDirectConnectVirtualInterfaceAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDirectConnectVirtualInterfaceAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDirectConnectVirtualInterfaceAttributesInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.VirtualInterfaceId == nil {
		invalidParams.Add(request.NewErrParamRequired("VirtualInterfaceId"))
	}
	if s.VirtualInterfaceName != nil && len(*s.VirtualInterfaceName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("VirtualInterfaceName", 1))
	}
	if s.VirtualInterfaceName != nil && len(*s.VirtualInterfaceName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("VirtualInterfaceName", 128, *s.VirtualInterfaceName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyDirectConnectVirtualInterfaceAttributesInput) SetDescription(v string) *ModifyDirectConnectVirtualInterfaceAttributesInput {
	s.Description = &v
	return s
}

// SetVirtualInterfaceId sets the VirtualInterfaceId field's value.
func (s *ModifyDirectConnectVirtualInterfaceAttributesInput) SetVirtualInterfaceId(v string) *ModifyDirectConnectVirtualInterfaceAttributesInput {
	s.VirtualInterfaceId = &v
	return s
}

// SetVirtualInterfaceName sets the VirtualInterfaceName field's value.
func (s *ModifyDirectConnectVirtualInterfaceAttributesInput) SetVirtualInterfaceName(v string) *ModifyDirectConnectVirtualInterfaceAttributesInput {
	s.VirtualInterfaceName = &v
	return s
}

type ModifyDirectConnectVirtualInterfaceAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyDirectConnectVirtualInterfaceAttributesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDirectConnectVirtualInterfaceAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyDirectConnectVirtualInterfaceAttributesOutput) SetRequestId(v string) *ModifyDirectConnectVirtualInterfaceAttributesOutput {
	s.RequestId = &v
	return s
}
