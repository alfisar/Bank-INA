package oauthhandler

import (
	"Bank-INA/domain"
	"Bank-INA/internal/errorhandler"
	"net/http"
)

type OAuthHandler struct {
}

func GetOAuthHandler() *OAuthHandler {
	return &OAuthHandler{}
}

func (obj *OAuthHandler) ValidateToken(token string) (err domain.ErrorData) {
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

}
