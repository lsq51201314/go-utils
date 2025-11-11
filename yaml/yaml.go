package yaml

import (
	"github.com/spf13/viper"
)

// 构建参数
func Build(config any, file ...string) error {
	str := "./config.yaml"
	if len(file) > 0 {
		str = file[0]
	}
	viper.SetConfigFile(str)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(config); err != nil {
		return err
	}
	return nil
}
