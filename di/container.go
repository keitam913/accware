package di

import (
	"database/sql"
	"io/ioutil"
	"os"

	"github.com/keitam913/accware-api/account"
	accountapp "github.com/keitam913/accware-api/application"
	"github.com/keitam913/accware-api/config"
	"github.com/keitam913/accware-api/person"
	"github.com/keitam913/accware-api/rest"
	"github.com/keitam913/accware-api/sqlite"

	"github.com/keitam913/accware-api/oidc"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Container struct{}

func (c Container) Router() *gin.Engine {
	return rest.NewRouter(c.OIDCService(), c.MonthHandler(), c.ItemHandler())
}

func (c Container) MonthHandler() *rest.MonthHandler {
	return &rest.MonthHandler{
		Service:       c.AccountApplicationService(),
		PersonService: c.PersonService(),
	}
}

func (c Container) ItemHandler() *rest.ItemHandler {
	return &rest.ItemHandler{
		Service: c.AccountApplicationService(),
	}
}

func (c Container) AccountApplicationService() *accountapp.Service {
	return &accountapp.Service{
		Repository: c.AccountRepository(),
	}
}

func (c Container) AccountRepository() account.Respository {
	return &sqlite.AccountRepository{
		DB: c.DB(),
	}
}

func (c Container) DB() *sql.DB {
	db, err := sql.Open("sqlite3", "accware.sqlite3")
	if err != nil {
		panic(err)
	}
	sc, err := ioutil.ReadFile("./schema.sql")
	if err != nil {
		panic(err)
	}
	if _, err := db.Exec(string(sc)); err != nil {
		panic(err)
	}
	return db
}

func (c Container) OIDCService() *oidc.Service {
	return &oidc.Service{}
}

func (c Container) PersonService() person.Service {
	return &config.PersonService{
		Config: c.Config(),
	}
}

func (c Container) Config() *config.Config {
	conf, err := config.Load(os.Args[1])
	if err != nil {
		panic(err)
	}
	return conf
}
