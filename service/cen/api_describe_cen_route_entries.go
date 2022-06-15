// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDescribeCenRouteEntriesCommon = "DescribeCenRouteEntries"

// DescribeCenRouteEntriesCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeCenRouteEntriesCommon operation. The "output" return
// value will be populated with the DescribeCenRouteEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCenRouteEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCenRouteEntriesCommon Send returns without error.
//
// See DescribeCenRouteEntriesCommon for more information on using the DescribeCenRouteEntriesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCenRouteEntriesCommonRequest method.
//    req, resp := client.DescribeCenRouteEntriesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DescribeCenRouteEntriesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCenRouteEntriesCommon,
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

// DescribeCenRouteEntriesCommon API operation for CEN.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CEN's
// API operation DescribeCenRouteEntriesCommon for usage and error information.
func (c *CEN) DescribeCenRouteEntriesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCenRouteEntriesCommonRequest(input)
	return out, req.Send()
}

// DescribeCenRouteEntriesCommonWithContext is the same as DescribeCenRouteEntriesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCenRouteEntriesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCenRouteEntriesCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCenRouteEntriesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCenRouteEntries = "DescribeCenRouteEntries"

// DescribeCenRouteEntriesRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeCenRouteEntries operation. The "output" return
// value will be populated with the DescribeCenRouteEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCenRouteEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCenRouteEntriesCommon Send returns without error.
//
// See DescribeCenRouteEntries for more information on using the DescribeCenRouteEntries
// API call, and error handling.
//
//    // Example sending a request using the DescribeCenRouteEntriesRequest method.
//    req, resp := client.DescribeCenRouteEntriesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DescribeCenRouteEntriesRequest(input *DescribeCenRouteEntriesInput) (req *request.Request, output *DescribeCenRouteEntriesOutput) {
	op := &request.Operation{
		Name:       opDescribeCenRouteEntries,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCenRouteEntriesInput{}
	}

	output = &DescribeCenRouteEntriesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeCenRouteEntries API operation for CEN.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CEN's
// API operation DescribeCenRouteEntries for usage and error information.
func (c *CEN) DescribeCenRouteEntries(input *DescribeCenRouteEntriesInput) (*DescribeCenRouteEntriesOutput, error) {
	req, out := c.DescribeCenRouteEntriesRequest(input)
	return out, req.Send()
}

// DescribeCenRouteEntriesWithContext is the same as DescribeCenRouteEntries with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCenRouteEntries for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCenRouteEntriesWithContext(ctx volcstack.Context, input *DescribeCenRouteEntriesInput, opts ...request.Option) (*DescribeCenRouteEntriesOutput, error) {
	req, out := c.DescribeCenRouteEntriesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CenRouteEntryForDescribeCenRouteEntriesOutput struct {
	_ struct{} `type:"structure"`

	AsPath []*string `type:"list"`

	CenId *string `type:"string"`

	DestinationCidrBlock *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceRegionId *string `type:"string"`

	InstanceType *string `type:"string"`

	PublishStatus *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s CenRouteEntryForDescribeCenRouteEntriesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s CenRouteEntryForDescribeCenRouteEntriesOutput) GoString() string {
	return s.String()
}

// SetAsPath sets the AsPath field's value.
func (s *CenRouteEntryForDescribeCenRouteEntriesOutput) SetAsPath(v []*string) *CenRouteEntryForDescribeCenRouteEntriesOutput {
	s.AsPath = v
	return s
}

// SetCenId sets the CenId field's value.
func (s *CenRouteEntryForDescribeCenRouteEntriesOutput) SetCenId(v string) *CenRouteEntryForDescribeCenRouteEntriesOutput {
	s.CenId = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *CenRouteEntryForDescribeCenRouteEntriesOutput) SetDestinationCidrBlock(v string) *CenRouteEntryForDescribeCenRouteEntriesOutput {
	s.DestinationCidrBlock = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CenRouteEntryForDescribeCenRouteEntriesOutput) SetInstanceId(v string) *CenRouteEntryForDescribeCenRouteEntriesOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceRegionId sets the InstanceRegionId field's value.
func (s *CenRouteEntryForDescribeCenRouteEntriesOutput) SetInstanceRegionId(v string) *CenRouteEntryForDescribeCenRouteEntriesOutput {
	s.InstanceRegionId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *CenRouteEntryForDescribeCenRouteEntriesOutput) SetInstanceType(v string) *CenRouteEntryForDescribeCenRouteEntriesOutput {
	s.InstanceType = &v
	return s
}

// SetPublishStatus sets the PublishStatus field's value.
func (s *CenRouteEntryForDescribeCenRouteEntriesOutput) SetPublishStatus(v string) *CenRouteEntryForDescribeCenRouteEntriesOutput {
	s.PublishStatus = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *CenRouteEntryForDescribeCenRouteEntriesOutput) SetStatus(v string) *CenRouteEntryForDescribeCenRouteEntriesOutput {
	s.Status = &v
	return s
}

type DescribeCenRouteEntriesInput struct {
	_ struct{} `type:"structure"`

	CenId *string `type:"string"`

	DestinationCidrBlock *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceRegionId *string `type:"string"`

	InstanceType *string `type:"string"`
}

// String returns the string representation
func (s DescribeCenRouteEntriesInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCenRouteEntriesInput) GoString() string {
	return s.String()
}

// SetCenId sets the CenId field's value.
func (s *DescribeCenRouteEntriesInput) SetCenId(v string) *DescribeCenRouteEntriesInput {
	s.CenId = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *DescribeCenRouteEntriesInput) SetDestinationCidrBlock(v string) *DescribeCenRouteEntriesInput {
	s.DestinationCidrBlock = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeCenRouteEntriesInput) SetInstanceId(v string) *DescribeCenRouteEntriesInput {
	s.InstanceId = &v
	return s
}

// SetInstanceRegionId sets the InstanceRegionId field's value.
func (s *DescribeCenRouteEntriesInput) SetInstanceRegionId(v string) *DescribeCenRouteEntriesInput {
	s.InstanceRegionId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *DescribeCenRouteEntriesInput) SetInstanceType(v string) *DescribeCenRouteEntriesInput {
	s.InstanceType = &v
	return s
}

type DescribeCenRouteEntriesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CenRouteEntries []*CenRouteEntryForDescribeCenRouteEntriesOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCenRouteEntriesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCenRouteEntriesOutput) GoString() string {
	return s.String()
}

// SetCenRouteEntries sets the CenRouteEntries field's value.
func (s *DescribeCenRouteEntriesOutput) SetCenRouteEntries(v []*CenRouteEntryForDescribeCenRouteEntriesOutput) *DescribeCenRouteEntriesOutput {
	s.CenRouteEntries = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCenRouteEntriesOutput) SetPageNumber(v int64) *DescribeCenRouteEntriesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCenRouteEntriesOutput) SetPageSize(v int64) *DescribeCenRouteEntriesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeCenRouteEntriesOutput) SetTotalCount(v int64) *DescribeCenRouteEntriesOutput {
	s.TotalCount = &v
	return s
}
