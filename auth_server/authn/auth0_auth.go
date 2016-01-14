package authn

type Auth0AuthConfig struct {
	Domain       string `yaml:"domain,omitempty"`
	Connection   string `yaml:"connection,omitempty"`
	ClientID     string `yaml:"client_id,omitempty"`
	ClientSecret string `yaml:"client_secret,omitempty"`
}

type Auth0Auth struct {
	config *Auth0AuthConfig
}

func NewAuth0Auth(c *Auth0AuthConfig) (*Auth0Auth, error) {
	return &Auth0Auth{
		config: c,
	}, nil
}

func (aa *Auth0Auth) Authenticate(account string, password PasswordString) (bool, error) {
	return true, nil
}
