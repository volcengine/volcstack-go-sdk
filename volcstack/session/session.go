package session

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/client"
	"github.com/volcengine/volcstack-go-sdk/volcstack/corehandlers"
	"github.com/volcengine/volcstack-go-sdk/volcstack/credentials"
	"github.com/volcengine/volcstack-go-sdk/volcstack/defaults"
	"github.com/volcengine/volcstack-go-sdk/volcstack/endpoints"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackerr"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const (
	// ErrCodeSharedConfig represents an error that occurs in the shared
	// configuration logic
	ErrCodeSharedConfig = "SharedConfigErr"
)

// ErrSharedConfigSourceCollision will be returned if a section contains both
// source_profile and credential_source
var ErrSharedConfigSourceCollision = volcstackerr.New(ErrCodeSharedConfig, "only source profile or credential source can be specified, not both", nil)

// ErrSharedConfigECSContainerEnvVarEmpty will be returned if the environment
// variables are empty and Environment was set as the credential source
var ErrSharedConfigECSContainerEnvVarEmpty = volcstackerr.New(ErrCodeSharedConfig, "EcsContainer was specified as the credential_source, but 'volcstack_CONTAINER_CREDENTIALS_RELATIVE_URI' was not set", nil)

// ErrSharedConfigInvalidCredSource will be returned if an invalid credential source was provided
var ErrSharedConfigInvalidCredSource = volcstackerr.New(ErrCodeSharedConfig, "credential source values must be EcsContainer, Ec2InstanceMetadata, or Environment", nil)

// A Session provides a central location to create service clients from and
// store configurations and request handlers for those services.
//
// Sessions are safe to create service clients concurrently, but it is not safe
// to mutate the Session concurrently.
//
// The Session satisfies the service client's client.ConfigProvider.
type Session struct {
	Config   *volcstack.Config
	Handlers request.Handlers
}

// New creates a new instance of the handlers merging in the provided configs
// on top of the SDK's default configurations. Once the Session is created it
// can be mutated to modify the Config or Handlers. The Session is safe to be
// read concurrently, but it should not be written to concurrently.
//
// If the volcstack_SDK_LOAD_CONFIG environment is set to a truthy value, the New
// method could now encounter an error when loading the configuration. When
// The environment variable is set, and an error occurs, New will return a
// session that will fail all requests reporting the error that occurred while
// loading the session. Use NewSession to get the error when creating the
// session.
//
// If the volcstack_SDK_LOAD_CONFIG environment variable is set to a truthy value
// the shared config file (~/.volcstack/config) will also be loaded, in addition to
// the shared credentials file (~/.volcstack/credentials). Values set in both the
// shared config, and shared credentials will be taken from the shared
// credentials file.
//
// Deprecated: Use NewSession functions to create sessions instead. NewSession
// has the same functionality as New except an error can be returned when the
// func is called instead of waiting to receive an error until a request is made.
func New(cfgs ...*volcstack.Config) *Session {
	// load initial config from environment
	envCfg := loadEnvConfig()

	if envCfg.EnableSharedConfig {
		var cfg volcstack.Config
		cfg.MergeIn(cfgs...)
		s, err := NewSessionWithOptions(Options{
			Config:            cfg,
			SharedConfigState: SharedConfigEnable,
		})
		if err != nil {
			// Old session.New expected all errors to be discovered when
			// a request is made, and would report the errors then. This
			// needs to be replicated if an error occurs while creating
			// the session.
			msg := "failed to create session with volcstack_SDK_LOAD_CONFIG enabled. " +
				"Use session.NewSession to handle errors occurring during session creation."

			// Session creation failed, need to report the error and prevent
			// any requests from succeeding.
			s = &Session{Config: defaults.Config()}
			s.Config.MergeIn(cfgs...)
			s.Config.Logger.Log("ERROR:", msg, "Error:", err)
			s.Handlers.Validate.PushBack(func(r *request.Request) {
				r.Error = err
			})
		}

		return s
	}

	s := deprecatedNewSession(cfgs...)

	return s
}

