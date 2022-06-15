// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDescribeScalingActivitiesCommon = "DescribeScalingActivities"

// DescribeScalingActivitiesCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeScalingActivitiesCommon operation. The "output" return
// value will be populated with the DescribeScalingActivitiesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeScalingActivitiesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeScalingActivitiesCommon Send returns without error.
//
// See DescribeScalingActivitiesCommon for more information on using the DescribeScalingActivitiesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeScalingActivitiesCommonRequest method.
//    req, resp := client.DescribeScalingActivitiesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DescribeScalingActivitiesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeScalingActivitiesCommon,
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

// DescribeScalingActivitiesCommon API operation for AUTO_SCALING.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for AUTO_SCALING's
// API operation DescribeScalingActivitiesCommon for usage and error information.
func (c *AUTOSCALING) DescribeScalingActivitiesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeScalingActivitiesCommonRequest(input)
	return out, req.Send()
}

// DescribeScalingActivitiesCommonWithContext is the same as DescribeScalingActivitiesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeScalingActivitiesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeScalingActivitiesCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeScalingActivitiesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeScalingActivities = "DescribeScalingActivities"

// DescribeScalingActivitiesRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeScalingActivities operation. The "output" return
// value will be populated with the DescribeScalingActivitiesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeScalingActivitiesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeScalingActivitiesCommon Send returns without error.
//
// See DescribeScalingActivities for more information on using the DescribeScalingActivities
// API call, and error handling.
//
//    // Example sending a request using the DescribeScalingActivitiesRequest method.
//    req, resp := client.DescribeScalingActivitiesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DescribeScalingActivitiesRequest(input *DescribeScalingActivitiesInput) (req *request.Request, output *DescribeScalingActivitiesOutput) {
	op := &request.Operation{
		Name:       opDescribeScalingActivities,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeScalingActivitiesInput{}
	}

	output = &DescribeScalingActivitiesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeScalingActivities API operation for AUTO_SCALING.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for AUTO_SCALING's
// API operation DescribeScalingActivities for usage and error information.
func (c *AUTOSCALING) DescribeScalingActivities(input *DescribeScalingActivitiesInput) (*DescribeScalingActivitiesOutput, error) {
	req, out := c.DescribeScalingActivitiesRequest(input)
	return out, req.Send()
}

// DescribeScalingActivitiesWithContext is the same as DescribeScalingActivities with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeScalingActivities for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeScalingActivitiesWithContext(ctx volcstack.Context, input *DescribeScalingActivitiesInput, opts ...request.Option) (*DescribeScalingActivitiesOutput, error) {
	req, out := c.DescribeScalingActivitiesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeScalingActivitiesInput struct {
	_ struct{} `type:"structure"`

	EndTime *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ScalingActivityIds []*string `type:"list"`

	ScalingGroupId *string `type:"string"`

	StartTime *string `type:"string"`

	StatusCode *string `type:"string"`
}

// String returns the string representation
func (s DescribeScalingActivitiesInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeScalingActivitiesInput) GoString() string {
	return s.String()
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeScalingActivitiesInput) SetEndTime(v string) *DescribeScalingActivitiesInput {
	s.EndTime = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeScalingActivitiesInput) SetPageNumber(v int32) *DescribeScalingActivitiesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeScalingActivitiesInput) SetPageSize(v int32) *DescribeScalingActivitiesInput {
	s.PageSize = &v
	return s
}

// SetScalingActivityIds sets the ScalingActivityIds field's value.
func (s *DescribeScalingActivitiesInput) SetScalingActivityIds(v []*string) *DescribeScalingActivitiesInput {
	s.ScalingActivityIds = v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DescribeScalingActivitiesInput) SetScalingGroupId(v string) *DescribeScalingActivitiesInput {
	s.ScalingGroupId = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeScalingActivitiesInput) SetStartTime(v string) *DescribeScalingActivitiesInput {
	s.StartTime = &v
	return s
}

// SetStatusCode sets the StatusCode field's value.
func (s *DescribeScalingActivitiesInput) SetStatusCode(v string) *DescribeScalingActivitiesInput {
	s.StatusCode = &v
	return s
}

type DescribeScalingActivitiesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ScalingActivities []*ScalingActivityForDescribeScalingActivitiesOutput `type:"list"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeScalingActivitiesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeScalingActivitiesOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeScalingActivitiesOutput) SetPageNumber(v int32) *DescribeScalingActivitiesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeScalingActivitiesOutput) SetPageSize(v int32) *DescribeScalingActivitiesOutput {
	s.PageSize = &v
	return s
}

// SetScalingActivities sets the ScalingActivities field's value.
func (s *DescribeScalingActivitiesOutput) SetScalingActivities(v []*ScalingActivityForDescribeScalingActivitiesOutput) *DescribeScalingActivitiesOutput {
	s.ScalingActivities = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeScalingActivitiesOutput) SetTotalCount(v int32) *DescribeScalingActivitiesOutput {
	s.TotalCount = &v
	return s
}

type RelatedInstanceForDescribeScalingActivitiesOutput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	Message *string `type:"string"`

	OperateType *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s RelatedInstanceForDescribeScalingActivitiesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s RelatedInstanceForDescribeScalingActivitiesOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *RelatedInstanceForDescribeScalingActivitiesOutput) SetInstanceId(v string) *RelatedInstanceForDescribeScalingActivitiesOutput {
	s.InstanceId = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *RelatedInstanceForDescribeScalingActivitiesOutput) SetMessage(v string) *RelatedInstanceForDescribeScalingActivitiesOutput {
	s.Message = &v
	return s
}

// SetOperateType sets the OperateType field's value.
func (s *RelatedInstanceForDescribeScalingActivitiesOutput) SetOperateType(v string) *RelatedInstanceForDescribeScalingActivitiesOutput {
	s.OperateType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *RelatedInstanceForDescribeScalingActivitiesOutput) SetStatus(v string) *RelatedInstanceForDescribeScalingActivitiesOutput {
	s.Status = &v
	return s
}

type ScalingActivityForDescribeScalingActivitiesOutput struct {
	_ struct{} `type:"structure"`

	ActivityType *string `type:"string"`

	ActualAdjustInstanceNumber *int32 `type:"int32"`

	Cooldown *int32 `type:"int32"`

	CreatedAt *string `type:"string"`

	CurrentInstanceNumber *int32 `type:"int32"`

	ExpectedRunTime *string `type:"string"`

	MaxInstanceNumber *int32 `type:"int32"`

	MinInstanceNumber *int32 `type:"int32"`

	RelatedInstances []*RelatedInstanceForDescribeScalingActivitiesOutput `type:"list"`

	ResultMsg *string `type:"string"`

	ScalingActivityId *string `type:"string"`

	ScalingGroupId *string `type:"string"`

	StatusCode *string `type:"string"`

	StoppedAt *string `type:"string"`

	TaskCategory *string `type:"string"`
}

// String returns the string representation
func (s ScalingActivityForDescribeScalingActivitiesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ScalingActivityForDescribeScalingActivitiesOutput) GoString() string {
	return s.String()
}

// SetActivityType sets the ActivityType field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetActivityType(v string) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.ActivityType = &v
	return s
}

// SetActualAdjustInstanceNumber sets the ActualAdjustInstanceNumber field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetActualAdjustInstanceNumber(v int32) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.ActualAdjustInstanceNumber = &v
	return s
}

// SetCooldown sets the Cooldown field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetCooldown(v int32) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.Cooldown = &v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetCreatedAt(v string) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.CreatedAt = &v
	return s
}

// SetCurrentInstanceNumber sets the CurrentInstanceNumber field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetCurrentInstanceNumber(v int32) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.CurrentInstanceNumber = &v
	return s
}

// SetExpectedRunTime sets the ExpectedRunTime field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetExpectedRunTime(v string) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.ExpectedRunTime = &v
	return s
}

// SetMaxInstanceNumber sets the MaxInstanceNumber field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetMaxInstanceNumber(v int32) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.MaxInstanceNumber = &v
	return s
}

// SetMinInstanceNumber sets the MinInstanceNumber field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetMinInstanceNumber(v int32) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.MinInstanceNumber = &v
	return s
}

// SetRelatedInstances sets the RelatedInstances field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetRelatedInstances(v []*RelatedInstanceForDescribeScalingActivitiesOutput) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.RelatedInstances = v
	return s
}

// SetResultMsg sets the ResultMsg field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetResultMsg(v string) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.ResultMsg = &v
	return s
}

// SetScalingActivityId sets the ScalingActivityId field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetScalingActivityId(v string) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.ScalingActivityId = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetScalingGroupId(v string) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.ScalingGroupId = &v
	return s
}

// SetStatusCode sets the StatusCode field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetStatusCode(v string) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.StatusCode = &v
	return s
}

// SetStoppedAt sets the StoppedAt field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetStoppedAt(v string) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.StoppedAt = &v
	return s
}

// SetTaskCategory sets the TaskCategory field's value.
func (s *ScalingActivityForDescribeScalingActivitiesOutput) SetTaskCategory(v string) *ScalingActivityForDescribeScalingActivitiesOutput {
	s.TaskCategory = &v
	return s
}
