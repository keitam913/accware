package config

import "github.com/keitam913/accware-api/person"

type PersonService struct {
	Config *Config
}

func (ps *PersonService) PaymentRatio(personID string) (float64, error) {
	for _, person := range ps.Config.Persons {
		if person.ID == personID {
			return person.Ratio, nil
		}
	}
	return 0, nil
}

func (ps *PersonService) Persons() ([]person.Person, error) {
	var p []person.Person
	for _, cp := range ps.Config.Persons {
		p = append(p, person.New(cp.ID, cp.Name))
	}
	return p, nil
}
