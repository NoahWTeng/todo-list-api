package injectors

import (
	"github.com/NoahWTeng/todo-api-go/config"
	"github.com/NoahWTeng/todo-api-go/src/infra/db/mongodb"
)

func MongodbProvider(config *config.GlobalConfig) *mongodb.Handler {
	return mongodb.NewMongodbConnection(&config.MongodbVariables)
}
