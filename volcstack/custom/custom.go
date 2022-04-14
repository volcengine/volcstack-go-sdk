package custom

import (
	"context"
	"net/http"

	"github.com/volcengine/volcstack-go-sdk/volcstack/credentials"
)

type ExtendHttpRequest func(ctx context.Context, request *http.Request)

type ExtraHttpParameters func(ctx context.Context) map[string]string

type LogAccount func(ctx context.Context) *string

type DynamicCredentials func(ctx context.Context) (*credentials.Credentials, *string)

type RequestInterceptor func(interceptor SdkInterceptor) []interface{}

type ResponseInterceptor func(interceptor SdkInterceptor)
