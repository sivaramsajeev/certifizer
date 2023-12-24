package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func printConfigFileUsage() {
	logger.Println(`
Create a config file by the name certifizer.yml either in your home directory or any specific file path.
NOTE: You need to set CERTIFIZER_CONFIG_PATH environment variable if config file is not in your home directory

SAMPLE FILE:
------------
cat ~/certifizer.yml
domain: 
email:
# offset: 7
	
ports:
- 3000
- 5000
	`)
}

func getPublicIP() (string, error) {
	resp, err := http.Get("http://ifconfig.me/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(ip)), nil
}

func checkDNSResolution(domain string) (string, error) {
	ip, err := net.LookupIP(domain)
	if err != nil {
		return "", err
	}

	return ip[0].String(), nil
}

func checkOSandPackageManager() (string, string) {
	pkgManagerCmd := ""
	osType := ""

	switch runtime.GOOS {
	case "linux":
		if _, err := os.Stat("/etc/redhat-release"); err == nil {
			pkgManagerCmd = "yum"
			osType = "redhat"
		} else if _, err := os.Stat("/etc/debian_version"); err == nil {
			pkgManagerCmd = "apt"
			osType = "debian"
		}
	default:
		logger.Panic("❌ Unsupported operating system.")
	}
	return pkgManagerCmd, osType
}

func installPackage(pkgCmd, pkgName string) {
	logger.Println("✅ Startinginstallation of", pkgName)

	Cmd := exec.Command(pkgCmd, "install", "-y", "pkgName")
	err := Cmd.Run()
	if err != nil {
		logger.Panic("❌ Failed to install :", pkgName, err)
	}

	logger.Println("✅ Packages installed successfully for ", pkgName)
}
