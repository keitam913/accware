package account

type Total struct {
	personID string
	amount   int
}

func NewTotal(personID string, amount int) Total {
	return Total{
		personID: personID,
		amount:   amount,
	}
}

func (t Total) PersonID() string {
	return t.personID
}

func (t Total) Amount() int {
	return t.amount
}
