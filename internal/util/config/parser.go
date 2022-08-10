package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

// Parse parses all configuration to a single Config object.
func Parse() *Config {

	return &Config{
		AppConfig: parseAppConfig(),
		LogConfig: parseLogConfig(),
	}
}

// Parse application configurations.
func parseAppConfig() AppConfig {

	content := read("app.yaml")

	cfg := AppConfig{}

	err := yaml.Unmarshal(content, &cfg)

	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return cfg
}

// Parse application configurations.
func parseLogConfig() LogConfig {

	content := read("logger.yaml")

	cfg := LogConfig{}

	err := yaml.Unmarshal(content, &cfg)

	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return cfg
}
