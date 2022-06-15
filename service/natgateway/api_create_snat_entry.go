// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package natgateway

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opCreateSnatEntryCommon = "CreateSnatEntry"

// CreateSnatEntryCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the CreateSnatEntryCommon operation. The "output" return
// value will be populated with the CreateSnatEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateSnatEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateSnatEntryCommon Send returns without error.
//
// See CreateSnatEntryCommon for more information on using the CreateSnatEntryCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateSnatEntryCommonRequest method.
//    req, resp := client.CreateSnatEntryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NATGATEWAY) CreateSnatEntryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateSnatEntryCommon,
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

// CreateSnatEntryCommon API operation for NATGATEWAY.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for NATGATEWAY's
// API operation CreateSnatEntryCommon for usage and error information.
func (c *NATGATEWAY) CreateSnatEntryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateSnatEntryCommonRequest(input)
	return out, req.Send()
}

// CreateSnatEntryCommonWithContext is the same as CreateSnatEntryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateSnatEntryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) CreateSnatEntryCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateSnatEntryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateSnatEntry = "CreateSnatEntry"

// CreateSnatEntryRequest generates a "volcstack/request.Request" representing the
// client's request for the CreateSnatEntry operation. The "output" return
// value will be populated with the CreateSnatEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateSnatEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateSnatEntryCommon Send returns without error.
//
// See CreateSnatEntry for more information on using the CreateSnatEntry
// API call, and error handling.
//
//    // Example sending a request using the CreateSnatEntryRequest method.
//    req, resp := client.CreateSnatEntryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NATGATEWAY) CreateSnatEntryRequest(input *CreateSnatEntryInput) (req *request.Request, output *CreateSnatEntryOutput) {
	op := &request.Operation{
		Name:       opCreateSnatEntry,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSnatEntryInput{}
	}

	output = &CreateSnatEntryOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateSnatEntry API operation for NATGATEWAY.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for NATGATEWAY's
// API operation CreateSnatEntry for usage and error information.
func (c *NATGATEWAY) CreateSnatEntry(input *CreateSnatEntryInput) (*CreateSnatEntryOutput, error) {
	req, out := c.CreateSnatEntryRequest(input)
	return out, req.Send()
}

// CreateSnatEntryWithContext is the same as CreateSnatEntry with the addition of
// the ability to pass a context and additional request options.
//
// See CreateSnatEntry for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) CreateSnatEntryWithContext(ctx volcstack.Context, input *CreateSnatEntryInput, opts ...request.Option) (*CreateSnatEntryOutput, error) {
	req, out := c.CreateSnatEntryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateSnatEntryInput struct {
	_ struct{} `type:"structure"`

	// EipId is a required field
	EipId *string `type:"string" required:"true"`

	// NatGatewayId is a required field
	NatGatewayId *string `type:"string" required:"true"`

	SnatEntryName *string `min:"1" max:"128" type:"string"`

	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSnatEntryInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateSnatEntryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSnatEntryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateSnatEntryInput"}
	if s.EipId == nil {
		invalidParams.Add(request.NewErrParamRequired("EipId"))
	}
	if s.NatGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("NatGatewayId"))
	}
	if s.SnatEntryName != nil && len(*s.SnatEntryName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SnatEntryName", 1))
	}
	if s.SnatEntryName != nil && len(*s.SnatEntryName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("SnatEntryName", 128, *s.SnatEntryName))
	}
	if s.SubnetId == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEipId sets the EipId field's value.
func (s *CreateSnatEntryInput) SetEipId(v string) *CreateSnatEntryInput {
	s.EipId = &v
	return s
}

// SetNatGatewayId sets the NatGatewayId field's value.
func (s *CreateSnatEntryInput) SetNatGatewayId(v string) *CreateSnatEntryInput {
	s.NatGatewayId = &v
	return s
}

// SetSnatEntryName sets the SnatEntryName field's value.
func (s *CreateSnatEntryInput) SetSnatEntryName(v string) *CreateSnatEntryInput {
	s.SnatEntryName = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *CreateSnatEntryInput) SetSubnetId(v string) *CreateSnatEntryInput {
	s.SubnetId = &v
	return s
}

type CreateSnatEntryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	SnatEntryId *string `type:"string"`
}

// String returns the string representation
func (s CreateSnatEntryOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateSnatEntryOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *CreateSnatEntryOutput) SetRequestId(v string) *CreateSnatEntryOutput {
	s.RequestId = &v
	return s
}

// SetSnatEntryId sets the SnatEntryId field's value.
func (s *CreateSnatEntryOutput) SetSnatEntryId(v string) *CreateSnatEntryOutput {
	s.SnatEntryId = &v
	return s
}