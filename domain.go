package main

type Domain struct {
	config *Config
}

func (d *Domain) validate() {
	d.init()
	d.confirmOwnership()
}

func (d *Domain) init() {
	d.config = readConfig()
	d.config.validate()
}

func (d *Domain) confirmOwnership() {
	publicIp, err := getPublicIP()
	if err != nil {
		logger.Fatal("❌ Unable to get public IP. Check connectivity...")
	}

	dnsIp, err := checkDNSResolution(d.config.Domain)
	if err != nil {
		logger.Fatal("❌ Failed DNS resolution for ", d.config.Domain)
	}

	if publicIp != dnsIp {
		logger.Fatal("❌ Ownership of domain can't be verified. FAILED!!!")
	}

	logger.Println("✅ Domain verification passed. We're golden")
}
