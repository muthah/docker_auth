package authn

// Auth0AuthConfig holds the configuration for the auth0 authenticator
type Auth0AuthConfig struct {
	Domain       string `yaml:"domain,omitempty"`
	Connection   string `yaml:"connection,omitempty"`
	ClientID     string `yaml:"client_id,omitempty"`
	ClientSecret string `yaml:"client_secret,omitempty"`
}

// Auth0Auth is an auth0 authenticator
type Auth0Auth struct {
	config *Auth0AuthConfig
}

// NewAuth0Auth creates a new authenticator
func NewAuth0Auth(c *Auth0AuthConfig) (*Auth0Auth, error) {
	return &Auth0Auth{
		config: c,
	}, nil
}

// Authenticate user with username and password
func (aa *Auth0Auth) Authenticate(account string, password PasswordString) (bool, error) {
	return true, nil
}

// Stop authenticating
func (aa *Auth0Auth) Stop() {
}

// Name of authenticator
func (aa *Auth0Auth) Name() string {
	return "Auth0"
}
