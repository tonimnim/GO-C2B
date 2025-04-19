package mpesa

import (
	"encoding/base64"
	"net/http"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

type Authorization struct {
	AuthResponse
}

func NewAuthorization(consumerKey, consumerSecret string, env Environment) (*Authorization, error) {
	auth := &Authorization{}
	authHeader := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(consumerKey+":"+consumerSecret)),
	}
	netPackage := newRequestPackage(nil, "/oauth/v1/generate?grant_type=client_credentials", http.MethodGet, authHeader, env)
	authResponse, err := newRequest[Authorization](netPackage)
	if err != nil {
		return nil, err
	}
	auth = &authResponse.Body
	return auth, nil
}
