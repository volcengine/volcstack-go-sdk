// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opModifyCenBandwidthPackageAttributesCommon = "ModifyCenBandwidthPackageAttributes"

// ModifyCenBandwidthPackageAttributesCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the ModifyCenBandwidthPackageAttributesCommon operation. The "output" return
// value will be populated with the ModifyCenBandwidthPackageAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyCenBandwidthPackageAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyCenBandwidthPackageAttributesCommon Send returns without error.
//
// See ModifyCenBandwidthPackageAttributesCommon for more information on using the ModifyCenBandwidthPackageAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyCenBandwidthPackageAttributesCommonRequest method.
//    req, resp := client.ModifyCenBandwidthPackageAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) ModifyCenBandwidthPackageAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyCenBandwidthPackageAttributesCommon,
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

// ModifyCenBandwidthPackageAttributesCommon API operation for CEN.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CEN's
// API operation ModifyCenBandwidthPackageAttributesCommon for usage and error information.
func (c *CEN) ModifyCenBandwidthPackageAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyCenBandwidthPackageAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyCenBandwidthPackageAttributesCommonWithContext is the same as ModifyCenBandwidthPackageAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyCenBandwidthPackageAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) ModifyCenBandwidthPackageAttributesCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyCenBandwidthPackageAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyCenBandwidthPackageAttributes = "ModifyCenBandwidthPackageAttributes"

// ModifyCenBandwidthPackageAttributesRequest generates a "volcstack/request.Request" representing the
// client's request for the ModifyCenBandwidthPackageAttributes operation. The "output" return
// value will be populated with the ModifyCenBandwidthPackageAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyCenBandwidthPackageAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyCenBandwidthPackageAttributesCommon Send returns without error.
//
// See ModifyCenBandwidthPackageAttributes for more information on using the ModifyCenBandwidthPackageAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyCenBandwidthPackageAttributesRequest method.
//    req, resp := client.ModifyCenBandwidthPackageAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) ModifyCenBandwidthPackageAttributesRequest(input *ModifyCenBandwidthPackageAttributesInput) (req *request.Request, output *ModifyCenBandwidthPackageAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyCenBandwidthPackageAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyCenBandwidthPackageAttributesInput{}
	}

	output = &ModifyCenBandwidthPackageAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyCenBandwidthPackageAttributes API operation for CEN.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CEN's
// API operation ModifyCenBandwidthPackageAttributes for usage and error information.
func (c *CEN) ModifyCenBandwidthPackageAttributes(input *ModifyCenBandwidthPackageAttributesInput) (*ModifyCenBandwidthPackageAttributesOutput, error) {
	req, out := c.ModifyCenBandwidthPackageAttributesRequest(input)
	return out, req.Send()
}

// ModifyCenBandwidthPackageAttributesWithContext is the same as ModifyCenBandwidthPackageAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyCenBandwidthPackageAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) ModifyCenBandwidthPackageAttributesWithContext(ctx volcstack.Context, input *ModifyCenBandwidthPackageAttributesInput, opts ...request.Option) (*ModifyCenBandwidthPackageAttributesOutput, error) {
	req, out := c.ModifyCenBandwidthPackageAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyCenBandwidthPackageAttributesInput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int64 `type:"integer"`

	// CenBandwidthPackageId is a required field
	CenBandwidthPackageId *string `type:"string" required:"true"`

	CenBandwidthPackageName *string `min:"1" max:"128" type:"string"`

	Description *string `min:"1" max:"255" type:"string"`
}

// String returns the string representation
func (s ModifyCenBandwidthPackageAttributesInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyCenBandwidthPackageAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyCenBandwidthPackageAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyCenBandwidthPackageAttributesInput"}
	if s.CenBandwidthPackageId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenBandwidthPackageId"))
	}
	if s.CenBandwidthPackageName != nil && len(*s.CenBandwidthPackageName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CenBandwidthPackageName", 1))
	}
	if s.CenBandwidthPackageName != nil && len(*s.CenBandwidthPackageName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("CenBandwidthPackageName", 128, *s.CenBandwidthPackageName))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBandwidth sets the Bandwidth field's value.
func (s *ModifyCenBandwidthPackageAttributesInput) SetBandwidth(v int64) *ModifyCenBandwidthPackageAttributesInput {
	s.Bandwidth = &v
	return s
}

// SetCenBandwidthPackageId sets the CenBandwidthPackageId field's value.
func (s *ModifyCenBandwidthPackageAttributesInput) SetCenBandwidthPackageId(v string) *ModifyCenBandwidthPackageAttributesInput {
	s.CenBandwidthPackageId = &v
	return s
}

// SetCenBandwidthPackageName sets the CenBandwidthPackageName field's value.
func (s *ModifyCenBandwidthPackageAttributesInput) SetCenBandwidthPackageName(v string) *ModifyCenBandwidthPackageAttributesInput {
	s.CenBandwidthPackageName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyCenBandwidthPackageAttributesInput) SetDescription(v string) *ModifyCenBandwidthPackageAttributesInput {
	s.Description = &v
	return s
}

type ModifyCenBandwidthPackageAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PreOrderNumber *string `type:"string"`
}

// String returns the string representation
func (s ModifyCenBandwidthPackageAttributesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyCenBandwidthPackageAttributesOutput) GoString() string {
	return s.String()
}

// SetPreOrderNumber sets the PreOrderNumber field's value.
func (s *ModifyCenBandwidthPackageAttributesOutput) SetPreOrderNumber(v string) *ModifyCenBandwidthPackageAttributesOutput {
	s.PreOrderNumber = &v
	return s
}
