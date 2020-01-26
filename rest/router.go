package rest

import (
	"errors"
	"log"
	"net/http"

	"github.com/keitam913/accware-api/oidc"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(idService *oidc.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := idService.Authenticate(ctx.Request.Header.Get("ID-TOKEN"))
		if errors.Is(err, oidc.InvalidToken) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		log.Printf("id: %s", id)
		ctx.Set("id", id)
		ctx.Next()
	}
}

func NewRouter(idService *oidc.Service, monthHandler *MonthHandler, itemHandler *ItemHandler) *gin.Engine {
	r := gin.Default()
	r.Use(AuthMiddleware(idService))
	r.GET("/v1/accounts/:year/:month", monthHandler.Get)
	r.POST("/v1/items", itemHandler.Post)
	return r
}
