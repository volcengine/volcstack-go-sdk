// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opCreateDirectConnectVirtualInterfaceCommon = "CreateDirectConnectVirtualInterface"

// CreateDirectConnectVirtualInterfaceCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the CreateDirectConnectVirtualInterfaceCommon operation. The "output" return
// value will be populated with the CreateDirectConnectVirtualInterfaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDirectConnectVirtualInterfaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDirectConnectVirtualInterfaceCommon Send returns without error.
//
// See CreateDirectConnectVirtualInterfaceCommon for more information on using the CreateDirectConnectVirtualInterfaceCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateDirectConnectVirtualInterfaceCommonRequest method.
//    req, resp := client.CreateDirectConnectVirtualInterfaceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) CreateDirectConnectVirtualInterfaceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDirectConnectVirtualInterfaceCommon,
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

// CreateDirectConnectVirtualInterfaceCommon API operation for DIRECTCONNECT.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation CreateDirectConnectVirtualInterfaceCommon for usage and error information.
func (c *DIRECTCONNECT) CreateDirectConnectVirtualInterfaceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDirectConnectVirtualInterfaceCommonRequest(input)
	return out, req.Send()
}

// CreateDirectConnectVirtualInterfaceCommonWithContext is the same as CreateDirectConnectVirtualInterfaceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDirectConnectVirtualInterfaceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) CreateDirectConnectVirtualInterfaceCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDirectConnectVirtualInterfaceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDirectConnectVirtualInterface = "CreateDirectConnectVirtualInterface"

// CreateDirectConnectVirtualInterfaceRequest generates a "volcstack/request.Request" representing the
// client's request for the CreateDirectConnectVirtualInterface operation. The "output" return
// value will be populated with the CreateDirectConnectVirtualInterfaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDirectConnectVirtualInterfaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDirectConnectVirtualInterfaceCommon Send returns without error.
//
// See CreateDirectConnectVirtualInterface for more information on using the CreateDirectConnectVirtualInterface
// API call, and error handling.
//
//    // Example sending a request using the CreateDirectConnectVirtualInterfaceRequest method.
//    req, resp := client.CreateDirectConnectVirtualInterfaceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) CreateDirectConnectVirtualInterfaceRequest(input *CreateDirectConnectVirtualInterfaceInput) (req *request.Request, output *CreateDirectConnectVirtualInterfaceOutput) {
	op := &request.Operation{
		Name:       opCreateDirectConnectVirtualInterface,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDirectConnectVirtualInterfaceInput{}
	}

	output = &CreateDirectConnectVirtualInterfaceOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateDirectConnectVirtualInterface API operation for DIRECTCONNECT.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation CreateDirectConnectVirtualInterface for usage and error information.
func (c *DIRECTCONNECT) CreateDirectConnectVirtualInterface(input *CreateDirectConnectVirtualInterfaceInput) (*CreateDirectConnectVirtualInterfaceOutput, error) {
	req, out := c.CreateDirectConnectVirtualInterfaceRequest(input)
	return out, req.Send()
}

