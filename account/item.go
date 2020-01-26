package account

import "time"

type Item struct {
	name     string
	amount   int
	personID string
	date     time.Time
}

func NewItem(name string, amount int, personID string, date time.Time) (Item, error) {
	return Item{
		name:     name,
		amount:   amount,
		personID: personID,
		date:     date,
	}, nil
}

func (a Item) Name() string {
	return a.name
}

func (a Item) Amount() int {
	return a.amount
}

func (a Item) PersonID() string {
	return a.personID
}

func (a Item) Date() time.Time {
	return a.date
}
