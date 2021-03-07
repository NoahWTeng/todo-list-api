package injectors

import (
	"github.com/NoahWTeng/todo-api-go/src/app/users"
	"github.com/NoahWTeng/todo-api-go/src/infra/db/mongodb"
)

func UsersControllersProvider(repo users.Services) users.Controllers {
	return &users.Repository{UsersServices: repo}
}

func UsersServicesProvider(mongodb *mongodb.Handler) users.Services {
	return &users.Database{
		Handler: mongodb, Collection: "users",
	}
}