// CreateDirectConnectVirtualInterfaceWithContext is the same as CreateDirectConnectVirtualInterface with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDirectConnectVirtualInterface for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) CreateDirectConnectVirtualInterfaceWithContext(ctx volcstack.Context, input *CreateDirectConnectVirtualInterfaceInput, opts ...request.Option) (*CreateDirectConnectVirtualInterfaceOutput, error) {
	req, out := c.CreateDirectConnectVirtualInterfaceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDirectConnectVirtualInterfaceInput struct {
	_ struct{} `type:"structure"`

	Description *string `min:"1" max:"255" type:"string"`

	// DirectConnectConnectionId is a required field
	DirectConnectConnectionId *string `type:"string" required:"true"`

	// DirectConnectGatewayId is a required field
	DirectConnectGatewayId *string `type:"string" required:"true"`

	// LocalIp is a required field
	LocalIp *string `type:"string" required:"true"`

	// PeerIp is a required field
	PeerIp *string `type:"string" required:"true"`

	RouteType *string `type:"string" enum:"RouteTypeForCreateDirectConnectVirtualInterfaceInput"`

	VirtualInterfaceName *string `min:"1" max:"128" type:"string"`

	// VlanId is a required field
	VlanId *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s CreateDirectConnectVirtualInterfaceInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDirectConnectVirtualInterfaceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDirectConnectVirtualInterfaceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateDirectConnectVirtualInterfaceInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.DirectConnectConnectionId == nil {
		invalidParams.Add(request.NewErrParamRequired("DirectConnectConnectionId"))
	}
	if s.DirectConnectGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("DirectConnectGatewayId"))
	}
	if s.LocalIp == nil {
		invalidParams.Add(request.NewErrParamRequired("LocalIp"))
	}
	if s.PeerIp == nil {
		invalidParams.Add(request.NewErrParamRequired("PeerIp"))
	}
	if s.VirtualInterfaceName != nil && len(*s.VirtualInterfaceName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("VirtualInterfaceName", 1))
	}
	if s.VirtualInterfaceName != nil && len(*s.VirtualInterfaceName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("VirtualInterfaceName", 128, *s.VirtualInterfaceName))
	}
	if s.VlanId == nil {
		invalidParams.Add(request.NewErrParamRequired("VlanId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *CreateDirectConnectVirtualInterfaceInput) SetDescription(v string) *CreateDirectConnectVirtualInterfaceInput {
	s.Description = &v
	return s
}

// SetDirectConnectConnectionId sets the DirectConnectConnectionId field's value.
func (s *CreateDirectConnectVirtualInterfaceInput) SetDirectConnectConnectionId(v string) *CreateDirectConnectVirtualInterfaceInput {
	s.DirectConnectConnectionId = &v
	return s
}

// SetDirectConnectGatewayId sets the DirectConnectGatewayId field's value.
func (s *CreateDirectConnectVirtualInterfaceInput) SetDirectConnectGatewayId(v string) *CreateDirectConnectVirtualInterfaceInput {
	s.DirectConnectGatewayId = &v
	return s
}

// SetLocalIp sets the LocalIp field's value.
func (s *CreateDirectConnectVirtualInterfaceInput) SetLocalIp(v string) *CreateDirectConnectVirtualInterfaceInput {
	s.LocalIp = &v
	return s
}

// SetPeerIp sets the PeerIp field's value.
func (s *CreateDirectConnectVirtualInterfaceInput) SetPeerIp(v string) *CreateDirectConnectVirtualInterfaceInput {
	s.PeerIp = &v
	return s
}

// SetRouteType sets the RouteType field's value.
func (s *CreateDirectConnectVirtualInterfaceInput) SetRouteType(v string) *CreateDirectConnectVirtualInterfaceInput {
	s.RouteType = &v
	return s
}

// SetVirtualInterfaceName sets the VirtualInterfaceName field's value.
func (s *CreateDirectConnectVirtualInterfaceInput) SetVirtualInterfaceName(v string) *CreateDirectConnectVirtualInterfaceInput {
	s.VirtualInterfaceName = &v
	return s
}

// SetVlanId sets the VlanId field's value.
func (s *CreateDirectConnectVirtualInterfaceInput) SetVlanId(v int64) *CreateDirectConnectVirtualInterfaceInput {
	s.VlanId = &v
	return s
}

type CreateDirectConnectVirtualInterfaceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	VirtualInterfaceId *string `type:"string"`
}

// String returns the string representation
func (s CreateDirectConnectVirtualInterfaceOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDirectConnectVirtualInterfaceOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *CreateDirectConnectVirtualInterfaceOutput) SetRequestId(v string) *CreateDirectConnectVirtualInterfaceOutput {
	s.RequestId = &v
	return s
}

// SetVirtualInterfaceId sets the VirtualInterfaceId field's value.
func (s *CreateDirectConnectVirtualInterfaceOutput) SetVirtualInterfaceId(v string) *CreateDirectConnectVirtualInterfaceOutput {
	s.VirtualInterfaceId = &v
	return s
}

const (
	// RouteTypeForCreateDirectConnectVirtualInterfaceInputStatic is a RouteTypeForCreateDirectConnectVirtualInterfaceInput enum value
	RouteTypeForCreateDirectConnectVirtualInterfaceInputStatic = "STATIC"

	// RouteTypeForCreateDirectConnectVirtualInterfaceInputBgp is a RouteTypeForCreateDirectConnectVirtualInterfaceInput enum value
	RouteTypeForCreateDirectConnectVirtualInterfaceInputBgp = "BGP"
)
