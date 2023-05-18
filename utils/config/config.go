package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type (
	Server struct {
		Host            string
		Port            string
		ShutdownTimeout string
	}

	Log struct {
		Level             string
		DevMode           bool
		DisableStacktrace bool
	}

	Config struct {
		Server  Server
		Log     Log
		Version string
	}
)

func LoadConfig(appName string) (*Config, error) {

	k := koanf.New(".")

	// Yaml config loader
	if err := k.Load(file.Provider("application.yaml"), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	// Env config loader: overrides the yaml config when present
	k.Load(env.Provider(getEnvVarPrefix(appName), ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, getEnvVarPrefix(appName))), "_", ".", -1)
	}), nil)

	// unmarshall the merged config
	var config Config
	if err := k.Unmarshal("", &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return &config, nil
}

func getEnvVarPrefix(appName string) string {
	return appName + "_"
}
