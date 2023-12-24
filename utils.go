package main

import (
	"io/ioutil"
	"net"
	"net/http"
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
