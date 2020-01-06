package account

import "time"

type Item struct {
	name   string
	amount int
	person string
	date   time.Time
}

func New(name string, amount int, person string) Item {
	return Item{
		name:   name,
		amount: amount,
		person: person,
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
