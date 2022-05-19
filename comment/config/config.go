package config

import "os"

type Config struct {
	ServerPort    string
	DatabaseHost  string
	DatabasePort  string
	DatabaseUname string
	DatabasePass  string
	DatabaseName  string

	OrgUrl  string
	OrgPort string
}

func InitConfig() Config {
	return Config{
		ServerPort:    get("COMMENT_SERVER_PORT", ":8889"),
		DatabaseHost:  get("COMMENT_DATABASE_URL", "localhost"),
		DatabasePort:  get("COMMENT_DATABASE_PORT", "5433"),
		DatabaseUname: get("COMMENT_DATABASE_USER", "root"),
		DatabasePass:  get("COMMENT_DATABASE_PASSWORD", "root"),
		DatabaseName:  get("COMMENT_DATABASE_NAME", "comment"),
		OrgUrl:        get("ORG_URL", "localhost"),
		OrgPort:       get("ORG_PORT", "8888"),
	}
}

func get(envName string, defaultValue string) string {
	if value, ok := os.LookupEnv(envName); ok {
		return value
	}

	return defaultValue
}
