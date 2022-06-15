// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opStopInstanceCommon = "StopInstance"

// StopInstanceCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the StopInstanceCommon operation. The "output" return
// value will be populated with the StopInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StopInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after StopInstanceCommon Send returns without error.
//
// See StopInstanceCommon for more information on using the StopInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the StopInstanceCommonRequest method.
//    req, resp := client.StopInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) StopInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStopInstanceCommon,
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

// StopInstanceCommon API operation for ECS.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for ECS's
// API operation StopInstanceCommon for usage and error information.
func (c *ECS) StopInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StopInstanceCommonRequest(input)
	return out, req.Send()
}

// StopInstanceCommonWithContext is the same as StopInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StopInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) StopInstanceCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StopInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStopInstance = "StopInstance"

// StopInstanceRequest generates a "volcstack/request.Request" representing the
// client's request for the StopInstance operation. The "output" return
// value will be populated with the StopInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StopInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after StopInstanceCommon Send returns without error.
//
// See StopInstance for more information on using the StopInstance
// API call, and error handling.
//
//    // Example sending a request using the StopInstanceRequest method.
//    req, resp := client.StopInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) StopInstanceRequest(input *StopInstanceInput) (req *request.Request, output *StopInstanceOutput) {
	op := &request.Operation{
		Name:       opStopInstance,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopInstanceInput{}
	}

	output = &StopInstanceOutput{}
	req = c.newRequest(op, input, output)

	return
}

// StopInstance API operation for ECS.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for ECS's
// API operation StopInstance for usage and error information.
func (c *ECS) StopInstance(input *StopInstanceInput) (*StopInstanceOutput, error) {
	req, out := c.StopInstanceRequest(input)
	return out, req.Send()
}

// StopInstanceWithContext is the same as StopInstance with the addition of
// the ability to pass a context and additional request options.
//
// See StopInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) StopInstanceWithContext(ctx volcstack.Context, input *StopInstanceInput, opts ...request.Option) (*StopInstanceOutput, error) {
	req, out := c.StopInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StopInstanceInput struct {
	_ struct{} `type:"structure"`

	ForceStop *bool `type:"boolean"`

	InstanceId *string `type:"string"`

	StoppedMode *string `type:"string"`
}

// String returns the string representation
func (s StopInstanceInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s StopInstanceInput) GoString() string {
	return s.String()
}

// SetForceStop sets the ForceStop field's value.
func (s *StopInstanceInput) SetForceStop(v bool) *StopInstanceInput {
	s.ForceStop = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *StopInstanceInput) SetInstanceId(v string) *StopInstanceInput {
	s.InstanceId = &v
	return s
}

// SetStoppedMode sets the StoppedMode field's value.
func (s *StopInstanceInput) SetStoppedMode(v string) *StopInstanceInput {
	s.StoppedMode = &v
	return s
}

type StopInstanceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s StopInstanceOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s StopInstanceOutput) GoString() string {
	return s.String()
}