// NewSession returns a new Session created from SDK defaults, config files,
// environment, and user provided config files. Once the Session is created
// it can be mutated to modify the Config or Handlers. The Session is safe to
// be read concurrently, but it should not be written to concurrently.
//
// If the volcstack_SDK_LOAD_CONFIG environment variable is set to a truthy value
// the shared config file (~/.volcstack/config) will also be loaded in addition to
// the shared credentials file (~/.volcstack/credentials). Values set in both the
// shared config, and shared credentials will be taken from the shared
// credentials file. Enabling the Shared Config will also allow the Session
// to be built with retrieving credentials with AssumeRole set in the config.
//
// See the NewSessionWithOptions func for information on how to override or
// control through code how the Session will be created, such as specifying the
// config profile, and controlling if shared config is enabled or not.
func NewSession(cfgs ...*volcstack.Config) (*Session, error) {
	opts := Options{}
	opts.Config.MergeIn(cfgs...)

	if opts.Config.Endpoint == nil {
		opts.Config.Endpoint = volcstack.String(volcstackutil.NewEndpoint().GetEndpoint())
	}

	return NewSessionWithOptions(opts)
}

// SharedConfigState provides the ability to optionally override the state
// of the session's creation based on the shared config being enabled or
// disabled.
type SharedConfigState int

const (
	// SharedConfigStateFromEnv does not override any state of the
	// volcstack_SDK_LOAD_CONFIG env var. It is the default value of the
	// SharedConfigState type.
	SharedConfigStateFromEnv SharedConfigState = iota

	// SharedConfigDisable overrides the volcstack_SDK_LOAD_CONFIG env var value
	// and disables the shared config functionality.
	SharedConfigDisable

	// SharedConfigEnable overrides the volcstack_SDK_LOAD_CONFIG env var value
	// and enables the shared config functionality.
	SharedConfigEnable
)

// Options provides the means to control how a Session is created and what
// configuration values will be loaded.
//
type Options struct {
	// Provides config values for the SDK to use when creating service clients
	// and making API requests to services. Any value set in with this field
	// will override the associated value provided by the SDK defaults,
	// environment or config files where relevant.
	//
	// If not set, configuration values from from SDK defaults, environment,
	// config will be used.
	Config volcstack.Config

	// Overrides the config profile the Session should be created from. If not
	// set the value of the environment variable will be loaded (volcstack_PROFILE,
	// or volcstack_DEFAULT_PROFILE if the Shared Config is enabled).
	//
	// If not set and environment variables are not set the "default"
	// (DefaultSharedConfigProfile) will be used as the profile to load the
	// session config from.
	Profile string

	// Instructs how the Session will be created based on the volcstack_SDK_LOAD_CONFIG
	// environment variable. By default a Session will be created using the
	// value provided by the volcstack_SDK_LOAD_CONFIG environment variable.
	//
	// Setting this value to SharedConfigEnable or SharedConfigDisable
	// will allow you to override the volcstack_SDK_LOAD_CONFIG environment variable
	// and enable or disable the shared config functionality.
	SharedConfigState SharedConfigState

	// Ordered list of files the session will load configuration from.
	// It will override environment variable volcstack_SHARED_CREDENTIALS_FILE, volcstack_CONFIG_FILE.
	SharedConfigFiles []string

	// When the SDK's shared config is configured to assume a role with MFA
	// this option is required in order to provide the mechanism that will
	// retrieve the MFA token. There is no default value for this field. If
	// it is not set an error will be returned when creating the session.
	//
	// This token provider will be called when ever the assumed role's
	// credentials need to be refreshed. Within the context of service clients
	// all sharing the same session the SDK will ensure calls to the token
	// provider are atomic. When sharing a token provider across multiple
	// sessions additional synchronization logic is needed to ensure the
	// token providers do not introduce race conditions. It is recommend to
	// share the session where possible.
	//
	// stscreds.StdinTokenProvider is a basic implementation that will prompt
	// from stdin for the MFA token code.
	//
	// This field is only used if the shared configuration is enabled, and
	// the config enables assume role wit MFA via the mfa_serial field.
	AssumeRoleTokenProvider func() (string, error)

	// When the SDK's shared config is configured to assume a role this option
	// may be provided to set the expiry duration of the STS credentials.
	// Defaults to 15 minutes if not set as documented in the
	// stscreds.AssumeRoleProvider.
	AssumeRoleDuration time.Duration

	// Reader for a custom Credentials Authority (CA) bundle in PEM format that
	// the SDK will use instead of the default system's root CA bundle. Use this
	// only if you want to replace the CA bundle the SDK uses for TLS requests.
	//
	// Enabling this option will attempt to merge the Transport into the SDK's HTTP
	// client. If the client's Transport is not a http.Transport an error will be
	// returned. If the Transport's TLS config is set this option will cause the SDK
	// to overwrite the Transport's TLS config's  RootCAs value. If the CA
	// bundle reader contains multiple certificates all of them will be loaded.
	//
	// The Session option CustomCABundle is also available when creating sessions
	// to also enable this feature. CustomCABundle session option field has priority
	// over the volcstack_CA_BUNDLE environment variable, and will be used if both are set.
	CustomCABundle io.Reader

	// The handlers that the session and all API clients will be created with.
	// This must be a complete set of handlers. Use the defaults.Handlers()
	// function to initialize this value before changing the handlers to be
	// used by the SDK.
	Handlers request.Handlers
}

