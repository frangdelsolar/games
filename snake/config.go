package snake

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2" // Assuming you've installed the library
)

var appConfig *AppConfig

type AppConfig struct {
	LogLevel string `yaml:"log_level"`
	Rows     int    `yaml:"rows"`
	Cols     int    `yaml:"cols"`
}

func LoadConfig(filePath string) (*AppConfig, error) {
	if appConfig != nil {
		return appConfig, nil
	}

	data, err := os.ReadFile(filePath) // Or use yaml.Unmarshal for structured reading
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Optionally use os.ExpandEnv to handle environment variables within the YAML:
	// replaced := os.ExpandEnv(string(data))

	var cfg AppConfig
	err = yaml.Unmarshal(data, &cfg) // Or yaml.UnmarshalStrict for stricter parsing
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	appConfig = &cfg

	return &cfg, nil
}
