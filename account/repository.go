package account

import (
	"time"
)

type Respository interface {
	Month(year int, month time.Month) (Month, error)
	Add(item Item) error
}
