package yaml

import (
	"github.com/spf13/viper"
)

// 构建参数
func Build(config any) error {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(config); err != nil {
		return err
	}
	return nil
}
