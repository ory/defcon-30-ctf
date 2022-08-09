package main

import "os"

type Config struct {
	ListenAddress  string
	DataSourceName string
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
		DataSourceName: fromEnvOrDefault("DATA_SOURCE_NAME", "file::memory:?cache=shared"),
	}, nil
}
