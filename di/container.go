package di

import (
	accountapp "github.com/keitam913/accware-api/application"
	"github.com/keitam913/accware-api/account"
	"github.com/keitam913/accware-api/person"
	"github.com/keitam913/accware-api/config"
	"github.com/keitam913/accware-api/rest"

	"github.com/keitam913/accware-api/oidc"

	"github.com/gin-gonic/gin"
)

type Container struct{}

func (c Container) Router() *gin.Engine {
	return rest.NewRouter(c.OIDCService(), c.MonthHandler())
}

func (c Container) MonthHandler() *rest.MonthHandler {
	return &rest.MonthHandler{
		Service:       c.AccountApplicationService(),
		PersonService: c.PersonService(),
	}
}

func (c Container) AccountApplicationService() *accountapp.Service {
	return &accountapp.Service{
		Repository: c.AccountRepository(),
	}
}

func (c Container) AccountRepository() account.Respository {
	return nil
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
	conf, err := config.Load("config.dev.yaml")
	if err != nil {
		panic(err)
	}
	return conf
}
