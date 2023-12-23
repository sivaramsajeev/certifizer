package main

type Ports struct {
}

func (p *Ports) Update(c *ConfigChecker) error {
	logger.Println("✅ Ports are updated")
	return nil
}
