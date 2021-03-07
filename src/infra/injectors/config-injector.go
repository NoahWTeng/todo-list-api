package injectors

import (
	"github.com/NoahWTeng/todo-api-go/config"
	"github.com/spf13/viper"
	"log"
)

func BaseConfigInjector() (*config.GlobalConfig, error) {

	viper.SetConfigFile("./config.yml")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	var config config.GlobalConfig
	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatal(err)
	}

	return &config, err
}
