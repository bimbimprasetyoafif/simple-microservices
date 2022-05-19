package config

import "os"

type Config struct {
	ServerPort    string
	DatabaseHost  string
	DatabasePort  string
	DatabaseUname string
	DatabasePass  string
	DatabaseName  string
}

func InitConfig() Config {
	return Config{
		ServerPort:    get("ORG_SERVER_PORT", ":8888"),
		DatabaseHost:  get("ORG_DATABASE_URL", "localhost"),
		DatabasePort:  get("ORG_DATABASE_PORT", "5432"),
		DatabaseUname: get("ORG_DATABASE_USER", "root"),
		DatabasePass:  get("ORG_DATABASE_PASSWORD", "root"),
		DatabaseName:  get("ORG_DATABASE_NAME", "organization"),
	}
}

func get(envName string, defaultValue string) string {
	if value, ok := os.LookupEnv(envName); ok {
		return value
	}

	return defaultValue
}
