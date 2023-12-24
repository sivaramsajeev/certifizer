package main

import (
	"io/ioutil"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Domain string `yaml:"domain"`
	Email  string `yaml:"email"`
	Ports  []int  `yaml:"ports"`
	Offset int    `yaml:"offset,omitempty"`
}

func (config *Config) validate() {
	if config.Domain == "" {
		logger.Panic("❌ Domain is empty in the configuration.")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(config.Email) {
		logger.Panic("❌ Invalid email format in the configuration.")
	}

	if len(config.Ports) == 0 {
		logger.Panic("❌ No ports specified in the configuration.")
	}

	logger.Println("Config validations passed")
}

func readConfig() *Config {
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		logger.Panic("❌ Failed to read the configuration file:", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logger.Panic("❌ Failed to unmarshal YAML:", err)
	}

	return &config
}
