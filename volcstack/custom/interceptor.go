package custom

import (
	"context"
	"net/http"
	"net/url"

	"github.com/volcengine/volcstack-go-sdk/volcstack/client/metadata"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
)

type SdkInterceptor struct {
	Before BeforeCall
	After  AfterCall
}

type RequestInfo struct {
	Context    context.Context
	Request    *http.Request
	Response   *http.Response
	Name       string
	Method     string
	ClientInfo metadata.ClientInfo
	URI        string
	Header     http.Header
	URL        *url.URL
	Input      interface{}
	Output     interface{}
	Metadata   response.ResponseMetadata
	Error      error
}

type BeforeCall func(RequestInfo) interface{}

type AfterCall func(RequestInfo, interface{})
