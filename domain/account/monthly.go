package account

import (
	"math"

	"github.com/keitam913/accware-api/domain/person"
)

type Monthly struct {
	items []Item
}

func (m Monthly) Items() []Item {
	is := make([]Item, len(m.items))
	is = append(is, m.items...)
	return is
}

func (m Monthly) Adjustments(personService *person.Service) []Adjustment {
	totals := map[string]int{}
	all := 0
	for _, item := range m.items {
		totals[item.Person()] += item.Amount()
		all += item.Amount()
	}
	adjustments := make([]Adjustment, len(m.items))
	for person, total := range totals {
		amount := int(math.Round(float64(all)*personService.PaymentRatio()[person])) - total
		adjustments = append(adjustments, NewAdjustment(person, amount))
	}
}

func (m Monthly) Totals(personService *person.Service) []Total {
	totalMap := map[string]int{}
	for _, item := range m.items {
		totalMap[item.Person()] += item.Amount()
	}
	for _, adjustment := range m.Adjustments(personService) {
		totalMap[adjustment.Person()] += adjustment.Amount()
	}
	var totals []Total
	for person, amount := range totalMap {
		totals = append(totals, NewTotal(person, amount))
	}
	return totals
}
