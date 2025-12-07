package gyaml

import (
	"io/fs"
	"os"

	"github.com/goccy/go-yaml"
)

// 载入文件
func Load(config any, file string) error {
	buf, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(buf, config)
}

// 保存文件
func Save(config any, file string) error {
	buf, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(file, buf, fs.ModePerm)
}
