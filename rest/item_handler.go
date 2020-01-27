package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/keitam913/accware-api/application"
)

type ItemHandler struct {
	Service *application.Service
}

func (ih *ItemHandler) Post(ctx *gin.Context) {
	var item Item
	if err := json.NewDecoder(ctx.Request.Body).Decode(&item); err != nil {
		panic(err)
	}

	d, err := time.ParseInLocation("2006-01-02", item.Date, time.Local)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid date string")
		return
	}

	if err := ih.Service.AddItem(item.Name, item.Amount, ctx.GetString("id"), d); err != nil {
		panic(err)
	}

	ctx.Status(http.StatusNoContent)
}
