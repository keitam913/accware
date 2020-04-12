package account

import (
	"github.com/google/uuid"
	"time"
)

type Item struct {
	id       string
	name     string
	amount   int
	personID string
	date     time.Time
}

func NewItem(name string, amount int, personID string, date time.Time) (Item, error) {
	return NewItemWithID(uuid.Must(uuid.NewRandom()).String(), name, amount, personID, date)
}

func NewItemWithID(id, name string, amount int, personID string, date time.Time) (Item, error) {
	return Item{
		id:       id,
		name:     name,
		amount:   amount,
		personID: personID,
		date:     date,
	}, nil
}

func (a Item) ID() string {
	return a.id
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
