// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDeleteSubnetCommon = "DeleteSubnet"

// DeleteSubnetCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DeleteSubnetCommon operation. The "output" return
// value will be populated with the DeleteSubnetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteSubnetCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteSubnetCommon Send returns without error.
//
// See DeleteSubnetCommon for more information on using the DeleteSubnetCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteSubnetCommonRequest method.
//    req, resp := client.DeleteSubnetCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DeleteSubnetCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteSubnetCommon,
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

// DeleteSubnetCommon API operation for VPC.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPC's
// API operation DeleteSubnetCommon for usage and error information.
func (c *VPC) DeleteSubnetCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteSubnetCommonRequest(input)
	return out, req.Send()
}

// DeleteSubnetCommonWithContext is the same as DeleteSubnetCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteSubnetCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteSubnetCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteSubnetCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteSubnet = "DeleteSubnet"

// DeleteSubnetRequest generates a "volcstack/request.Request" representing the
// client's request for the DeleteSubnet operation. The "output" return
// value will be populated with the DeleteSubnetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteSubnetCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteSubnetCommon Send returns without error.
//
// See DeleteSubnet for more information on using the DeleteSubnet
// API call, and error handling.
//
//    // Example sending a request using the DeleteSubnetRequest method.
//    req, resp := client.DeleteSubnetRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DeleteSubnetRequest(input *DeleteSubnetInput) (req *request.Request, output *DeleteSubnetOutput) {
	op := &request.Operation{
		Name:       opDeleteSubnet,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSubnetInput{}
	}

	output = &DeleteSubnetOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteSubnet API operation for VPC.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPC's
// API operation DeleteSubnet for usage and error information.
func (c *VPC) DeleteSubnet(input *DeleteSubnetInput) (*DeleteSubnetOutput, error) {
	req, out := c.DeleteSubnetRequest(input)
	return out, req.Send()
}

// DeleteSubnetWithContext is the same as DeleteSubnet with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteSubnet for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteSubnetWithContext(ctx volcstack.Context, input *DeleteSubnetInput, opts ...request.Option) (*DeleteSubnetOutput, error) {
	req, out := c.DeleteSubnetRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteSubnetInput struct {
	_ struct{} `type:"structure"`

	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSubnetInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteSubnetInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSubnetInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteSubnetInput"}
	if s.SubnetId == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSubnetId sets the SubnetId field's value.
func (s *DeleteSubnetInput) SetSubnetId(v string) *DeleteSubnetInput {
	s.SubnetId = &v
	return s
}

type DeleteSubnetOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteSubnetOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteSubnetOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteSubnetOutput) SetRequestId(v string) *DeleteSubnetOutput {
	s.RequestId = &v
	return s
}
