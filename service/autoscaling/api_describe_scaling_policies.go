// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDescribeScalingPoliciesCommon = "DescribeScalingPolicies"

// DescribeScalingPoliciesCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeScalingPoliciesCommon operation. The "output" return
// value will be populated with the DescribeScalingPoliciesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeScalingPoliciesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeScalingPoliciesCommon Send returns without error.
//
// See DescribeScalingPoliciesCommon for more information on using the DescribeScalingPoliciesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeScalingPoliciesCommonRequest method.
//    req, resp := client.DescribeScalingPoliciesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DescribeScalingPoliciesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeScalingPoliciesCommon,
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

// DescribeScalingPoliciesCommon API operation for AUTO_SCALING.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for AUTO_SCALING's
// API operation DescribeScalingPoliciesCommon for usage and error information.
func (c *AUTOSCALING) DescribeScalingPoliciesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeScalingPoliciesCommonRequest(input)
	return out, req.Send()
}

// DescribeScalingPoliciesCommonWithContext is the same as DescribeScalingPoliciesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeScalingPoliciesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeScalingPoliciesCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeScalingPoliciesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeScalingPolicies = "DescribeScalingPolicies"

// DescribeScalingPoliciesRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeScalingPolicies operation. The "output" return
// value will be populated with the DescribeScalingPoliciesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeScalingPoliciesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeScalingPoliciesCommon Send returns without error.
//
// See DescribeScalingPolicies for more information on using the DescribeScalingPolicies
// API call, and error handling.
//
//    // Example sending a request using the DescribeScalingPoliciesRequest method.
//    req, resp := client.DescribeScalingPoliciesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DescribeScalingPoliciesRequest(input *DescribeScalingPoliciesInput) (req *request.Request, output *DescribeScalingPoliciesOutput) {
	op := &request.Operation{
		Name:       opDescribeScalingPolicies,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeScalingPoliciesInput{}
	}

	output = &DescribeScalingPoliciesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeScalingPolicies API operation for AUTO_SCALING.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for AUTO_SCALING's
// API operation DescribeScalingPolicies for usage and error information.
func (c *AUTOSCALING) DescribeScalingPolicies(input *DescribeScalingPoliciesInput) (*DescribeScalingPoliciesOutput, error) {
	req, out := c.DescribeScalingPoliciesRequest(input)
	return out, req.Send()
}

// DescribeScalingPoliciesWithContext is the same as DescribeScalingPolicies with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeScalingPolicies for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeScalingPoliciesWithContext(ctx volcstack.Context, input *DescribeScalingPoliciesInput, opts ...request.Option) (*DescribeScalingPoliciesOutput, error) {
	req, out := c.DescribeScalingPoliciesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AlarmPolicyForDescribeScalingPoliciesOutput struct {
	_ struct{} `type:"structure"`

	Condition *ConditionForDescribeScalingPoliciesOutput `type:"structure"`

	EvaluationCount *int32 `type:"int32"`

	RuleType *string `type:"string"`
}

// String returns the string representation
func (s AlarmPolicyForDescribeScalingPoliciesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s AlarmPolicyForDescribeScalingPoliciesOutput) GoString() string {
	return s.String()
}

// SetCondition sets the Condition field's value.
func (s *AlarmPolicyForDescribeScalingPoliciesOutput) SetCondition(v *ConditionForDescribeScalingPoliciesOutput) *AlarmPolicyForDescribeScalingPoliciesOutput {
	s.Condition = v
	return s
}

// SetEvaluationCount sets the EvaluationCount field's value.
func (s *AlarmPolicyForDescribeScalingPoliciesOutput) SetEvaluationCount(v int32) *AlarmPolicyForDescribeScalingPoliciesOutput {
	s.EvaluationCount = &v
	return s
}

// SetRuleType sets the RuleType field's value.
func (s *AlarmPolicyForDescribeScalingPoliciesOutput) SetRuleType(v string) *AlarmPolicyForDescribeScalingPoliciesOutput {
	s.RuleType = &v
	return s
}

type ConditionForDescribeScalingPoliciesOutput struct {
	_ struct{} `type:"structure"`

	ComparisonOperator *string `type:"string"`

	MetricName *string `type:"string"`

	MetricUnit *string `type:"string"`

	Threshold *string `type:"string"`
}

// String returns the string representation
func (s ConditionForDescribeScalingPoliciesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ConditionForDescribeScalingPoliciesOutput) GoString() string {
	return s.String()
}

