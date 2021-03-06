// Package query provides serialization of VOLCSTACK volcstackquery requests, and responses.
package query

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"net/url"

	"github.com/volcengine/volcstack-go-sdk/private/protocol/query/queryutil"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackerr"
)

// BuildHandler is a named request handler for building volcstackquery protocol requests
var BuildHandler = request.NamedHandler{Name: "awssdk.volcstackquery.Build", Fn: Build}

// Build builds a request for an VOLCSTACK Query service.
func Build(r *request.Request) {
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.ClientInfo.APIVersion},
	}
	if err := queryutil.Parse(body, r.Params, false); err != nil {
		r.Error = volcstackerr.New(request.ErrCodeSerialization, "failed encoding Query request", err)
		return
	}

	if !r.IsPresigned() {
		r.HTTPRequest.Method = "POST"
		r.HTTPRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		r.SetBufferBody([]byte(body.Encode()))
	} else { // This is a pre-signed request
		r.HTTPRequest.Method = "GET"
		r.HTTPRequest.URL.RawQuery = body.Encode()
	}
}
