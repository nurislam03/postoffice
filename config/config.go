package config

import (
	"log"
	"sync"

	// viper/remote
	_ "github.com/spf13/viper/remote"
)

// Config ...
type Config struct {
}

func loadConfig() {
	log.Println("Loading configurations...")
	log.Println("Configurations loaded")
}

var config *Config
var configOnce = &sync.Once{}

// NewConfig ...
func NewConfig() *Config {
	configOnce.Do(func() {
		loadConfig()
		config = &Config{
		}
	})
	return config
}
