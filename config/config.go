package config

import "github.com/NoahWTeng/todo-api-go/src/infra/db/mongodb"

type BaseVariables struct {
	Environment string `mapstructure:"environment" json:"environment"`
	Port        int    `mapstructure:"port" json:"port"`
}

type GlobalConfig struct {
	BaseVariables    *BaseVariables `mapstructure:"base" json:"base"`
	MongodbVariables mongodb.Config `mapstructure:"mongodb" json:"mongodb"`
}
