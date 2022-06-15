// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"encoding/json"

	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDescribeVolumesCommon = "DescribeVolumes"

// DescribeVolumesCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeVolumesCommon operation. The "output" return
// value will be populated with the DescribeVolumesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVolumesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVolumesCommon Send returns without error.
//
// See DescribeVolumesCommon for more information on using the DescribeVolumesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeVolumesCommonRequest method.
//    req, resp := client.DescribeVolumesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) DescribeVolumesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVolumesCommon,
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

// DescribeVolumesCommon API operation for STORAGE_EBS.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for STORAGE_EBS's
// API operation DescribeVolumesCommon for usage and error information.
func (c *STORAGEEBS) DescribeVolumesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVolumesCommonRequest(input)
	return out, req.Send()
}

// DescribeVolumesCommonWithContext is the same as DescribeVolumesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVolumesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) DescribeVolumesCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVolumesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVolumes = "DescribeVolumes"

// DescribeVolumesRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeVolumes operation. The "output" return
// value will be populated with the DescribeVolumesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVolumesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVolumesCommon Send returns without error.
//
// See DescribeVolumes for more information on using the DescribeVolumes
// API call, and error handling.
//
//    // Example sending a request using the DescribeVolumesRequest method.
//    req, resp := client.DescribeVolumesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STORAGEEBS) DescribeVolumesRequest(input *DescribeVolumesInput) (req *request.Request, output *DescribeVolumesOutput) {
	op := &request.Operation{
		Name:       opDescribeVolumes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVolumesInput{}
	}

	output = &DescribeVolumesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeVolumes API operation for STORAGE_EBS.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for STORAGE_EBS's
// API operation DescribeVolumes for usage and error information.
func (c *STORAGEEBS) DescribeVolumes(input *DescribeVolumesInput) (*DescribeVolumesOutput, error) {
	req, out := c.DescribeVolumesRequest(input)
	return out, req.Send()
}

// DescribeVolumesWithContext is the same as DescribeVolumes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVolumes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) DescribeVolumesWithContext(ctx volcstack.Context, input *DescribeVolumesInput, opts ...request.Option) (*DescribeVolumesOutput, error) {
	req, out := c.DescribeVolumesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeVolumesInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	Kind *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	VolumeIds []*string `type:"list"`

	VolumeName *string `type:"string"`

	VolumeStatus *string `type:"string"`

	VolumeType *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeVolumesInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVolumesInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeVolumesInput) SetInstanceId(v string) *DescribeVolumesInput {
	s.InstanceId = &v
	return s
}

// SetKind sets the Kind field's value.
func (s *DescribeVolumesInput) SetKind(v string) *DescribeVolumesInput {
	s.Kind = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVolumesInput) SetPageNumber(v int32) *DescribeVolumesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVolumesInput) SetPageSize(v int32) *DescribeVolumesInput {
	s.PageSize = &v
	return s
}

// SetVolumeIds sets the VolumeIds field's value.
func (s *DescribeVolumesInput) SetVolumeIds(v []*string) *DescribeVolumesInput {
	s.VolumeIds = v
	return s
}

// SetVolumeName sets the VolumeName field's value.
func (s *DescribeVolumesInput) SetVolumeName(v string) *DescribeVolumesInput {
	s.VolumeName = &v
	return s
}

// SetVolumeStatus sets the VolumeStatus field's value.
func (s *DescribeVolumesInput) SetVolumeStatus(v string) *DescribeVolumesInput {
	s.VolumeStatus = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *DescribeVolumesInput) SetVolumeType(v string) *DescribeVolumesInput {
	s.VolumeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeVolumesInput) SetZoneId(v string) *DescribeVolumesInput {
	s.ZoneId = &v
	return s
}

type DescribeVolumesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`

	Volumes []*VolumeForDescribeVolumesOutput `type:"list"`
}

// String returns the string representation
func (s DescribeVolumesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVolumesOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVolumesOutput) SetPageNumber(v int32) *DescribeVolumesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVolumesOutput) SetPageSize(v int32) *DescribeVolumesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeVolumesOutput) SetTotalCount(v int32) *DescribeVolumesOutput {
	s.TotalCount = &v
	return s
}

// SetVolumes sets the Volumes field's value.
func (s *DescribeVolumesOutput) SetVolumes(v []*VolumeForDescribeVolumesOutput) *DescribeVolumesOutput {
	s.Volumes = v
	return s
}

type VolumeForDescribeVolumesOutput struct {
	_ struct{} `type:"structure"`

	BillingType *int32 `type:"int32"`

	CreatedAt *string `type:"string"`

	DeleteWithInstance *bool `type:"boolean"`

	Description *string `type:"string"`

	DeviceName *string `type:"string"`

	ExpiredTime *string `type:"string"`

	ImageId *string `type:"string"`

	InstanceId *string `type:"string"`

	Kind *string `type:"string"`

	PayType *string `type:"string"`

	RenewType *int32 `type:"int32"`

	Size *json.Number `type:"json_number"`

	Status *string `type:"string"`

	TradeStatus *int32 `type:"int32"`

	UpdatedAt *string `type:"string"`

	VolumeId *string `type:"string"`

	VolumeName *string `type:"string"`

	VolumeType *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s VolumeForDescribeVolumesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s VolumeForDescribeVolumesOutput) GoString() string {
	return s.String()
}

// SetBillingType sets the BillingType field's value.
func (s *VolumeForDescribeVolumesOutput) SetBillingType(v int32) *VolumeForDescribeVolumesOutput {
	s.BillingType = &v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *VolumeForDescribeVolumesOutput) SetCreatedAt(v string) *VolumeForDescribeVolumesOutput {
	s.CreatedAt = &v
	return s
}

// SetDeleteWithInstance sets the DeleteWithInstance field's value.
func (s *VolumeForDescribeVolumesOutput) SetDeleteWithInstance(v bool) *VolumeForDescribeVolumesOutput {
	s.DeleteWithInstance = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *VolumeForDescribeVolumesOutput) SetDescription(v string) *VolumeForDescribeVolumesOutput {
	s.Description = &v
	return s
}

// SetDeviceName sets the DeviceName field's value.
func (s *VolumeForDescribeVolumesOutput) SetDeviceName(v string) *VolumeForDescribeVolumesOutput {
	s.DeviceName = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *VolumeForDescribeVolumesOutput) SetExpiredTime(v string) *VolumeForDescribeVolumesOutput {
	s.ExpiredTime = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *VolumeForDescribeVolumesOutput) SetImageId(v string) *VolumeForDescribeVolumesOutput {
	s.ImageId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *VolumeForDescribeVolumesOutput) SetInstanceId(v string) *VolumeForDescribeVolumesOutput {
	s.InstanceId = &v
	return s
}

// SetKind sets the Kind field's value.
func (s *VolumeForDescribeVolumesOutput) SetKind(v string) *VolumeForDescribeVolumesOutput {
	s.Kind = &v
	return s
}

// SetPayType sets the PayType field's value.
func (s *VolumeForDescribeVolumesOutput) SetPayType(v string) *VolumeForDescribeVolumesOutput {
	s.PayType = &v
	return s
}

// SetRenewType sets the RenewType field's value.
func (s *VolumeForDescribeVolumesOutput) SetRenewType(v int32) *VolumeForDescribeVolumesOutput {
	s.RenewType = &v
	return s
}

// SetSize sets the Size field's value.
func (s *VolumeForDescribeVolumesOutput) SetSize(v json.Number) *VolumeForDescribeVolumesOutput {
	s.Size = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *VolumeForDescribeVolumesOutput) SetStatus(v string) *VolumeForDescribeVolumesOutput {
	s.Status = &v
	return s
}

// SetTradeStatus sets the TradeStatus field's value.
func (s *VolumeForDescribeVolumesOutput) SetTradeStatus(v int32) *VolumeForDescribeVolumesOutput {
	s.TradeStatus = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *VolumeForDescribeVolumesOutput) SetUpdatedAt(v string) *VolumeForDescribeVolumesOutput {
	s.UpdatedAt = &v
	return s
}

// SetVolumeId sets the VolumeId field's value.
func (s *VolumeForDescribeVolumesOutput) SetVolumeId(v string) *VolumeForDescribeVolumesOutput {
	s.VolumeId = &v
	return s
}

// SetVolumeName sets the VolumeName field's value.
func (s *VolumeForDescribeVolumesOutput) SetVolumeName(v string) *VolumeForDescribeVolumesOutput {
	s.VolumeName = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *VolumeForDescribeVolumesOutput) SetVolumeType(v string) *VolumeForDescribeVolumesOutput {
	s.VolumeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *VolumeForDescribeVolumesOutput) SetZoneId(v string) *VolumeForDescribeVolumesOutput {
	s.ZoneId = &v
	return s
}