// SetComparisonOperator sets the ComparisonOperator field's value.
func (s *ConditionForDescribeScalingPoliciesOutput) SetComparisonOperator(v string) *ConditionForDescribeScalingPoliciesOutput {
	s.ComparisonOperator = &v
	return s
}

// SetMetricName sets the MetricName field's value.
func (s *ConditionForDescribeScalingPoliciesOutput) SetMetricName(v string) *ConditionForDescribeScalingPoliciesOutput {
	s.MetricName = &v
	return s
}

// SetMetricUnit sets the MetricUnit field's value.
func (s *ConditionForDescribeScalingPoliciesOutput) SetMetricUnit(v string) *ConditionForDescribeScalingPoliciesOutput {
	s.MetricUnit = &v
	return s
}

// SetThreshold sets the Threshold field's value.
func (s *ConditionForDescribeScalingPoliciesOutput) SetThreshold(v string) *ConditionForDescribeScalingPoliciesOutput {
	s.Threshold = &v
	return s
}

type DescribeScalingPoliciesInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ScalingGroupId *string `type:"string"`

	ScalingPolicyIds []*string `type:"list"`

	ScalingPolicyNames []*string `type:"list"`

	ScalingPolicyType *string `type:"string"`
}

// String returns the string representation
func (s DescribeScalingPoliciesInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeScalingPoliciesInput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeScalingPoliciesInput) SetPageNumber(v int32) *DescribeScalingPoliciesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeScalingPoliciesInput) SetPageSize(v int32) *DescribeScalingPoliciesInput {
	s.PageSize = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DescribeScalingPoliciesInput) SetScalingGroupId(v string) *DescribeScalingPoliciesInput {
	s.ScalingGroupId = &v
	return s
}

// SetScalingPolicyIds sets the ScalingPolicyIds field's value.
func (s *DescribeScalingPoliciesInput) SetScalingPolicyIds(v []*string) *DescribeScalingPoliciesInput {
	s.ScalingPolicyIds = v
	return s
}

// SetScalingPolicyNames sets the ScalingPolicyNames field's value.
func (s *DescribeScalingPoliciesInput) SetScalingPolicyNames(v []*string) *DescribeScalingPoliciesInput {
	s.ScalingPolicyNames = v
	return s
}

// SetScalingPolicyType sets the ScalingPolicyType field's value.
func (s *DescribeScalingPoliciesInput) SetScalingPolicyType(v string) *DescribeScalingPoliciesInput {
	s.ScalingPolicyType = &v
	return s
}

type DescribeScalingPoliciesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ScalingPolicies []*ScalingPolicyForDescribeScalingPoliciesOutput `type:"list"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeScalingPoliciesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeScalingPoliciesOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeScalingPoliciesOutput) SetPageNumber(v int32) *DescribeScalingPoliciesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeScalingPoliciesOutput) SetPageSize(v int32) *DescribeScalingPoliciesOutput {
	s.PageSize = &v
	return s
}

// SetScalingPolicies sets the ScalingPolicies field's value.
func (s *DescribeScalingPoliciesOutput) SetScalingPolicies(v []*ScalingPolicyForDescribeScalingPoliciesOutput) *DescribeScalingPoliciesOutput {
	s.ScalingPolicies = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeScalingPoliciesOutput) SetTotalCount(v int32) *DescribeScalingPoliciesOutput {
	s.TotalCount = &v
	return s
}

type ScalingPolicyForDescribeScalingPoliciesOutput struct {
	_ struct{} `type:"structure"`

	AdjustmentType *string `type:"string"`

	AdjustmentValue *int32 `type:"int32"`

	AlarmPolicy *AlarmPolicyForDescribeScalingPoliciesOutput `type:"structure"`

	Cooldown *int32 `type:"int32"`

	ScalingGroupId *string `type:"string"`

	ScalingPolicyId *string `type:"string"`

	ScalingPolicyName *string `type:"string"`

	ScalingPolicyType *string `type:"string"`

	ScheduledPolicy *ScheduledPolicyForDescribeScalingPoliciesOutput `type:"structure"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s ScalingPolicyForDescribeScalingPoliciesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ScalingPolicyForDescribeScalingPoliciesOutput) GoString() string {
	return s.String()
}

// SetAdjustmentType sets the AdjustmentType field's value.
func (s *ScalingPolicyForDescribeScalingPoliciesOutput) SetAdjustmentType(v string) *ScalingPolicyForDescribeScalingPoliciesOutput {
	s.AdjustmentType = &v
	return s
}

// SetAdjustmentValue sets the AdjustmentValue field's value.
func (s *ScalingPolicyForDescribeScalingPoliciesOutput) SetAdjustmentValue(v int32) *ScalingPolicyForDescribeScalingPoliciesOutput {
	s.AdjustmentValue = &v
	return s
}

