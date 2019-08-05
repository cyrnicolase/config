package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var configPath string

// init 是初始化函数
// 在没有指定路径的情况下，将当前目录下的yaml目录作为配置目录
func init() {
	currentDir, _ := os.Getwd()
	configPath = currentDir + "/yaml"
}

// Boot 是启动函数
func Boot() {
	unmarshalApp()
	unmarshalDatabases()
}

func readYaml(name string) (v *viper.Viper, err error) {
	v = viper.New()
	v.SetConfigName(name)
	v.AddConfigPath(configPath)
	err = v.ReadInConfig()
	if nil != err {
		fmt.Errorf("读取配置文件%s失败%v", name, err)
		return
	}

	return
}
