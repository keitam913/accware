package person

type Service interface {
	PaymentRatio(personID string) (float64, error)
	Persons() ([]Person, error)
}
