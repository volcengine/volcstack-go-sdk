// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDeleteCertificateCommon = "DeleteCertificate"

// DeleteCertificateCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DeleteCertificateCommon operation. The "output" return
// value will be populated with the DeleteCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCertificateCommon Send returns without error.
//
// See DeleteCertificateCommon for more information on using the DeleteCertificateCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteCertificateCommonRequest method.
//    req, resp := client.DeleteCertificateCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DeleteCertificateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteCertificateCommon,
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

// DeleteCertificateCommon API operation for CLB.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CLB's
// API operation DeleteCertificateCommon for usage and error information.
func (c *CLB) DeleteCertificateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteCertificateCommonRequest(input)
	return out, req.Send()
}

// DeleteCertificateCommonWithContext is the same as DeleteCertificateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCertificateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DeleteCertificateCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteCertificateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteCertificate = "DeleteCertificate"

// DeleteCertificateRequest generates a "volcstack/request.Request" representing the
// client's request for the DeleteCertificate operation. The "output" return
// value will be populated with the DeleteCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCertificateCommon Send returns without error.
//
// See DeleteCertificate for more information on using the DeleteCertificate
// API call, and error handling.
//
//    // Example sending a request using the DeleteCertificateRequest method.
//    req, resp := client.DeleteCertificateRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DeleteCertificateRequest(input *DeleteCertificateInput) (req *request.Request, output *DeleteCertificateOutput) {
	op := &request.Operation{
		Name:       opDeleteCertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCertificateInput{}
	}

	output = &DeleteCertificateOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteCertificate API operation for CLB.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CLB's
// API operation DeleteCertificate for usage and error information.
func (c *CLB) DeleteCertificate(input *DeleteCertificateInput) (*DeleteCertificateOutput, error) {
	req, out := c.DeleteCertificateRequest(input)
	return out, req.Send()
}

// DeleteCertificateWithContext is the same as DeleteCertificate with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCertificate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DeleteCertificateWithContext(ctx volcstack.Context, input *DeleteCertificateInput, opts ...request.Option) (*DeleteCertificateOutput, error) {
	req, out := c.DeleteCertificateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteCertificateInput struct {
	_ struct{} `type:"structure"`

	// CertificateId is a required field
	CertificateId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCertificateInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCertificateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCertificateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteCertificateInput"}
	if s.CertificateId == nil {
		invalidParams.Add(request.NewErrParamRequired("CertificateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCertificateId sets the CertificateId field's value.
func (s *DeleteCertificateInput) SetCertificateId(v string) *DeleteCertificateInput {
	s.CertificateId = &v
	return s
}

type DeleteCertificateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteCertificateOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCertificateOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteCertificateOutput) SetRequestId(v string) *DeleteCertificateOutput {
	s.RequestId = &v
	return s
}
