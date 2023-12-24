package main

import (
	"os"
	"os/exec"
	"runtime"
)

type NginxProxy struct {
}

func (n *NginxProxy) setUp() error {
	installPackages()
	new(Domain).validate()
	return nil
}

func (n *NginxProxy) restart() error {
	logger.Println("✅ Nginx restarted")
	return nil
}

func installPackages() {
	logger.Println("✅ Starting package installation...")

	var pkgManagerCmd string
	switch runtime.GOOS {
	case "linux":
		if _, err := os.Stat("/etc/redhat-release"); err == nil {
			pkgManagerCmd = "yum"
		} else if _, err := os.Stat("/etc/debian_version"); err == nil {
			pkgManagerCmd = "apt"
		} else {
			logger.Panic("❌ Unsupported Linux distribution.")
		}
	default:
		logger.Panic("❌ Unsupported operating system.")
	}

	installNginxCmd := exec.Command(pkgManagerCmd, "install", "-y", "nginx")
	err := installNginxCmd.Run()
	if err != nil {
		logger.Panic("❌ Failed to install Nginx:", err)
	}

	installCertbotCmd := exec.Command(pkgManagerCmd, "install", "-y", "certbot")
	err = installCertbotCmd.Run()
	if err != nil {
		logger.Panic("❌ Failed to install Certbot:", err)
	}

	logger.Println("✅ Packages installed successfully.")
}
