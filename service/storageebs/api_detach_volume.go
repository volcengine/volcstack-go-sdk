// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDetachVolumeCommon = "DetachVolume"

// DetachVolumeCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DetachVolumeCommon operation. The "output" return
// value will be populated with the DetachVolumeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachVolumeCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachVolumeCommon Send returns without error.
//
// See DetachVolumeCommon for more information on using the DetachVolumeCommon
// API call, and error handling.
//
//    // Example sending a request using the DetachVolumeCommonRequest method.
//    req, resp := client.DetachVolumeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) DetachVolumeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDetachVolumeCommon,
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

// DetachVolumeCommon API operation for STORAGE_EBS.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for STORAGE_EBS's
// API operation DetachVolumeCommon for usage and error information.
func (c *STORAGEEBS) DetachVolumeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DetachVolumeCommonRequest(input)
	return out, req.Send()
}

// DetachVolumeCommonWithContext is the same as DetachVolumeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DetachVolumeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) DetachVolumeCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DetachVolumeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDetachVolume = "DetachVolume"

// DetachVolumeRequest generates a "volcstack/request.Request" representing the
// client's request for the DetachVolume operation. The "output" return
// value will be populated with the DetachVolumeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachVolumeCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachVolumeCommon Send returns without error.
//
// See DetachVolume for more information on using the DetachVolume
// API call, and error handling.
//
//    // Example sending a request using the DetachVolumeRequest method.
//    req, resp := client.DetachVolumeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) DetachVolumeRequest(input *DetachVolumeInput) (req *request.Request, output *DetachVolumeOutput) {
	op := &request.Operation{
		Name:       opDetachVolume,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachVolumeInput{}
	}

	output = &DetachVolumeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DetachVolume API operation for STORAGE_EBS.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for STORAGE_EBS's
// API operation DetachVolume for usage and error information.
func (c *STORAGEEBS) DetachVolume(input *DetachVolumeInput) (*DetachVolumeOutput, error) {
	req, out := c.DetachVolumeRequest(input)
	return out, req.Send()
}

// DetachVolumeWithContext is the same as DetachVolume with the addition of
// the ability to pass a context and additional request options.
//
// See DetachVolume for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) DetachVolumeWithContext(ctx volcstack.Context, input *DetachVolumeInput, opts ...request.Option) (*DetachVolumeOutput, error) {
	req, out := c.DetachVolumeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DetachVolumeInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	VolumeId *string `type:"string"`
}

// String returns the string representation
func (s DetachVolumeInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachVolumeInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *DetachVolumeInput) SetInstanceId(v string) *DetachVolumeInput {
	s.InstanceId = &v
	return s
}

// SetVolumeId sets the VolumeId field's value.
func (s *DetachVolumeInput) SetVolumeId(v string) *DetachVolumeInput {
	s.VolumeId = &v
	return s
}

type DetachVolumeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DetachVolumeOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachVolumeOutput) GoString() string {
	return s.String()
}