// NewSessionWithOptions returns a new Session created from SDK defaults, config files,
// environment, and user provided config files. This func uses the Options
// values to configure how the Session is created.
//
// If the volcstack_SDK_LOAD_CONFIG environment variable is set to a truthy value
// the shared config file (~/.volcstack/config) will also be loaded in addition to
// the shared credentials file (~/.volcstack/credentials). Values set in both the
// shared config, and shared credentials will be taken from the shared
// credentials file. Enabling the Shared Config will also allow the Session
// to be built with retrieving credentials with AssumeRole set in the config.
//
//     // Equivalent to session.New
//     sess := session.Must(session.NewSessionWithOptions(session.Options{}))
//
//     // Specify profile to load for the session's config
//     sess := session.Must(session.NewSessionWithOptions(session.Options{
//          Profile: "profile_name",
//     }))
//
//     // Specify profile for config and region for requests
//     sess := session.Must(session.NewSessionWithOptions(session.Options{
//          Config: volcstack.Config{Region: volcstack.String("us-east-1")},
//          Profile: "profile_name",
//     }))
//
//     // Force enable Shared Config support
//     sess := session.Must(session.NewSessionWithOptions(session.Options{
//         SharedConfigState: session.SharedConfigEnable,
//     }))
func NewSessionWithOptions(opts Options) (*Session, error) {
	var envCfg envConfig
	if opts.SharedConfigState == SharedConfigEnable {
		envCfg = loadSharedEnvConfig()
	} else {
		envCfg = loadEnvConfig()
	}

	if len(opts.Profile) != 0 {
		envCfg.Profile = opts.Profile
	}

	switch opts.SharedConfigState {
	case SharedConfigDisable:
		envCfg.EnableSharedConfig = false
	case SharedConfigEnable:
		envCfg.EnableSharedConfig = true
	}

	// Only use volcstack_CA_BUNDLE if session option is not provided.
	if len(envCfg.CustomCABundle) != 0 && opts.CustomCABundle == nil {
		f, err := os.Open(envCfg.CustomCABundle)
		if err != nil {
			return nil, volcstackerr.New("LoadCustomCABundleError",
				"failed to open custom CA bundle PEM file", err)
		}
		defer f.Close()
		opts.CustomCABundle = f
	}

	return newSession(opts, envCfg, &opts.Config)
}

// Must is a helper function to ensure the Session is valid and there was no
// error when calling a NewSession function.
//
// This helper is intended to be used in variable initialization to load the
// Session and configuration at startup. Such as:
//
//     var sess = session.Must(session.NewSession())
func Must(sess *Session, err error) *Session {
	if err != nil {
		panic(err)
	}

	return sess
}

func deprecatedNewSession(cfgs ...*volcstack.Config) *Session {
	cfg := defaults.Config()
	handlers := defaults.Handlers()

	// Apply the passed in configs so the configuration can be applied to the
	// default credential chain
	cfg.MergeIn(cfgs...)
	cfg.Credentials = defaults.CredChain(cfg, handlers)

	// Reapply any passed in configs to override credentials if set
	cfg.MergeIn(cfgs...)

	s := &Session{
		Config:   cfg,
		Handlers: handlers,
	}

	initHandlers(s)
	return s
}

