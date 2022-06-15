// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opAttachInstanceToCenCommon = "AttachInstanceToCen"

// AttachInstanceToCenCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the AttachInstanceToCenCommon operation. The "output" return
// value will be populated with the AttachInstanceToCenCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachInstanceToCenCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachInstanceToCenCommon Send returns without error.
//
// See AttachInstanceToCenCommon for more information on using the AttachInstanceToCenCommon
// API call, and error handling.
//
//    // Example sending a request using the AttachInstanceToCenCommonRequest method.
//    req, resp := client.AttachInstanceToCenCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) AttachInstanceToCenCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAttachInstanceToCenCommon,
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

// AttachInstanceToCenCommon API operation for CEN.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CEN's
// API operation AttachInstanceToCenCommon for usage and error information.
func (c *CEN) AttachInstanceToCenCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AttachInstanceToCenCommonRequest(input)
	return out, req.Send()
}

// AttachInstanceToCenCommonWithContext is the same as AttachInstanceToCenCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AttachInstanceToCenCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) AttachInstanceToCenCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AttachInstanceToCenCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAttachInstanceToCen = "AttachInstanceToCen"

// AttachInstanceToCenRequest generates a "volcstack/request.Request" representing the
// client's request for the AttachInstanceToCen operation. The "output" return
// value will be populated with the AttachInstanceToCenCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachInstanceToCenCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachInstanceToCenCommon Send returns without error.
//
// See AttachInstanceToCen for more information on using the AttachInstanceToCen
// API call, and error handling.
//
//    // Example sending a request using the AttachInstanceToCenRequest method.
//    req, resp := client.AttachInstanceToCenRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) AttachInstanceToCenRequest(input *AttachInstanceToCenInput) (req *request.Request, output *AttachInstanceToCenOutput) {
	op := &request.Operation{
		Name:       opAttachInstanceToCen,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachInstanceToCenInput{}
	}

	output = &AttachInstanceToCenOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AttachInstanceToCen API operation for CEN.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CEN's
// API operation AttachInstanceToCen for usage and error information.
func (c *CEN) AttachInstanceToCen(input *AttachInstanceToCenInput) (*AttachInstanceToCenOutput, error) {
	req, out := c.AttachInstanceToCenRequest(input)
	return out, req.Send()
}

// AttachInstanceToCenWithContext is the same as AttachInstanceToCen with the addition of
// the ability to pass a context and additional request options.
//
// See AttachInstanceToCen for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) AttachInstanceToCenWithContext(ctx volcstack.Context, input *AttachInstanceToCenInput, opts ...request.Option) (*AttachInstanceToCenOutput, error) {
	req, out := c.AttachInstanceToCenRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachInstanceToCenInput struct {
	_ struct{} `type:"structure"`

	// CenId is a required field
	CenId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// InstanceRegionId is a required field
	InstanceRegionId *string `type:"string" required:"true"`

	// InstanceType is a required field
	InstanceType *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AttachInstanceToCenInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachInstanceToCenInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachInstanceToCenInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AttachInstanceToCenInput"}
	if s.CenId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceRegionId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceRegionId"))
	}
	if s.InstanceType == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCenId sets the CenId field's value.
func (s *AttachInstanceToCenInput) SetCenId(v string) *AttachInstanceToCenInput {
	s.CenId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *AttachInstanceToCenInput) SetInstanceId(v string) *AttachInstanceToCenInput {
	s.InstanceId = &v
	return s
}

// SetInstanceRegionId sets the InstanceRegionId field's value.
func (s *AttachInstanceToCenInput) SetInstanceRegionId(v string) *AttachInstanceToCenInput {
	s.InstanceRegionId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *AttachInstanceToCenInput) SetInstanceType(v string) *AttachInstanceToCenInput {
	s.InstanceType = &v
	return s
}

type AttachInstanceToCenOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AttachInstanceToCenOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachInstanceToCenOutput) GoString() string {
	return s.String()
}