// SetAlarmPolicy sets the AlarmPolicy field's value.
func (s *ScalingPolicyForDescribeScalingPoliciesOutput) SetAlarmPolicy(v *AlarmPolicyForDescribeScalingPoliciesOutput) *ScalingPolicyForDescribeScalingPoliciesOutput {
	s.AlarmPolicy = v
	return s
}

// SetCooldown sets the Cooldown field's value.
func (s *ScalingPolicyForDescribeScalingPoliciesOutput) SetCooldown(v int32) *ScalingPolicyForDescribeScalingPoliciesOutput {
	s.Cooldown = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *ScalingPolicyForDescribeScalingPoliciesOutput) SetScalingGroupId(v string) *ScalingPolicyForDescribeScalingPoliciesOutput {
	s.ScalingGroupId = &v
	return s
}

// SetScalingPolicyId sets the ScalingPolicyId field's value.
func (s *ScalingPolicyForDescribeScalingPoliciesOutput) SetScalingPolicyId(v string) *ScalingPolicyForDescribeScalingPoliciesOutput {
	s.ScalingPolicyId = &v
	return s
}

// SetScalingPolicyName sets the ScalingPolicyName field's value.
func (s *ScalingPolicyForDescribeScalingPoliciesOutput) SetScalingPolicyName(v string) *ScalingPolicyForDescribeScalingPoliciesOutput {
	s.ScalingPolicyName = &v
	return s
}

// SetScalingPolicyType sets the ScalingPolicyType field's value.
func (s *ScalingPolicyForDescribeScalingPoliciesOutput) SetScalingPolicyType(v string) *ScalingPolicyForDescribeScalingPoliciesOutput {
	s.ScalingPolicyType = &v
	return s
}

// SetScheduledPolicy sets the ScheduledPolicy field's value.
func (s *ScalingPolicyForDescribeScalingPoliciesOutput) SetScheduledPolicy(v *ScheduledPolicyForDescribeScalingPoliciesOutput) *ScalingPolicyForDescribeScalingPoliciesOutput {
	s.ScheduledPolicy = v
	return s
}

// SetStatus sets the Status field's value.
func (s *ScalingPolicyForDescribeScalingPoliciesOutput) SetStatus(v string) *ScalingPolicyForDescribeScalingPoliciesOutput {
	s.Status = &v
	return s
}

type ScheduledPolicyForDescribeScalingPoliciesOutput struct {
	_ struct{} `type:"structure"`

	LaunchTime *string `type:"string"`

	RecurrenceEndTime *string `type:"string"`

	RecurrenceStartTime *string `type:"string"`

	RecurrenceType *string `type:"string"`

	RecurrenceValue *string `type:"string"`
}

// String returns the string representation
func (s ScheduledPolicyForDescribeScalingPoliciesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ScheduledPolicyForDescribeScalingPoliciesOutput) GoString() string {
	return s.String()
}

// SetLaunchTime sets the LaunchTime field's value.
func (s *ScheduledPolicyForDescribeScalingPoliciesOutput) SetLaunchTime(v string) *ScheduledPolicyForDescribeScalingPoliciesOutput {
	s.LaunchTime = &v
	return s
}

// SetRecurrenceEndTime sets the RecurrenceEndTime field's value.
func (s *ScheduledPolicyForDescribeScalingPoliciesOutput) SetRecurrenceEndTime(v string) *ScheduledPolicyForDescribeScalingPoliciesOutput {
	s.RecurrenceEndTime = &v
	return s
}

// SetRecurrenceStartTime sets the RecurrenceStartTime field's value.
func (s *ScheduledPolicyForDescribeScalingPoliciesOutput) SetRecurrenceStartTime(v string) *ScheduledPolicyForDescribeScalingPoliciesOutput {
	s.RecurrenceStartTime = &v
	return s
}

// SetRecurrenceType sets the RecurrenceType field's value.
func (s *ScheduledPolicyForDescribeScalingPoliciesOutput) SetRecurrenceType(v string) *ScheduledPolicyForDescribeScalingPoliciesOutput {
	s.RecurrenceType = &v
	return s
}

// SetRecurrenceValue sets the RecurrenceValue field's value.
func (s *ScheduledPolicyForDescribeScalingPoliciesOutput) SetRecurrenceValue(v string) *ScheduledPolicyForDescribeScalingPoliciesOutput {
	s.RecurrenceValue = &v
	return s
}
