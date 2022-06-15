// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opAllocateEipAddressCommon = "AllocateEipAddress"

// AllocateEipAddressCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the AllocateEipAddressCommon operation. The "output" return
// value will be populated with the AllocateEipAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AllocateEipAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after AllocateEipAddressCommon Send returns without error.
//
// See AllocateEipAddressCommon for more information on using the AllocateEipAddressCommon
// API call, and error handling.
//
//    // Example sending a request using the AllocateEipAddressCommonRequest method.
//    req, resp := client.AllocateEipAddressCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) AllocateEipAddressCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAllocateEipAddressCommon,
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

// AllocateEipAddressCommon API operation for VPC.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPC's
// API operation AllocateEipAddressCommon for usage and error information.
func (c *VPC) AllocateEipAddressCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AllocateEipAddressCommonRequest(input)
	return out, req.Send()
}

// AllocateEipAddressCommonWithContext is the same as AllocateEipAddressCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AllocateEipAddressCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AllocateEipAddressCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AllocateEipAddressCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAllocateEipAddress = "AllocateEipAddress"

// AllocateEipAddressRequest generates a "volcstack/request.Request" representing the
// client's request for the AllocateEipAddress operation. The "output" return
// value will be populated with the AllocateEipAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AllocateEipAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after AllocateEipAddressCommon Send returns without error.
//
// See AllocateEipAddress for more information on using the AllocateEipAddress
// API call, and error handling.
//
//    // Example sending a request using the AllocateEipAddressRequest method.
//    req, resp := client.AllocateEipAddressRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) AllocateEipAddressRequest(input *AllocateEipAddressInput) (req *request.Request, output *AllocateEipAddressOutput) {
	op := &request.Operation{
		Name:       opAllocateEipAddress,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AllocateEipAddressInput{}
	}

	output = &AllocateEipAddressOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AllocateEipAddress API operation for VPC.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPC's
// API operation AllocateEipAddress for usage and error information.
func (c *VPC) AllocateEipAddress(input *AllocateEipAddressInput) (*AllocateEipAddressOutput, error) {
	req, out := c.AllocateEipAddressRequest(input)
	return out, req.Send()
}

// AllocateEipAddressWithContext is the same as AllocateEipAddress with the addition of
// the ability to pass a context and additional request options.
//
// See AllocateEipAddress for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AllocateEipAddressWithContext(ctx volcstack.Context, input *AllocateEipAddressInput, opts ...request.Option) (*AllocateEipAddressOutput, error) {
	req, out := c.AllocateEipAddressRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AllocateEipAddressInput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int64 `min:"1" max:"500" type:"integer"`

	BillingType *int64 `min:"1" max:"3" type:"integer"`

	Description *string `min:"1" max:"255" type:"string"`

	ISP *string `type:"string" enum:"ISPForAllocateEipAddressInput"`

	Name *string `min:"1" max:"128" type:"string"`

	Period *int64 `type:"integer"`

	PeriodUnit *int64 `min:"1" max:"2" type:"integer"`

	SecurityProtectionTypes []*string `type:"list"`
}

// String returns the string representation
func (s AllocateEipAddressInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s AllocateEipAddressInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AllocateEipAddressInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AllocateEipAddressInput"}
	if s.Bandwidth != nil && *s.Bandwidth < 1 {
		invalidParams.Add(request.NewErrParamMinValue("Bandwidth", 1))
	}
	if s.Bandwidth != nil && *s.Bandwidth > 500 {
		invalidParams.Add(request.NewErrParamMaxValue("Bandwidth", 500))
	}
	if s.BillingType != nil && *s.BillingType < 1 {
		invalidParams.Add(request.NewErrParamMinValue("BillingType", 1))
	}
	if s.BillingType != nil && *s.BillingType > 3 {
		invalidParams.Add(request.NewErrParamMaxValue("BillingType", 3))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Name", 1))
	}
	if s.Name != nil && len(*s.Name) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("Name", 128, *s.Name))
	}
	if s.PeriodUnit != nil && *s.PeriodUnit < 1 {
		invalidParams.Add(request.NewErrParamMinValue("PeriodUnit", 1))
	}
	if s.PeriodUnit != nil && *s.PeriodUnit > 2 {
		invalidParams.Add(request.NewErrParamMaxValue("PeriodUnit", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBandwidth sets the Bandwidth field's value.
func (s *AllocateEipAddressInput) SetBandwidth(v int64) *AllocateEipAddressInput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *AllocateEipAddressInput) SetBillingType(v int64) *AllocateEipAddressInput {
	s.BillingType = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *AllocateEipAddressInput) SetDescription(v string) *AllocateEipAddressInput {
	s.Description = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *AllocateEipAddressInput) SetISP(v string) *AllocateEipAddressInput {
	s.ISP = &v
	return s
}

// SetName sets the Name field's value.
func (s *AllocateEipAddressInput) SetName(v string) *AllocateEipAddressInput {
	s.Name = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *AllocateEipAddressInput) SetPeriod(v int64) *AllocateEipAddressInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *AllocateEipAddressInput) SetPeriodUnit(v int64) *AllocateEipAddressInput {
	s.PeriodUnit = &v
	return s
}

// SetSecurityProtectionTypes sets the SecurityProtectionTypes field's value.
func (s *AllocateEipAddressInput) SetSecurityProtectionTypes(v []*string) *AllocateEipAddressInput {
	s.SecurityProtectionTypes = v
	return s
}

type AllocateEipAddressOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AllocationId *string `type:"string"`

	EipAddress *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s AllocateEipAddressOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s AllocateEipAddressOutput) GoString() string {
	return s.String()
}

// SetAllocationId sets the AllocationId field's value.
func (s *AllocateEipAddressOutput) SetAllocationId(v string) *AllocateEipAddressOutput {
	s.AllocationId = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *AllocateEipAddressOutput) SetEipAddress(v string) *AllocateEipAddressOutput {
	s.EipAddress = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *AllocateEipAddressOutput) SetRequestId(v string) *AllocateEipAddressOutput {
	s.RequestId = &v
	return s
}

const (
	// ISPForAllocateEipAddressInputBgp is a ISPForAllocateEipAddressInput enum value
	ISPForAllocateEipAddressInputBgp = "BGP"

	// ISPForAllocateEipAddressInputChinaMobile is a ISPForAllocateEipAddressInput enum value
	ISPForAllocateEipAddressInputChinaMobile = "ChinaMobile"

	// ISPForAllocateEipAddressInputChinaUnicom is a ISPForAllocateEipAddressInput enum value
	ISPForAllocateEipAddressInputChinaUnicom = "ChinaUnicom"

	// ISPForAllocateEipAddressInputChinaTelecom is a ISPForAllocateEipAddressInput enum value
	ISPForAllocateEipAddressInputChinaTelecom = "ChinaTelecom"
)
