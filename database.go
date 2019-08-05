package config

import (
	"errors"
	"fmt"
)

// DatabaseConfig 是app的配置结构体
type DatabaseConfig struct {
	Host     string
	User     string
	DB       string `mapstructure:"database"`
	Password string
	Sslmode  bool
}

// DatabaseConfigs 是整个数据库配置文件结构体
type DatabaseConfigs struct {
	Development DatabaseConfig
	Testing     DatabaseConfig
	Production  DatabaseConfig
}

// Databases 是所有环境的配置情况
var Databases DatabaseConfigs

// Database 是DatabaseConfig的具体对象
var Database DatabaseConfig

func unmarshalDatabases() {
	v, _ := readYaml("databases")
	v.Unmarshal(&Databases)
}

func registerDatabase() error {
	runMode := App.RunMode
	switch runMode {
	case "Development":
		Database = Databases.Development
	case "Testing":
		Database = Databases.Testing
	case "Production":
		Database = Databases.Production
	default:
		msg := fmt.Sprintf("未找到运行模式:%s 对应数据库配置", runMode)
		return errors.New(msg)
	}

	return nil
}
