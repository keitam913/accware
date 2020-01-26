package rest

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/keitam913/accware-api/account"
	"github.com/keitam913/accware-api/application"
	"github.com/keitam913/accware-api/person"
)

type MonthHandler struct {
	Service       *application.Service
	PersonService person.Service
}

func (mc *MonthHandler) Get(ctx *gin.Context) {
	y, err := strconv.Atoi(ctx.Param("year"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("year must be a number"))
		return
	}
	m, err := strconv.Atoi(ctx.Param("month"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("month must be a number"))
		return
	}
	monthAccount, err := mc.Service.GetMonth(y, time.Month(m))
	if err != nil {
		panic(err)
	}
	res, err := mc.makeResponse(monthAccount)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (mc *MonthHandler) makeResponse(monthAccount account.Month) (Month, error) {
	ps, err := mc.persons()
	if err != nil {
		return Month{}, err
	}
	var is []Item
	for _, i := range monthAccount.Items() {
		is = append(is, Item{
			Name:     i.Name(),
			Amount:   i.Amount(),
			PersonID: i.PersonID(),
			Date:     i.Date(),
		})
	}
	var as []Amount
	for _, a := range monthAccount.Adjustments(mc.PersonService) {
		as = append(as, Amount{PersonID: a.PersonID(), Amount: a.Amount()})
	}
	var ts []Amount
	for _, t := range monthAccount.Totals(mc.PersonService) {
		ts = append(ts, Amount{PersonID: t.PersonID(), Amount: t.Amount()})
	}
	return Month{
		Persons:     ps,
		Items:       is,
		Adjustments: as,
		Totals:      ts,
	}, nil
}

func (mc *MonthHandler) persons() (map[string]string, error) {
	m := map[string]string{}
	ps, err := mc.PersonService.Persons()
	if err != nil {
		return nil, err
	}
	for _, p := range ps {
		m[p.ID()] = p.Name()
	}
	return m, nil
}
