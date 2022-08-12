package main

import "os"

type Config struct {
	ListenAddress  string
	DataSourceName string
	FlagSeed       string
}

func fromEnvOrDefault(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func newConfig() (*Config, error) {
	return &Config{
		ListenAddress:  fromEnvOrDefault("LISTEN_ADDRESS", ":8000"),
		DataSourceName: fromEnvOrDefault("DSN", "sqlite3://file::memory:?cache=shared"),
		FlagSeed:       fromEnvOrDefault("FLAG_SEED", ""),
	}, nil
}
