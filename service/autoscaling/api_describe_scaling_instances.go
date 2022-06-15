// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDescribeScalingInstancesCommon = "DescribeScalingInstances"

// DescribeScalingInstancesCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeScalingInstancesCommon operation. The "output" return
// value will be populated with the DescribeScalingInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeScalingInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeScalingInstancesCommon Send returns without error.
//
// See DescribeScalingInstancesCommon for more information on using the DescribeScalingInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeScalingInstancesCommonRequest method.
//    req, resp := client.DescribeScalingInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DescribeScalingInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeScalingInstancesCommon,
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

// DescribeScalingInstancesCommon API operation for AUTO_SCALING.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for AUTO_SCALING's
// API operation DescribeScalingInstancesCommon for usage and error information.
func (c *AUTOSCALING) DescribeScalingInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeScalingInstancesCommonRequest(input)
	return out, req.Send()
}

// DescribeScalingInstancesCommonWithContext is the same as DescribeScalingInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeScalingInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeScalingInstancesCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeScalingInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeScalingInstances = "DescribeScalingInstances"

// DescribeScalingInstancesRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeScalingInstances operation. The "output" return
// value will be populated with the DescribeScalingInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeScalingInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeScalingInstancesCommon Send returns without error.
//
// See DescribeScalingInstances for more information on using the DescribeScalingInstances
// API call, and error handling.
//
//    // Example sending a request using the DescribeScalingInstancesRequest method.
//    req, resp := client.DescribeScalingInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DescribeScalingInstancesRequest(input *DescribeScalingInstancesInput) (req *request.Request, output *DescribeScalingInstancesOutput) {
	op := &request.Operation{
		Name:       opDescribeScalingInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeScalingInstancesInput{}
	}

	output = &DescribeScalingInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeScalingInstances API operation for AUTO_SCALING.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for AUTO_SCALING's
// API operation DescribeScalingInstances for usage and error information.
func (c *AUTOSCALING) DescribeScalingInstances(input *DescribeScalingInstancesInput) (*DescribeScalingInstancesOutput, error) {
	req, out := c.DescribeScalingInstancesRequest(input)
	return out, req.Send()
}

// DescribeScalingInstancesWithContext is the same as DescribeScalingInstances with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeScalingInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeScalingInstancesWithContext(ctx volcstack.Context, input *DescribeScalingInstancesInput, opts ...request.Option) (*DescribeScalingInstancesOutput, error) {
	req, out := c.DescribeScalingInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeScalingInstancesInput struct {
	_ struct{} `type:"structure"`

	CreationType *string `type:"string"`

	InstanceIds []*string `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ScalingConfigurationId *string `type:"string"`

	ScalingGroupId *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s DescribeScalingInstancesInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeScalingInstancesInput) GoString() string {
	return s.String()
}

// SetCreationType sets the CreationType field's value.
func (s *DescribeScalingInstancesInput) SetCreationType(v string) *DescribeScalingInstancesInput {
	s.CreationType = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *DescribeScalingInstancesInput) SetInstanceIds(v []*string) *DescribeScalingInstancesInput {
	s.InstanceIds = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeScalingInstancesInput) SetPageNumber(v int32) *DescribeScalingInstancesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeScalingInstancesInput) SetPageSize(v int32) *DescribeScalingInstancesInput {
	s.PageSize = &v
	return s
}

// SetScalingConfigurationId sets the ScalingConfigurationId field's value.
func (s *DescribeScalingInstancesInput) SetScalingConfigurationId(v string) *DescribeScalingInstancesInput {
	s.ScalingConfigurationId = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DescribeScalingInstancesInput) SetScalingGroupId(v string) *DescribeScalingInstancesInput {
	s.ScalingGroupId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeScalingInstancesInput) SetStatus(v string) *DescribeScalingInstancesInput {
	s.Status = &v
	return s
}

type DescribeScalingInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ScalingInstances []*ScalingInstanceForDescribeScalingInstancesOutput `type:"list"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeScalingInstancesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeScalingInstancesOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeScalingInstancesOutput) SetPageNumber(v int32) *DescribeScalingInstancesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeScalingInstancesOutput) SetPageSize(v int32) *DescribeScalingInstancesOutput {
	s.PageSize = &v
	return s
}

// SetScalingInstances sets the ScalingInstances field's value.
func (s *DescribeScalingInstancesOutput) SetScalingInstances(v []*ScalingInstanceForDescribeScalingInstancesOutput) *DescribeScalingInstancesOutput {
	s.ScalingInstances = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeScalingInstancesOutput) SetTotalCount(v int32) *DescribeScalingInstancesOutput {
	s.TotalCount = &v
	return s
}

type ScalingInstanceForDescribeScalingInstancesOutput struct {
	_ struct{} `type:"structure"`

	CreatedTime *string `type:"string"`

	CreationType *string `type:"string"`

	Entrusted *bool `type:"boolean"`

	InstanceId *string `type:"string"`

	ScalingConfigurationId *string `type:"string"`

	ScalingGroupId *string `type:"string"`

	ScalingPolicyId *string `type:"string"`

	Status *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s ScalingInstanceForDescribeScalingInstancesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ScalingInstanceForDescribeScalingInstancesOutput) GoString() string {
	return s.String()
}

// SetCreatedTime sets the CreatedTime field's value.
func (s *ScalingInstanceForDescribeScalingInstancesOutput) SetCreatedTime(v string) *ScalingInstanceForDescribeScalingInstancesOutput {
	s.CreatedTime = &v
	return s
}

// SetCreationType sets the CreationType field's value.
func (s *ScalingInstanceForDescribeScalingInstancesOutput) SetCreationType(v string) *ScalingInstanceForDescribeScalingInstancesOutput {
	s.CreationType = &v
	return s
}

// SetEntrusted sets the Entrusted field's value.
func (s *ScalingInstanceForDescribeScalingInstancesOutput) SetEntrusted(v bool) *ScalingInstanceForDescribeScalingInstancesOutput {
	s.Entrusted = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ScalingInstanceForDescribeScalingInstancesOutput) SetInstanceId(v string) *ScalingInstanceForDescribeScalingInstancesOutput {
	s.InstanceId = &v
	return s
}

// SetScalingConfigurationId sets the ScalingConfigurationId field's value.
func (s *ScalingInstanceForDescribeScalingInstancesOutput) SetScalingConfigurationId(v string) *ScalingInstanceForDescribeScalingInstancesOutput {
	s.ScalingConfigurationId = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *ScalingInstanceForDescribeScalingInstancesOutput) SetScalingGroupId(v string) *ScalingInstanceForDescribeScalingInstancesOutput {
	s.ScalingGroupId = &v
	return s
}

// SetScalingPolicyId sets the ScalingPolicyId field's value.
func (s *ScalingInstanceForDescribeScalingInstancesOutput) SetScalingPolicyId(v string) *ScalingInstanceForDescribeScalingInstancesOutput {
	s.ScalingPolicyId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ScalingInstanceForDescribeScalingInstancesOutput) SetStatus(v string) *ScalingInstanceForDescribeScalingInstancesOutput {
	s.Status = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ScalingInstanceForDescribeScalingInstancesOutput) SetZoneId(v string) *ScalingInstanceForDescribeScalingInstancesOutput {
	s.ZoneId = &v
	return s
}
