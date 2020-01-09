package account

import (
	"time"

	"github.com/keitam913/accware-api/domain/account"
)

type Service struct {
	Repository account.Respository
}

func (s *Service) GetMonth(year int, month time.Month) (account.Month, error) {
	return s.Repository.Month(year, month)
}

func (s *Service) AddItem(name string, amount int, person string, date time.Time) error {
	item := account.NewItem(name, amount, person, date)
	return s.Repository.Add(item)
}
