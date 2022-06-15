// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opCreateSubnetCommon = "CreateSubnet"

// CreateSubnetCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the CreateSubnetCommon operation. The "output" return
// value will be populated with the CreateSubnetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateSubnetCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateSubnetCommon Send returns without error.
//
// See CreateSubnetCommon for more information on using the CreateSubnetCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateSubnetCommonRequest method.
//    req, resp := client.CreateSubnetCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateSubnetCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateSubnetCommon,
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

// CreateSubnetCommon API operation for VPC.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPC's
// API operation CreateSubnetCommon for usage and error information.
func (c *VPC) CreateSubnetCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateSubnetCommonRequest(input)
	return out, req.Send()
}

// CreateSubnetCommonWithContext is the same as CreateSubnetCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateSubnetCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateSubnetCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateSubnetCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateSubnet = "CreateSubnet"

// CreateSubnetRequest generates a "volcstack/request.Request" representing the
// client's request for the CreateSubnet operation. The "output" return
// value will be populated with the CreateSubnetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateSubnetCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateSubnetCommon Send returns without error.
//
// See CreateSubnet for more information on using the CreateSubnet
// API call, and error handling.
//
//    // Example sending a request using the CreateSubnetRequest method.
//    req, resp := client.CreateSubnetRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateSubnetRequest(input *CreateSubnetInput) (req *request.Request, output *CreateSubnetOutput) {
	op := &request.Operation{
		Name:       opCreateSubnet,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSubnetInput{}
	}

	output = &CreateSubnetOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateSubnet API operation for VPC.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPC's
// API operation CreateSubnet for usage and error information.
func (c *VPC) CreateSubnet(input *CreateSubnetInput) (*CreateSubnetOutput, error) {
	req, out := c.CreateSubnetRequest(input)
	return out, req.Send()
}

// CreateSubnetWithContext is the same as CreateSubnet with the addition of
// the ability to pass a context and additional request options.
//
// See CreateSubnet for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateSubnetWithContext(ctx volcstack.Context, input *CreateSubnetInput, opts ...request.Option) (*CreateSubnetOutput, error) {
	req, out := c.CreateSubnetRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateSubnetInput struct {
	_ struct{} `type:"structure"`

	// CidrBlock is a required field
	CidrBlock *string `type:"string" required:"true"`

	Description *string `min:"1" max:"255" type:"string"`

	SubnetName *string `min:"1" max:"128" type:"string"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`

	// ZoneId is a required field
	ZoneId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSubnetInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateSubnetInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSubnetInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateSubnetInput"}
	if s.CidrBlock == nil {
		invalidParams.Add(request.NewErrParamRequired("CidrBlock"))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.SubnetName != nil && len(*s.SubnetName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SubnetName", 1))
	}
	if s.SubnetName != nil && len(*s.SubnetName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("SubnetName", 128, *s.SubnetName))
	}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}
	if s.ZoneId == nil {
		invalidParams.Add(request.NewErrParamRequired("ZoneId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCidrBlock sets the CidrBlock field's value.
func (s *CreateSubnetInput) SetCidrBlock(v string) *CreateSubnetInput {
	s.CidrBlock = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateSubnetInput) SetDescription(v string) *CreateSubnetInput {
	s.Description = &v
	return s
}

// SetSubnetName sets the SubnetName field's value.
func (s *CreateSubnetInput) SetSubnetName(v string) *CreateSubnetInput {
	s.SubnetName = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *CreateSubnetInput) SetVpcId(v string) *CreateSubnetInput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *CreateSubnetInput) SetZoneId(v string) *CreateSubnetInput {
	s.ZoneId = &v
	return s
}

type CreateSubnetOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	SubnetId *string `type:"string"`
}

// String returns the string representation
func (s CreateSubnetOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateSubnetOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *CreateSubnetOutput) SetRequestId(v string) *CreateSubnetOutput {
	s.RequestId = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *CreateSubnetOutput) SetSubnetId(v string) *CreateSubnetOutput {
	s.SubnetId = &v
	return s
}
