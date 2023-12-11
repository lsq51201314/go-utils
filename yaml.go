package utils

import "github.com/spf13/viper"

func YamlLoad(config interface{}) (err error) {
	viper.SetConfigFile("./config.yaml")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(config)
	return
}
