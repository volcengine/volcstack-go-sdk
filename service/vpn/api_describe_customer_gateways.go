// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDescribeCustomerGatewaysCommon = "DescribeCustomerGateways"

// DescribeCustomerGatewaysCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeCustomerGatewaysCommon operation. The "output" return
// value will be populated with the DescribeCustomerGatewaysCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCustomerGatewaysCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCustomerGatewaysCommon Send returns without error.
//
// See DescribeCustomerGatewaysCommon for more information on using the DescribeCustomerGatewaysCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCustomerGatewaysCommonRequest method.
//    req, resp := client.DescribeCustomerGatewaysCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DescribeCustomerGatewaysCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCustomerGatewaysCommon,
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

// DescribeCustomerGatewaysCommon API operation for VPN.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPN's
// API operation DescribeCustomerGatewaysCommon for usage and error information.
func (c *VPN) DescribeCustomerGatewaysCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCustomerGatewaysCommonRequest(input)
	return out, req.Send()
}

// DescribeCustomerGatewaysCommonWithContext is the same as DescribeCustomerGatewaysCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCustomerGatewaysCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DescribeCustomerGatewaysCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCustomerGatewaysCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCustomerGateways = "DescribeCustomerGateways"

// DescribeCustomerGatewaysRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeCustomerGateways operation. The "output" return
// value will be populated with the DescribeCustomerGatewaysCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCustomerGatewaysCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCustomerGatewaysCommon Send returns without error.
//
// See DescribeCustomerGateways for more information on using the DescribeCustomerGateways
// API call, and error handling.
//
//    // Example sending a request using the DescribeCustomerGatewaysRequest method.
//    req, resp := client.DescribeCustomerGatewaysRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DescribeCustomerGatewaysRequest(input *DescribeCustomerGatewaysInput) (req *request.Request, output *DescribeCustomerGatewaysOutput) {
	op := &request.Operation{
		Name:       opDescribeCustomerGateways,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCustomerGatewaysInput{}
	}

	output = &DescribeCustomerGatewaysOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeCustomerGateways API operation for VPN.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPN's
// API operation DescribeCustomerGateways for usage and error information.
func (c *VPN) DescribeCustomerGateways(input *DescribeCustomerGatewaysInput) (*DescribeCustomerGatewaysOutput, error) {
	req, out := c.DescribeCustomerGatewaysRequest(input)
	return out, req.Send()
}

// DescribeCustomerGatewaysWithContext is the same as DescribeCustomerGateways with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCustomerGateways for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DescribeCustomerGatewaysWithContext(ctx volcstack.Context, input *DescribeCustomerGatewaysInput, opts ...request.Option) (*DescribeCustomerGatewaysOutput, error) {
	req, out := c.DescribeCustomerGatewaysRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CustomerGatewayForDescribeCustomerGatewaysOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	ConnectionCount *int64 `type:"integer"`

	CreationTime *string `type:"string"`

	CustomerGatewayId *string `type:"string"`

	CustomerGatewayName *string `type:"string"`

	Description *string `type:"string"`

	IpAddress *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s CustomerGatewayForDescribeCustomerGatewaysOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s CustomerGatewayForDescribeCustomerGatewaysOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *CustomerGatewayForDescribeCustomerGatewaysOutput) SetAccountId(v string) *CustomerGatewayForDescribeCustomerGatewaysOutput {
	s.AccountId = &v
	return s
}

// SetConnectionCount sets the ConnectionCount field's value.
func (s *CustomerGatewayForDescribeCustomerGatewaysOutput) SetConnectionCount(v int64) *CustomerGatewayForDescribeCustomerGatewaysOutput {
	s.ConnectionCount = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *CustomerGatewayForDescribeCustomerGatewaysOutput) SetCreationTime(v string) *CustomerGatewayForDescribeCustomerGatewaysOutput {
	s.CreationTime = &v
	return s
}

// SetCustomerGatewayId sets the CustomerGatewayId field's value.
func (s *CustomerGatewayForDescribeCustomerGatewaysOutput) SetCustomerGatewayId(v string) *CustomerGatewayForDescribeCustomerGatewaysOutput {
	s.CustomerGatewayId = &v
	return s
}

// SetCustomerGatewayName sets the CustomerGatewayName field's value.
func (s *CustomerGatewayForDescribeCustomerGatewaysOutput) SetCustomerGatewayName(v string) *CustomerGatewayForDescribeCustomerGatewaysOutput {
	s.CustomerGatewayName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CustomerGatewayForDescribeCustomerGatewaysOutput) SetDescription(v string) *CustomerGatewayForDescribeCustomerGatewaysOutput {
	s.Description = &v
	return s
}

// SetIpAddress sets the IpAddress field's value.
func (s *CustomerGatewayForDescribeCustomerGatewaysOutput) SetIpAddress(v string) *CustomerGatewayForDescribeCustomerGatewaysOutput {
	s.IpAddress = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *CustomerGatewayForDescribeCustomerGatewaysOutput) SetStatus(v string) *CustomerGatewayForDescribeCustomerGatewaysOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *CustomerGatewayForDescribeCustomerGatewaysOutput) SetUpdateTime(v string) *CustomerGatewayForDescribeCustomerGatewaysOutput {
	s.UpdateTime = &v
	return s
}

type DescribeCustomerGatewaysInput struct {
	_ struct{} `type:"structure"`

	CustomerGatewayIds []*string `type:"list"`

	CustomerGatewayName *string `type:"string"`

	IPAddress *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s DescribeCustomerGatewaysInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCustomerGatewaysInput) GoString() string {
	return s.String()
}

// SetCustomerGatewayIds sets the CustomerGatewayIds field's value.
func (s *DescribeCustomerGatewaysInput) SetCustomerGatewayIds(v []*string) *DescribeCustomerGatewaysInput {
	s.CustomerGatewayIds = v
	return s
}

// SetCustomerGatewayName sets the CustomerGatewayName field's value.
func (s *DescribeCustomerGatewaysInput) SetCustomerGatewayName(v string) *DescribeCustomerGatewaysInput {
	s.CustomerGatewayName = &v
	return s
}

// SetIPAddress sets the IPAddress field's value.
func (s *DescribeCustomerGatewaysInput) SetIPAddress(v string) *DescribeCustomerGatewaysInput {
	s.IPAddress = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCustomerGatewaysInput) SetPageNumber(v int64) *DescribeCustomerGatewaysInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCustomerGatewaysInput) SetPageSize(v int64) *DescribeCustomerGatewaysInput {
	s.PageSize = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeCustomerGatewaysInput) SetStatus(v string) *DescribeCustomerGatewaysInput {
	s.Status = &v
	return s
}

type DescribeCustomerGatewaysOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CustomerGateways []*CustomerGatewayForDescribeCustomerGatewaysOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCustomerGatewaysOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCustomerGatewaysOutput) GoString() string {
	return s.String()
}

// SetCustomerGateways sets the CustomerGateways field's value.
func (s *DescribeCustomerGatewaysOutput) SetCustomerGateways(v []*CustomerGatewayForDescribeCustomerGatewaysOutput) *DescribeCustomerGatewaysOutput {
	s.CustomerGateways = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCustomerGatewaysOutput) SetPageNumber(v int64) *DescribeCustomerGatewaysOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCustomerGatewaysOutput) SetPageSize(v int64) *DescribeCustomerGatewaysOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeCustomerGatewaysOutput) SetRequestId(v string) *DescribeCustomerGatewaysOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeCustomerGatewaysOutput) SetTotalCount(v int64) *DescribeCustomerGatewaysOutput {
	s.TotalCount = &v
	return s
}
