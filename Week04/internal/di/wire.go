// +build wireinject

package di

import (
	"Go-000/Week04/internal/dao"
	"Go-000/Week04/internal/server/http"
	"Go-000/Week04/internal/service"
	"github.com/google/wire"
)

func InitApp() (*App, func(), error){
	panic(wire.Build(dao.Provider, service.New, http.New, NewApp))
}
