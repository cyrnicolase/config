package config

import (
	"testing"
)

func TestUnmarshalApp(t *testing.T) {
	unmarshalApp()

	tApp := AppConfig{
		AppName:    "一个应用名",
		Version:    "0.0.1",
		DateFormat: "2006-01-02 15:04:05",
		RunMode:    "development",
	}

	if tApp != App {
		t.Errorf("App格式希望是%v,实际是%v", tApp, App)
	}
}

func TestUnmarshalDatabases(t *testing.T) {
	unmarshalDatabases()
	tDatabases := DatabaseConfigs{
		Development: DatabaseConfig{
			Host:     "192.168.3.224",
			User:     "opmp",
			DB:       "opmp",
			Password: "123456",
			Sslmode:  false,
		},
		Testing: DatabaseConfig{
			Host:     "192.168.0.11",
			User:     "opmp",
			DB:       "opmp",
			Password: "abc123",
			Sslmode:  false,
		},
		Production: DatabaseConfig{
			Host:     "10.0.0.1",
			User:     "opmp",
			DB:       "opmp",
			Password: "hook@761&%ass,11",
			Sslmode:  true,
		},
	}

	if tDatabases != Databases {
		t.Errorf("Databases格式希望是%v,实际是%v", tDatabases, Databases)
	}
}

func TestingRegisterDatabase(t *testing.T) {
	Boot()
	registerDatabase()

	tDatabase := DatabaseConfig{
		Host:     "192.168.3.224",
		User:     "opmp",
		DB:       "opmp",
		Password: "123456",
		Sslmode:  false,
	}

	if tDatabase != Database {
		t.Errorf("Database 格式期望是%v, 实际是%v", tDatabase, Databases)
	}
}
