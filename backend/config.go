package main

import "os"

type config struct {
	ListenAddress  string
	DataSourceName string
}

func fromEnvOrDefault(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func newConfig() (*config, error) {
	return &config{
		ListenAddress:  fromEnvOrDefault("LISTEN_ADDRESS", ":8000"),
		DataSourceName: fromEnvOrDefault("DATA_SOURCE_NAME", "file::memory:?cache=shared"),
	}, nil
}
