package di

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Container struct{}

func (c Container) Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello")
	})
	return r
}
