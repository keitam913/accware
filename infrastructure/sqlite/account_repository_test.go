package sqlite_test

import (
	"database/sql"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/keitam913/accware-api/domain/account"
	"github.com/keitam913/accware-api/infrastructure/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func WithDB(f func(db *sql.DB)) {
	db, err := sql.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic(err)
	}
	defer os.Remove("test.sqlite3")
	defer db.Close()
	sc, err := ioutil.ReadFile("../../schema.sql")
	if err != nil {
		panic(err)
	}
	if _, err := db.Exec(string(sc)); err != nil {
		panic(err)
	}
	f(db)
}

func TestMonth(t *testing.T) {
	WithDB(func(db *sql.DB) {
		if _, err := db.Exec(`insert into item (name, person_id, amount, date) values ("A", "b@mail", 100, "2019-12-31 23:59:59")`); err != nil {
			panic(err)
		}
		if _, err := db.Exec(`insert into item (name, person_id, amount, date) values ("B", "a@mail", 100, "2020-01-01 00:00:00")`); err != nil {
			panic(err)
		}
		if _, err := db.Exec(`insert into item (name, person_id, amount, date) values ("C", "b@mail", 100, "2020-01-31 23:59:59")`); err != nil {
			panic(err)
		}
		if _, err := db.Exec(`insert into item (name, person_id, amount, date) values ("D", "b@mail", 100, "2020-02-01 00:00:00")`); err != nil {
			panic(err)
		}
		repo := &sqlite.AccountRepository{DB: db}

		m, err := repo.Month(2020, time.January)
		if err != nil {
			t.Fail()
			return
		}

		tb, err := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
		if err != nil {
			panic(err)
		}
		ib, err := account.NewItem("B", 100, "a@mail", tb)
		if err != nil {
			panic(err)
		}
		tc, err := time.Parse(time.RFC3339, "2020-01-31T23:59:59Z")
		if err != nil {
			panic(err)
		}
		ic, err := account.NewItem("C", 100, "b@mail", tc)
		if err != nil {
			panic(err)
		}
		want := []account.Item{ib, ic}
		if !reflect.DeepEqual(m.Items(), want) {
			t.Errorf("want %#v got %#v", want, m.Items())
		}
	})
}