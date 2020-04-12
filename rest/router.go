package rest

import (
	"errors"
	"net/http"
	"strings"

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
	}
}

func NewRouter(idService *oidc.Service, monthHandler *MonthHandler, itemHandler *ItemHandler) *gin.Engine {
	v1r := gin.New()
	v1r.Use(AuthMiddleware(idService))
	v1 := v1r.Group("/v1")
	v1.GET("/accounts/:year/:month", monthHandler.Get)
	v1.POST("/items", itemHandler.Post)
	v1.DELETE("/items/:id", itemHandler.Delete)

	sr := gin.New()
	sr.Static("/static", "/usr/share/accware/assets/static")

	tr := gin.New()
	tr.Any("/*resource", func(ctx *gin.Context) {
		ctx.File("/usr/share/accware/assets/index.html")
	})

	r := gin.Default()

	r.Any("/*resource", func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Param("resource"), "/v1") {
			v1r.ServeHTTP(ctx.Writer, ctx.Request)
			return
		}
		if strings.HasPrefix(ctx.Param("resource"), "/static") {
			sr.ServeHTTP(ctx.Writer, ctx.Request)
			return
		}
		tr.ServeHTTP(ctx.Writer, ctx.Request)
		return
	})

	return r
}
