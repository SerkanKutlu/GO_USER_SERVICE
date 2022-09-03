package config

import "github.com/spf13/viper"

type configurationManager struct {
	applicationConfig *ApplicationConfig
}

func NewConfigurationManager(path string, file string, env string) *configurationManager {
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	appConfig := readApplicationConfigFile(env, file)
	return &configurationManager{
		applicationConfig: appConfig,
	}
}
func readApplicationConfigFile(env string, file string) *ApplicationConfig {

	viper.SetConfigName(file)
	if err := viper.ReadInConfig(); err != nil {
		panic("Can not load application config file 1")
	}
	var appConfig ApplicationConfig
	envSub := viper.Sub(env)
	if err := envSub.Unmarshal(&appConfig); err != nil {
		panic(err.Error())
	}
	return &appConfig
}

func (cm *configurationManager) GetMongoConfiguration() *MongoConfig {
	return &cm.applicationConfig.Mongo
}
func (cm *configurationManager) GetJwtConfiguration() *Jwt {
	return &cm.applicationConfig.Jwt
}
