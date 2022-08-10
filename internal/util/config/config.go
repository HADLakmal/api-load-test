package config

import "github.com/tryfix/log"

// Config holds all other config structs.
type Config struct {
	AppConfig      AppConfig
	LogConfig      LogConfig
	ServicesConfig []ServiceConfig
}

// AppConfig holds application configurations.
type AppConfig struct {
	Name     string     `yaml:"name"`
	Mode     string     `yaml:"mode"`
	Host     string     `yaml:"host"`
	Port     int        `yaml:"port"`
	Timezone string     `yaml:"timezone"`
	Metrics  Metrics    `yaml:"metrics"`
	Data     DataConfig `yaml:"data"`
}

type DataConfig struct {
	RadiusSyncInterval int `yaml:"radius_sync_interval"`
}

// Metrics holds application metric configurations.
type Metrics struct {
	Enabled bool   `yaml:"enabled"`
	Port    int    `yaml:"port"`
	Route   string `yaml:"route"`
}

// LogConfig holds application log configurations.
type LogConfig struct {
	Level           log.Level `yaml:"level"`
	Remote          bool      `yaml:"remote_log"`
	FilePathEnabled bool      `yaml:"file_path_enabled"`
	Colors          bool      `yaml:"colors"`
	Console         bool      `yaml:"console"`
	File            bool      `yaml:"file"`
	Directory       string    `yaml:"directory"`
}

// ServiceConfig holds service configurations.
type ServiceConfig struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Timeout int    `yaml:"timeout"`
}
