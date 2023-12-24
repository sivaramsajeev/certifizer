package main

import (
	"log"
	"os"
)

var (
	logger         *log.Logger
	configFilePath string
)

type Certifizer struct{}

func (c *Certifizer) start() {
	logger.Println("🚀 Starting the program...")
	config := new(ConfigChecker)
	config.check()

	if config.isFresh {
		logger.Println("✅ Starting a fresh setup...")
		proxy := new(NginxProxy)
		if err := proxy.setUp(); err != nil {
			logger.Fatal("🔥 Proxy set up has failed!!!")
		}

	}

	logger.Println("✅ Program has finished succesfully.")
}

func main() {
	new(Certifizer).start()
}

func init() {
	logger = log.New(os.Stdout, "[Certifizer] ", log.Ldate|log.Ltime)
}
