package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Certbot struct {
	config *Config
}

func (c *Certbot) run() error {
	if c.certsExists() {
		logger.Println(" Certs are already generated. No need to invoke again.")
		return nil
	}
	cmd := exec.Command("certbot", "--nginx", "--non-interactive", "--agree-tos", "--email", c.config.Email, "-d", c.config.Domain)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (c *Certbot) certsExists() bool {
	_, err := os.Stat(fmt.Sprintf("/etc/letsencrypt/live/%s/fullchain.pem", c.config.Domain))
	return err == nil
}
