package injectors

import (
	"github.com/NoahWTeng/todo-api-go/src/app/tasks"
	"github.com/NoahWTeng/todo-api-go/src/infra/db/mongodb"
)

func TasksControllersProvider(repo tasks.Services) tasks.Controllers {
	return &tasks.Repository{TasksRepository: repo}
}

func TasksServicesProvider(mongodb *mongodb.Handler) tasks.Services {
	return &tasks.Database{
		Handler: mongodb, Collection: "tasks",
	}
}
