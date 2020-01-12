package account

type Adjustment struct {
	personID string
	amount   int
}

func NewAdjustment(personID string, amount int) Adjustment {
	return Adjustment{
		personID: personID,
		amount:   amount,
	}
}

func (t Adjustment) PersonID() string {
	return t.personID
}

func (t Adjustment) Amount() int {
	return t.amount
}
