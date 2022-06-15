// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDescribeAclsCommon = "DescribeAcls"

// DescribeAclsCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeAclsCommon operation. The "output" return
// value will be populated with the DescribeAclsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAclsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAclsCommon Send returns without error.
//
// See DescribeAclsCommon for more information on using the DescribeAclsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeAclsCommonRequest method.
//    req, resp := client.DescribeAclsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeAclsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeAclsCommon,
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

// DescribeAclsCommon API operation for CLB.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CLB's
// API operation DescribeAclsCommon for usage and error information.
func (c *CLB) DescribeAclsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeAclsCommonRequest(input)
	return out, req.Send()
}

// DescribeAclsCommonWithContext is the same as DescribeAclsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAclsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeAclsCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeAclsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeAcls = "DescribeAcls"

// DescribeAclsRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeAcls operation. The "output" return
// value will be populated with the DescribeAclsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAclsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAclsCommon Send returns without error.
//
// See DescribeAcls for more information on using the DescribeAcls
// API call, and error handling.
//
//    // Example sending a request using the DescribeAclsRequest method.
//    req, resp := client.DescribeAclsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeAclsRequest(input *DescribeAclsInput) (req *request.Request, output *DescribeAclsOutput) {
	op := &request.Operation{
		Name:       opDescribeAcls,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAclsInput{}
	}

	output = &DescribeAclsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeAcls API operation for CLB.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CLB's
// API operation DescribeAcls for usage and error information.
func (c *CLB) DescribeAcls(input *DescribeAclsInput) (*DescribeAclsOutput, error) {
	req, out := c.DescribeAclsRequest(input)
	return out, req.Send()
}

// DescribeAclsWithContext is the same as DescribeAcls with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAcls for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeAclsWithContext(ctx volcstack.Context, input *DescribeAclsInput, opts ...request.Option) (*DescribeAclsOutput, error) {
	req, out := c.DescribeAclsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AclForDescribeAclsOutput struct {
	_ struct{} `type:"structure"`

	AclEntryCount *int64 `type:"integer"`

	AclId *string `type:"string"`

	AclName *string `type:"string"`

	CreateTime *string `type:"string"`

	Description *string `type:"string"`

	Listeners []*string `type:"list"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s AclForDescribeAclsOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s AclForDescribeAclsOutput) GoString() string {
	return s.String()
}

// SetAclEntryCount sets the AclEntryCount field's value.
func (s *AclForDescribeAclsOutput) SetAclEntryCount(v int64) *AclForDescribeAclsOutput {
	s.AclEntryCount = &v
	return s
}

// SetAclId sets the AclId field's value.
func (s *AclForDescribeAclsOutput) SetAclId(v string) *AclForDescribeAclsOutput {
	s.AclId = &v
	return s
}

// SetAclName sets the AclName field's value.
func (s *AclForDescribeAclsOutput) SetAclName(v string) *AclForDescribeAclsOutput {
	s.AclName = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *AclForDescribeAclsOutput) SetCreateTime(v string) *AclForDescribeAclsOutput {
	s.CreateTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *AclForDescribeAclsOutput) SetDescription(v string) *AclForDescribeAclsOutput {
	s.Description = &v
	return s
}

// SetListeners sets the Listeners field's value.
func (s *AclForDescribeAclsOutput) SetListeners(v []*string) *AclForDescribeAclsOutput {
	s.Listeners = v
	return s
}

// SetStatus sets the Status field's value.
func (s *AclForDescribeAclsOutput) SetStatus(v string) *AclForDescribeAclsOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *AclForDescribeAclsOutput) SetUpdateTime(v string) *AclForDescribeAclsOutput {
	s.UpdateTime = &v
	return s
}

type DescribeAclsInput struct {
	_ struct{} `type:"structure"`

	AclIds []*string `type:"list"`

	AclName *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeAclsInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAclsInput) GoString() string {
	return s.String()
}

// SetAclIds sets the AclIds field's value.
func (s *DescribeAclsInput) SetAclIds(v []*string) *DescribeAclsInput {
	s.AclIds = v
	return s
}

// SetAclName sets the AclName field's value.
func (s *DescribeAclsInput) SetAclName(v string) *DescribeAclsInput {
	s.AclName = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeAclsInput) SetPageNumber(v int64) *DescribeAclsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeAclsInput) SetPageSize(v int64) *DescribeAclsInput {
	s.PageSize = &v
	return s
}

type DescribeAclsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Acls []*AclForDescribeAclsOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeAclsOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAclsOutput) GoString() string {
	return s.String()
}

// SetAcls sets the Acls field's value.
func (s *DescribeAclsOutput) SetAcls(v []*AclForDescribeAclsOutput) *DescribeAclsOutput {
	s.Acls = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeAclsOutput) SetPageNumber(v int64) *DescribeAclsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeAclsOutput) SetPageSize(v int64) *DescribeAclsOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeAclsOutput) SetRequestId(v string) *DescribeAclsOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeAclsOutput) SetTotalCount(v int64) *DescribeAclsOutput {
	s.TotalCount = &v
	return s
}
