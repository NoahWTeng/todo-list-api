// +build wireinject

package app

import (
	"github.com/NoahWTeng/todo-api-go/src/infra/injectors"
	"github.com/google/wire"
)

func CreateNewApp() (*Container, error) {
	panic(wire.Build(
		injectors.BaseConfigInjector,
		injectors.HttpServerInjector,
		injectors.RouterInjector,
		injectors.MongodbProvider,
		injectors.UsersServicesProvider,
		injectors.UsersControllersProvider,
		wire.Struct(new(Container), "*"),
	))
}
