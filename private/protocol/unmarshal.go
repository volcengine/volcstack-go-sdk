package protocol

import (
	"io"
	"io/ioutil"

	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/request"
)

// UnmarshalDiscardBodyHandler is a named request handler to empty and close a response's volcstackbody
var UnmarshalDiscardBodyHandler = request.NamedHandler{Name: "volcstacksdk.shared.UnmarshalDiscardBody", Fn: UnmarshalDiscardBody}

// UnmarshalDiscardBody is a request handler to empty a response's volcstackbody and closing it.
func UnmarshalDiscardBody(r *request.Request) {
	if r.HTTPResponse == nil || r.HTTPResponse.Body == nil {
		return
	}

	io.Copy(ioutil.Discard, r.HTTPResponse.Body)
	r.HTTPResponse.Body.Close()
}
