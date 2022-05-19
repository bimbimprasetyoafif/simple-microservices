package config

import "os"

type Config struct {
	ServerPort        string
	OrgServerUrl      string
	OrgServerPort     string
	CommentServerUrl  string
	CommentServerPort string
}

func InitConfig() Config {
	return Config{
		ServerPort:        get("API_SERVER_PORT", ":8000"),
		OrgServerUrl:      get("ORG_SERVER_URL", "localhost"),
		OrgServerPort:     get("ORG_SERVER_PORT", "8888"),
		CommentServerUrl:  get("COMMENT_SERVER_URL", "localhost"),
		CommentServerPort: get("COMMENT_SERVER_PORT", "8889"),
	}
}

func get(envName string, defaultValue string) string {
	if value, ok := os.LookupEnv(envName); ok {
		return value
	}

	return defaultValue
}
