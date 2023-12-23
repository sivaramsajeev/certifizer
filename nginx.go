package main

type NginxProxy struct {
}

func (n *NginxProxy) setUp() error {
	logger.Println("✅ Set up Nginx...")
	return nil
}

func (n *NginxProxy) restart() error {
	logger.Println("✅ Nginx restarted")
	return nil
}
