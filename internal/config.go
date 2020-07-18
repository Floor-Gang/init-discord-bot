package internal

import (
	"log"

	utilConfig "github.com/Floor-Gang/utilpkg/config"
)

// Config structure.
type Config struct {
	Token     string `yaml:"token"`
	Prefix    string `yaml:"prefix"`
	ChannelID string `yaml:"channel"`
	LeadDevID string `yaml:"leadev"`
	Location  string `yaml:"location"`
}

// GetConfig retrieves a configuration.
func GetConfig(configPath string) Config {
	config := Config{
		Token:     "",
		Prefix:    "",
		ChannelID: "",
		LeadDevID: "",
		Location:  configPath,
	}
	err := utilConfig.GetConfig(configPath, &config)

	if err != nil {
		log.Fatalln(err)
	}

	return config
}

// Save saves configuration
func (config *Config) Save() {
	if err := utilConfig.Save(config.Location, config); err != nil {
		log.Println("Failed to save config", err)
	}
}
