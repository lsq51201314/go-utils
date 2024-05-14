package utils

import "github.com/spf13/viper"

//载入配置
func YamlLoad(config interface{}, file ...string) (err error) {
	//配置文件
	path := "./config.yaml"
	if len(file) > 0 {
		path = file[0]
	}
	//载入配置
	viper.SetConfigFile(path)
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(config)
	return
}
