package rest

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/keitam913/accware/api/oidc"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(idService *oidc.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := idService.Authenticate(ctx.Request.Header.Get("ID-TOKEN"))
		if err != nil {
			if errors.Is(err, oidc.InvalidToken) {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.Set("id", id)
		ctx.Next()
	}
}

func NewRouter(idService *oidc.Service, monthHandler *MonthHandler, itemHandler *ItemHandler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://localhost",
			"https://accware.keitam.com",
		},
		AllowHeaders: []string{
			"ID-Token",
		},
		AllowMethods: []string{"GET", "POST", "DELETE"},
	}))

	r.Use(AuthMiddleware(idService))

	r.GET("/v1/accounts/:year/:month", monthHandler.Get)
	r.POST("/v1/items", itemHandler.Post)
	r.DELETE("/v1/items/:id", itemHandler.Delete)

	return r
}
