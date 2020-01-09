package account

import (
	"time"

	"github.com/keitam913/accware-api/domain/account"
)

type Respository interface {
	Month(year int, month time.Month) (Month, error)
	Add(item account.Item) error
}
