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
		if _, err := db.Exec(`insert into item (name, person_id, amount, date) values ("A", "a@mail", 100, "2020-01-01 00:00:00")`); err != nil {
			panic(err)
		}
		if _, err := db.Exec(`insert into item (name, person_id, amount, date) values ("B", "b@mail", 100, "2020-01-31 23:59:59")`); err != nil {
			panic(err)
		}
		repo := &sqlite.AccountRepository{DB: db}

		m, err := repo.Month(2020, time.January)
		if err != nil {
			t.Fail()
			return
		}

		t0, err := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
		if err != nil {
			panic(err)
		}
		n0, err := account.NewItem("A", 100, "a@mail", t0)
		if err != nil {
			panic(err)
		}
		t1, err := time.Parse(time.RFC3339, "2020-01-01T23:59:59Z")
		if err != nil {
			panic(err)
		}
		n1, err := account.NewItem("B", 100, "b@mail", t1)
		if err != nil {
			panic(err)
		}
		want := []account.Item{n0, n1}
		if !reflect.DeepEqual(m.Items(), want) {
			t.Errorf("want %#v, but %#v", want, m.Items())
		}
	})
}
