package sqlite_test

import (
	"database/sql"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/keitam913/accware/api/account"
	"github.com/keitam913/accware/api/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func WithDB(f func(db *sql.DB)) {
	db, err := sql.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic(err)
	}
	defer os.Remove("test.sqlite3")
	defer db.Close()
	sc, err := ioutil.ReadFile("../schema.sql")
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
		if _, err := db.Exec(`insert into item (id, name, person_id, amount, date) values ("0", "A", "b@mail", 100, "2019-12-31 23:59:59")`); err != nil {
			panic(err)
		}
		if _, err := db.Exec(`insert into item (id, name, person_id, amount, date) values ("1", "B", "a@mail", 100, "2020-01-01 00:00:00")`); err != nil {
			panic(err)
		}
		if _, err := db.Exec(`insert into item (id, name, person_id, amount, date) values ("2", "C", "b@mail", 100, "2020-01-31 23:59:59")`); err != nil {
			panic(err)
		}
		if _, err := db.Exec(`insert into item (id, name, person_id, amount, date) values ("3", "D", "b@mail", 100, "2020-02-01 00:00:00")`); err != nil {
			panic(err)
		}
		repo := &sqlite.AccountRepository{DB: db}

		m, err := repo.Month(2020, time.January)
		if err != nil {
			t.Fail()
			return
		}

		t0, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-01-01 00:00:00", time.Local)
		if err != nil {
			panic(err)
		}
		wantLen := 2
		want0, err := account.NewItemWithID("1", "B", 100, "a@mail", t0)
		if err != nil {
			panic(err)
		}
		t1, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-01-31 23:59:59", time.Local)
		if err != nil {
			panic(err)
		}
		want1, err := account.NewItemWithID("2", "C", 100, "b@mail", t1)
		if err != nil {
			panic(err)
		}
		if len(m.Items()) != wantLen {
			t.Errorf("len(items) = %d; want %d", len(m.Items()), wantLen)
			return
		}
		if !reflect.DeepEqual(m.Items()[0], want0) {
			t.Errorf("items[0] = %+v; want %+v", m.Items()[0], want0)
		}
		if !reflect.DeepEqual(m.Items()[1], want1) {
			t.Errorf("items[1] == %+v; want %+v", m.Items()[1], want1)
		}
	})
}

func TestAdd(t *testing.T) {
	WithDB(func(db *sql.DB) {
		repo := &sqlite.AccountRepository{DB: db}
		ti, err := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
		if err != nil {
			panic(err)
		}
		item, err := account.NewItem("a", 100, "a@mail", ti)
		if err != nil {
			panic(err)
		}
		if err := repo.Add(item); err != nil {
			t.Fatal(err)
			return
		}
		row := db.QueryRow("select name, person_id, amount, date from item")

		var (
			name, personID, date string
			amount               int
		)
		if err := row.Scan(&name, &personID, &amount, &date); err != nil {
			t.Fatal(err)
			return
		}
		got := []interface{}{name, personID, amount, date}
		want := []interface{}{"a", "a@mail", 100, "2020-01-01 00:00:00"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, want %#v", got, want)
		}
	})
}

func TestRemove(t *testing.T) {
	WithDB(func(db *sql.DB) {
		repo := &sqlite.AccountRepository{DB: db}
		if _, err := db.Exec(`insert into item (id, name, person_id, amount, date) values ("0", "A", "b@mail", 100, "2019-12-31 23:59:59")`); err != nil {
			t.Fatal(err)
		}
		if _, err := db.Exec(`insert into item (id, name, person_id, amount, date) values ("1", "B", "a@mail", 100, "2020-01-01 00:00:00")`); err != nil {
			t.Fatal(err)
		}
		if err := repo.Remove("0"); err != nil {
			t.Fatal(err)
		}
		row, err := db.Query("select id from item")
		if err != nil {
			t.Fatal(err)
		}
		var got []string
		for row.Next() {
			var id string
			if err := row.Scan(&id); err != nil {
				t.Fatal(err)
			}
			got = append(got, id)
		}
		want := []string{"1"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ids = %+v; want %+v", got, want)
		}
	})
}
