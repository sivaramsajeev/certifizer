package main

import (
	"os"
	"os/exec"
)

type Certbot struct {
	config *Config
}

func (c *Certbot) run() error {
	cmd := exec.Command("certbot", "--nginx", "--non-interactive", "--agree-tos", "--email", c.config.Email, "-d", c.config.Domain)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
