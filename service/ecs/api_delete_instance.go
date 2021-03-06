// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDeleteInstanceCommon = "DeleteInstance"

// DeleteInstanceCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DeleteInstanceCommon operation. The "output" return
// value will be populated with the DeleteInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteInstanceCommon Send returns without error.
//
// See DeleteInstanceCommon for more information on using the DeleteInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteInstanceCommonRequest method.
//    req, resp := client.DeleteInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DeleteInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteInstanceCommon,
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

// DeleteInstanceCommon API operation for ECS.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for ECS's
// API operation DeleteInstanceCommon for usage and error information.
func (c *ECS) DeleteInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteInstanceCommonRequest(input)
	return out, req.Send()
}

// DeleteInstanceCommonWithContext is the same as DeleteInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DeleteInstanceCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteInstance = "DeleteInstance"

// DeleteInstanceRequest generates a "volcstack/request.Request" representing the
// client's request for the DeleteInstance operation. The "output" return
// value will be populated with the DeleteInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteInstanceCommon Send returns without error.
//
// See DeleteInstance for more information on using the DeleteInstance
// API call, and error handling.
//
//    // Example sending a request using the DeleteInstanceRequest method.
//    req, resp := client.DeleteInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DeleteInstanceRequest(input *DeleteInstanceInput) (req *request.Request, output *DeleteInstanceOutput) {
	op := &request.Operation{
		Name:       opDeleteInstance,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteInstanceInput{}
	}

	output = &DeleteInstanceOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteInstance API operation for ECS.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for ECS's
// API operation DeleteInstance for usage and error information.
func (c *ECS) DeleteInstance(input *DeleteInstanceInput) (*DeleteInstanceOutput, error) {
	req, out := c.DeleteInstanceRequest(input)
	return out, req.Send()
}

// DeleteInstanceWithContext is the same as DeleteInstance with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DeleteInstanceWithContext(ctx volcstack.Context, input *DeleteInstanceInput, opts ...request.Option) (*DeleteInstanceOutput, error) {
	req, out := c.DeleteInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteInstanceInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s DeleteInstanceInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteInstanceInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *DeleteInstanceInput) SetInstanceId(v string) *DeleteInstanceInput {
	s.InstanceId = &v
	return s
}

type DeleteInstanceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteInstanceOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteInstanceOutput) GoString() string {
	return s.String()
}
