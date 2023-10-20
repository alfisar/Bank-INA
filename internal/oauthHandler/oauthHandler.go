package oauthhandler

import (
	"Bank-INA/domain"
	"Bank-INA/internal/errorhandler"
	"net/http"
)

type OAuthHandler struct {
	ClientID     string
	ClientSecret string
	TokenUrl     string
}

func GetOAuthHandler() *OAuthHandler {
	return &OAuthHandler{}
}

func (obj *OAuthHandler) ValidateToken(token string) (err domain.ErrorData) {
	// config := &oauth2.Config{
	// 	ClientID:     obj.ClientID,
	// 	ClientSecret: obj.ClientSecret,
	// 	Endpoint: oauth2.Endpoint{
	// 		TokenURL: obj.TokenUrl,
	// 	},
	// }

	// tokenOauth := &oauth2.Token{AccessToken: token}
	// httpClient := config.Client(oauth2.NoContext, tokenOauth)

	resp, errData := http.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=" + token)
	if errData != nil {
		err = errorhandler.Failed(domain.FailedTokenCode, domain.InvalidToken, errData)
		return
	}

	if resp.StatusCode != 200 {
		err = errorhandler.Failed(domain.FailedTokenCode, domain.InvalidToken, errData)
		return
	}
	return
	// defer resp.Body.Close()

}
