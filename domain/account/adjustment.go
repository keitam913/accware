package account

type Adjustment struct {
	person string
	amount int
}

func NewAdjustment(person string, amount int) Adjustment {
	return Adjustment{
		person: person,
		amount: amount,
	}
}

func (t Adjustment) Person() string {
	return t.person
}

func (t Adjustment) Amount() int {
	return t.amount
}
