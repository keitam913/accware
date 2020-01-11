package rest

import (
	"errors"
	"net/http"

	"github.com/keitam913/accware-api/infrastructure/oidc"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(idService *oidc.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := idService.Authenticate(ctx.Request.Header.Get("ID-TOKEN"))
		if errors.Is(err, oidc.InvalidToken) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("id", id)
		ctx.Next()
	}
}

func NewRouter(idService *oidc.Service) *gin.Engine {
	r := gin.Default()
	r.Use(AuthMiddleware(idService))
	return r
}
