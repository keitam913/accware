package account

type Total struct {
	person string
	amount int
}

func NewTotal(person string, amount int) Total {
	return Total{
		person: person,
		amount: amount,
	}
}

func (t Total) Person() string {
	return t.person
}

func (t Total) Amount() int {
	return t.amount
}
