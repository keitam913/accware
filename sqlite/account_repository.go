package sqlite

import (
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/keitam913/accware-api/account"
)

type AccountRepository struct {
	DB *sql.DB
}

func (ar *AccountRepository) Month(year int, month time.Month) (account.Month, error) {
	var (
		nextYear  int
		nextMonth time.Month
	)
	if month == time.December {
		nextYear = year + 1
		nextMonth = time.January
	} else {
		nextYear = year
		nextMonth = month + 1
	}
	rs, err := sq.
		Select("name", "person_id", "amount", "date").
		From("item").
		Where(sq.And{sq.GtOrEq{"date": fmt.Sprintf("%04d-%02d-01", year, month)}, sq.Lt{"date": fmt.Sprintf("%04d-%02d-01", nextYear, nextMonth)}}).
		RunWith(ar.DB).
		Query()
	if err != nil {
		return account.Month{}, err
	}
	var items []account.Item
	for rs.Next() {
		var (
			name, personID, dateStr string
			amount                  int
		)
		if err := rs.Scan(&name, &personID, &amount, &dateStr); err != nil {
			panic(err)
		}
		date, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, time.Local)
		if err != nil {
			panic(err)
		}
		item, err := account.NewItem(name, amount, personID, date)
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	return account.NewMonth(items), nil
}

func (ar *AccountRepository) Add(item account.Item) error {
	d := item.Date().Format("2006-01-02 15:04:05")
	if _, err := ar.DB.Exec("insert into item (name, person_id, amount, date) values (?, ?, ?, ?)", item.Name(), item.PersonID(), item.Amount(), d); err != nil {
		return err
	}
	return nil
}
