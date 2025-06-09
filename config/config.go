package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds all configuration options
type Config struct {
	APIKey      string `mapstructure:"api_key"`
	InputFolder string `mapstructure:"input_folder"`
	DropletName string `mapstructure:"droplet_name"`
	Region      string `mapstructure:"region"`
	Size        string `mapstructure:"size"`
	Image       string `mapstructure:"image"`
}

// LoadConfig loads the config from the file
func LoadConfig(configFile string) (*Config, error) {
	viper.SetConfigFile(configFile)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return &config, nil
}
