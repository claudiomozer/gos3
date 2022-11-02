package services

import "github.com/spf13/viper"

type EnvironmentVars struct {
	AccessKey string `mapstructure:"ACCESS_KEY"`
	SecretKey string `mapstructure:"SECRET_KEY"`
}

func LoadEnvironmentVarsFromFile(filePath string) (env *EnvironmentVars, err error) {
	if len(filePath) > 0 {
		viper.SetConfigName("app")
		viper.SetConfigType("env")
		viper.AddConfigPath(filePath)
	} else {
		viper.SetConfigFile("app.env")
	}
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&env)
	return
}
