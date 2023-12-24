package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

type ConfigChecker struct {
	isFresh          bool
	isUpdateRequired bool
}

func (c *ConfigChecker) check() {
	checkConfigFile()
	c.isFresh = isFreshInstallation()

	logger.Println("✅ Config checks are done")
}

func (c *ConfigChecker) displayPortMappingInfo() {
	logger.Println("✅ The port mappings")
}

func checkConfigFile() {
	configFilePath = os.Getenv("CERTIFIZER_CONFIG_PATH")
	if configFilePath == "" {
		homePath, err := os.UserHomeDir()
		if err != nil {
			logger.Fatal("🔥 Error reading home path")
		}
		configFilePath = filepath.Join(homePath, "certifizer.yml")
		logger.Println("✅ Home path: ", configFilePath)
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		printConfigFileUsage()
		logger.Fatal("🔥 Please re-try after creating the config file as per the instructions...")
	}
}

func isFreshInstallation() bool {
	err := exec.Command("nginx", "-v").Run()
	return err != nil
}
