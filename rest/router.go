package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/keitam913/accware/api/oidc"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(idService *oidc.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !strings.HasPrefix(ctx.Request.URL.Path, "/v") {
			ctx.Next()
			return
		}
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
			"https://localhost:3000",
			"https://accware.keitam.com",
		},
		AllowHeaders: []string{
			"ID-Token",
		},
		AllowMethods: []string{"GET", "POST", "DELETE"},
	}))

	r.Use(AuthMiddleware(idService))

	v1r := gin.Default()
	v1 := v1r.Group("/v1")
	v1.GET("/accounts/:year/:month", monthHandler.Get)
	v1.POST("/items", itemHandler.Post)
	v1.DELETE("/items/:id", itemHandler.Delete)

	sr := gin.Default()
	sr.Static("/static", "assets/static")

	r.GET("/*resource", func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Param("resource"), "/v1") {
			v1r.HandleContext(ctx)
			return
		}
		if strings.HasPrefix(ctx.Param("resource"), "/static") {
			sr.HandleContext(ctx)
			return
		}
		ctx.File("assets/index.html")
	})

	return r
}
