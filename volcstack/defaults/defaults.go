// Package defaults is a collection of helpers to retrieve the SDK's default
// configuration and handlers.
//
// Generally this package shouldn't be used directly, but session.Session
// instead. This package is useful when you need to reset the defaults
// of a session or service client to the SDK defaults before setting
// additional parameters.
package defaults

import (
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/volcstackerr"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"code.byted.org/iaasng/volcstack-go-sdk/volcstack"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/corehandlers"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/credentials"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/credentials/endpointcreds"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/request"
)

// A Defaults provides a collection of default values for SDK clients.
type Defaults struct {
	Config   *volcstack.Config
	Handlers request.Handlers
}

// Get returns the SDK's default values with Config and handlers pre-configured.
func Get() Defaults {
	cfg := Config()
	handlers := Handlers()
	cfg.Credentials = CredChain(cfg, handlers)

	return Defaults{
		Config:   cfg,
		Handlers: handlers,
	}
}

// Config returns the default configuration without credentials.
// To retrieve a config with credentials also included use
// `defaults.Get().Config` instead.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the configuration of an
// existing service client or session.
func Config() *volcstack.Config {
	return volcstack.NewConfig().
		WithCredentials(credentials.AnonymousCredentials).
		WithRegion(os.Getenv("VOLCSTACK_REGION")).
		WithHTTPClient(http.DefaultClient).
		WithMaxRetries(volcstack.UseServiceDefaultRetries).
		WithLogger(volcstack.NewDefaultLogger()).
		WithLogLevel(volcstack.LogOff)
}

// Handlers returns the default request handlers.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the request handlers of an
// existing service client or session.
func Handlers() request.Handlers {
	var handlers request.Handlers

	handlers.Validate.PushBackNamed(corehandlers.ValidateEndpointHandler)
	handlers.Validate.AfterEachFn = request.HandlerListStopOnError
	handlers.Build.PushBackNamed(corehandlers.SDKVersionUserAgentHandler)
	handlers.Build.PushBackNamed(corehandlers.AddHostExecEnvUserAgentHandler)
	handlers.Build.AfterEachFn = request.HandlerListStopOnError
	handlers.Sign.PushBackNamed(corehandlers.BuildContentLengthHandler)
	handlers.Send.PushBackNamed(corehandlers.ValidateReqSigHandler)
	handlers.Send.PushBackNamed(corehandlers.SendHandler)
	handlers.AfterRetry.PushBackNamed(corehandlers.AfterRetryHandler)
	handlers.ValidateResponse.PushBackNamed(corehandlers.ValidateResponseHandler)

	return handlers
}

// CredChain returns the default credential chain.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the credentials of an
// existing service client or session's Config.
func CredChain(cfg *volcstack.Config, handlers request.Handlers) *credentials.Credentials {
	return credentials.NewCredentials(&credentials.ChainProvider{
		VerboseErrors: volcstack.BoolValue(cfg.CredentialsChainVerboseErrors),
		Providers:     CredProviders(cfg, handlers),
	})
}

// CredProviders returns the slice of providers used in
// the default credential chain.
//
// For applications that need to use some other provider (for example use
// different  environment variables for legacy reasons) but still fall back
// on the default chain of providers. This allows that default chaint to be
// automatically updated
func CredProviders(cfg *volcstack.Config, handlers request.Handlers) []credentials.Provider {
	return []credentials.Provider{
		&credentials.EnvProvider{},
		&credentials.SharedCredentialsProvider{Filename: "", Profile: ""},
	}
}

const (
	httpProviderAuthorizationEnvVar = "VOLCSTACK_CONTAINER_AUTHORIZATION_TOKEN"
	httpProviderEnvVar              = "VOLCSTACK_CONTAINER_CREDENTIALS_FULL_URI"
)

var lookupHostFn = net.LookupHost

func isLoopbackHost(host string) (bool, error) {
	ip := net.ParseIP(host)
	if ip != nil {
		return ip.IsLoopback(), nil
	}

	// Host is not an ip, perform lookup
	addrs, err := lookupHostFn(host)
	if err != nil {
		return false, err
	}
	for _, addr := range addrs {
		if !net.ParseIP(addr).IsLoopback() {
			return false, nil
		}
	}

	return true, nil
}

func localHTTPCredProvider(cfg volcstack.Config, handlers request.Handlers, u string) credentials.Provider {
	var errMsg string

	parsed, err := url.Parse(u)
	if err != nil {
		errMsg = fmt.Sprintf("invalid URL, %v", err)
	} else {
		host := volcstack.URLHostname(parsed)
		if len(host) == 0 {
			errMsg = "unable to parse host from local HTTP cred provider URL"
		} else if isLoopback, loopbackErr := isLoopbackHost(host); loopbackErr != nil {
			errMsg = fmt.Sprintf("failed to resolve host %q, %v", host, loopbackErr)
		} else if !isLoopback {
			errMsg = fmt.Sprintf("invalid endpoint host, %q, only loopback hosts are allowed.", host)
		}
	}

	if len(errMsg) > 0 {
		if cfg.Logger != nil {
			cfg.Logger.Log("Ignoring, HTTP credential provider", errMsg, err)
		}
		return credentials.ErrorProvider{
			Err:          volcstackerr.New("CredentialsEndpointError", errMsg, err),
			ProviderName: endpointcreds.ProviderName,
		}
	}

	return httpCredProvider(cfg, handlers, u)
}

func httpCredProvider(cfg volcstack.Config, handlers request.Handlers, u string) credentials.Provider {
	return endpointcreds.NewProviderClient(cfg, handlers, u,
		func(p *endpointcreds.Provider) {
			p.ExpiryWindow = 5 * time.Minute
			p.AuthorizationToken = os.Getenv(httpProviderAuthorizationEnvVar)
		},
	)
}
