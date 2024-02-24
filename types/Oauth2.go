package types

type Oauth2 struct {
	ClientID    string `json:"client_id"`
	RedirectURI string `json:"redirect_uri"`
	Login       string `json:"login"`
	Scope       string `json:"scope"`
	State       string `json:"state"`
	AllowSignup string `json:"allow_signup"`
}

type GHAccessResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}
type AuthToken struct {
	Access_token string
}
