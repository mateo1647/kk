package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config is the env config; can be overwritten with env vars
type Config struct {
	LogLevel        string   `default:"debug" envconfig:"LOG_LEVEL"` // debugging
	EnvMatchRegex   string   `default:".*"`
	KubeClusterName string   `default:""`
	KubeConfig      string   `default:"" envconfig:"KUBECONFIG"`
	AllowedOrigins  []string `default:"*"`
}

// Get returns the environment configuration
func Get() Config {
	var cfg Config
	err := envconfig.Process("kk", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
