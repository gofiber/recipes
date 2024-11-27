package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct for structuring the config data, can be extended accordingly
type Config struct {
	Database struct {
		URL     string `yaml:"url"`
		DB      string `yaml:"db"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

// NewConfig  a function to read the configuration file and return config struct
func NewConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
