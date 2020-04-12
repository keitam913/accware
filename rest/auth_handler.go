package rest

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	OAuthClientID string
}

func (ah *AuthHandler) Proxy(ctx *gin.Context) {
	u, err := url.Parse("https://accounts.google.com/o/oauth2/v2")
	if err != nil {
		panic(err)
	}

	ctx.Request.Host = "accounts.google.com"

	q := url.Values{}
	q.Add("client_id", ah.OAuthClientID)
	u.RawQuery = q.Encode()

	httputil.NewSingleHostReverseProxy(u).ServeHTTP(ctx.Writer, ctx.Request)
}
