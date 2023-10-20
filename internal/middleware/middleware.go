package middleware

import (
	"Bank-INA/domain"
	oauthhandler "Bank-INA/internal/oauthHandler"
	"Bank-INA/internal/response"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticateMiddleware struct {
	oauth *oauthhandler.OAuthHandler
}

func NewAuthenticateMiddleware(oauth *oauthhandler.OAuthHandler) *AuthenticateMiddleware {
	return &AuthenticateMiddleware{
		oauth: oauth,
	}
}

func (obj *AuthenticateMiddleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, errData := getTokenRequest(ctx.Request)
		if errData.ResponseCode != "" {
			err := domain.ErrorData{
				ResponseCode: domain.FailedGetHeaderTokenCode,
				Message:      domain.FailedGetHeaderToken,
			}
			response.Unauth(ctx, err)
			ctx.Abort()
			return
		}

		errData = obj.oauth.ValidateToken(tokenString)
		if errData.ResponseCode != "" {
			response.Unauth(ctx, errData)
			ctx.Abort()
			return
		}
		newCtx := context.WithValue(ctx.Request.Context(), "", "")
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
		// next.ServeHTTP(ctx, ctx.Request.WithContext(newCtx))
	}
}

func getTokenRequest(r *http.Request) (tokenString string, err domain.ErrorData) {
	const bearerSchema = "Bearer "
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) < len(bearerSchema) {
		err = domain.ErrorData{
			ResponseCode: domain.FailedMiddleAuthTokenCode,
			Message:      domain.FailedMiddleAuthToken,
		}
		return
	}

	tokenString = authHeader[len(bearerSchema):]
	return
}