func newSession(opts Options, envCfg envConfig, cfgs ...*volcstack.Config) (*Session, error) {
	cfg := defaults.Config()

	handlers := opts.Handlers
	if handlers.IsEmpty() {
		handlers = defaults.Handlers()
	}

	// Get a merged version of the user provided config to determine if
	// credentials were.
	userCfg := &volcstack.Config{}
	userCfg.MergeIn(cfgs...)
	cfg.MergeIn(userCfg)

	// Ordered config files will be loaded in with later files overwriting
	// previous config file values.
	var cfgFiles []string
	if opts.SharedConfigFiles != nil {
		cfgFiles = opts.SharedConfigFiles
	} else {
		cfgFiles = []string{envCfg.SharedConfigFile, envCfg.SharedCredentialsFile}
		if !envCfg.EnableSharedConfig {
			// The shared config file (~/.volcstack/config) is only loaded if instructed
			// to load via the envConfig.EnableSharedConfig (volcstack_SDK_LOAD_CONFIG).
			cfgFiles = cfgFiles[1:]
		}
	}

	// Load additional config from file(s)
	sharedCfg, err := loadSharedConfig(envCfg.Profile, cfgFiles, envCfg.EnableSharedConfig)
	if err != nil {
		if len(envCfg.Profile) == 0 && !envCfg.EnableSharedConfig && (envCfg.Creds.HasKeys() || userCfg.Credentials != nil) {
			// Special case where the user has not explicitly specified an volcstack_PROFILE,
			// or session.Options.profile, shared config is not enabled, and the
			// environment has credentials, allow the shared config file to fail to
			// load since the user has already provided credentials, and nothing else
			// is required to be read file. Github(volcstack/volcstack-go-sdk#2455)
		} else if _, ok := err.(SharedConfigProfileNotExistsError); !ok {
			return nil, err
		}
	}

	if err := mergeConfigSrcs(cfg, userCfg, envCfg, sharedCfg, handlers, opts); err != nil {
		return nil, err
	}

	s := &Session{
		Config:   cfg,
		Handlers: handlers,
	}

	initHandlers(s)

	// Setup HTTP client with custom cert bundle if enabled
	if opts.CustomCABundle != nil {
		if err := loadCustomCABundle(s, opts.CustomCABundle); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func getCABundleTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func loadCustomCABundle(s *Session, bundle io.Reader) error {
	var t *http.Transport
	switch v := s.Config.HTTPClient.Transport.(type) {
	case *http.Transport:
		t = v
	default:
		if s.Config.HTTPClient.Transport != nil {
			return volcstackerr.New("LoadCustomCABundleError",
				"unable to load custom CA bundle, HTTPClient's transport unsupported type", nil)
		}
	}
	if t == nil {
		// Nil transport implies `http.DefaultTransport` should be used. Since
		// the SDK cannot modify, nor copy the `DefaultTransport` specifying
		// the values the next closest behavior.
		t = getCABundleTransport()
	}

	p, err := loadCertPool(bundle)
	if err != nil {
		return err
	}
	if t.TLSClientConfig == nil {
		t.TLSClientConfig = &tls.Config{}
	}
	t.TLSClientConfig.RootCAs = p

	s.Config.HTTPClient.Transport = t

	return nil
}

func loadCertPool(r io.Reader) (*x509.CertPool, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, volcstackerr.New("LoadCustomCABundleError",
			"failed to read custom CA bundle PEM file", err)
	}

	p := x509.NewCertPool()
	if !p.AppendCertsFromPEM(b) {
		return nil, volcstackerr.New("LoadCustomCABundleError",
			"failed to load custom CA bundle PEM file", err)
	}

	return p, nil
}

func mergeConfigSrcs(cfg, userCfg *volcstack.Config,
	envCfg envConfig, sharedCfg sharedConfig,
	handlers request.Handlers,
	sessOpts Options,
) error {

	// Region if not already set by user
	if len(volcstack.StringValue(cfg.Region)) == 0 {
		if len(envCfg.Region) > 0 {
			cfg.WithRegion(envCfg.Region)
		} else if envCfg.EnableSharedConfig && len(sharedCfg.Region) > 0 {
			cfg.WithRegion(sharedCfg.Region)
		}
	}

	if cfg.EnableEndpointDiscovery == nil {
		if envCfg.EnableEndpointDiscovery != nil {
			cfg.WithEndpointDiscovery(*envCfg.EnableEndpointDiscovery)
		} else if envCfg.EnableSharedConfig && sharedCfg.EnableEndpointDiscovery != nil {
			cfg.WithEndpointDiscovery(*sharedCfg.EnableEndpointDiscovery)
		}
	}

	// Configure credentials if not already set by the user when creating the
	// Session.
	if cfg.Credentials == credentials.AnonymousCredentials && userCfg.Credentials == nil {
		creds, err := resolveCredentials(cfg, envCfg, sharedCfg, handlers, sessOpts)
		if err != nil {
			return err
		}
		cfg.Credentials = creds
	}

	return nil
}

func initHandlers(s *Session) {
	// Add the Validate parameter handler if it is not disabled.
	s.Handlers.Validate.Remove(corehandlers.ValidateParametersHandler)
	// Temporary notes by xuyaming@bytedance.com because some validate field is relation.
	//if !volcstack.BoolValue(s.Config.DisableParamValidation) {
	//	s.Handlers.Validate.PushBackNamed(corehandlers.ValidateParametersHandler)
	//}
}

// Copy creates and returns a copy of the current Session, copying the config
// and handlers. If any additional configs are provided they will be merged
// on top of the Session's copied config.
//
//     // Create a copy of the current Session, configured for the us-west-2 region.
//     sess.Copy(&volcstack.Config{Region: volcstack.String("us-west-2")})
func (s *Session) Copy(cfgs ...*volcstack.Config) *Session {
	newSession := &Session{
		Config:   s.Config.Copy(cfgs...),
		Handlers: s.Handlers.Copy(),
	}

	initHandlers(newSession)

	return newSession
}

// ClientConfig satisfies the client.ConfigProvider interface and is used to
// configure the service client instances. Passing the Session to the service
// client's constructor (New) will use this method to configure the client.
func (s *Session) ClientConfig(serviceName string, cfgs ...*volcstack.Config) client.Config {
	// Backwards compatibility, the error will be eaten if user calls ClientConfig
	// directly. All SDK services will use ClientconfigWithError.
	cfg, _ := s.clientConfigWithErr(serviceName, cfgs...)

	return cfg
}

func (s *Session) clientConfigWithErr(serviceName string, cfgs ...*volcstack.Config) (client.Config, error) {
	s = s.Copy(cfgs...)

	var resolved endpoints.ResolvedEndpoint
	var err error

	region := volcstack.StringValue(s.Config.Region)

	if endpoint := volcstack.StringValue(s.Config.Endpoint); len(endpoint) != 0 {
		resolved.URL = endpoints.AddScheme(endpoint, volcstack.BoolValue(s.Config.DisableSSL))
		resolved.SigningRegion = region
	}

	return client.Config{
		Config:             s.Config,
		Handlers:           s.Handlers,
		Endpoint:           resolved.URL,
		SigningRegion:      resolved.SigningRegion,
		SigningNameDerived: resolved.SigningNameDerived,
		SigningName:        resolved.SigningName,
	}, err
}

// ClientConfigNoResolveEndpoint is the same as ClientConfig with the exception
// that the EndpointResolver will not be used to resolve the endpoint. The only
// endpoint set must come from the volcstack.Config.Endpoint field.
func (s *Session) ClientConfigNoResolveEndpoint(cfgs ...*volcstack.Config) client.Config {
	s = s.Copy(cfgs...)

	var resolved endpoints.ResolvedEndpoint

	region := volcstack.StringValue(s.Config.Region)

	if ep := volcstack.StringValue(s.Config.Endpoint); len(ep) > 0 {
		resolved.URL = endpoints.AddScheme(ep, volcstack.BoolValue(s.Config.DisableSSL))
		resolved.SigningRegion = region
	}

	return client.Config{
		Config:             s.Config,
		Handlers:           s.Handlers,
		Endpoint:           resolved.URL,
		SigningRegion:      resolved.SigningRegion,
		SigningNameDerived: resolved.SigningNameDerived,
		SigningName:        resolved.SigningName,
	}
}
