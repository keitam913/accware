package account

import "time"

type Item struct {
	name   string
	amount int
	person string
	date   time.Time
}

func NewItem(name string, amount int, person string, date time.Time) Item {
	return Item{
		name:   name,
		amount: amount,
		person: person,
		date:   date,
	}
}

func (a Item) Name() string {
	return a.name
}

func (a Item) Amount() int {
	return a.amount
}

func (a Item) Person() string {
	return a.person
}

func (a Item) Date() time.Time {
	return a.date
}
