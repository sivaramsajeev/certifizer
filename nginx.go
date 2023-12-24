package main

import (
	"fmt"
	"html/template"
	"os"
)

type NginxProxy struct {
	osType            string
	packageManagerCmd string
}

func (n *NginxProxy) setUp() error {
	n.packageManagerCmd, n.osType = checkOSandPackageManager()
	if n.osType == "" || n.packageManagerCmd == "" {
		logger.Fatal("❌ Unsupported OS detected.")
	}
	n.installPackages()

	domain := new(Domain)
	domain.validate()

	bot := &Certbot{
		config: domain.config,
	}
	if err := bot.run(); err != nil {
		logger.Fatal("❌ Certbot failed receiving the certs")
	}

	n.createServers(domain)
	return nil
}

func (n *NginxProxy) restart() error {
	logger.Println("✅ Nginx restarted")
	return nil
}

func (n *NginxProxy) installPackages() {
	logger.Println("✅ Starting installation of ALL required packages...")

	for _, pkg := range []string{"nginx", "certbot", "python3-certbot-nginx"} {
		installPackage(n.packageManagerCmd, pkg)
	}

	logger.Println("✅ ALL the Packages have been installed successfully.")
}

func (n *NginxProxy) createServers(domain *Domain) {
	// The template string
	serverTmpl := `
{{ range .Ports }}
	server {
		listen {{ . }} ssl;
		server_name {{ $.Domain }} www.{{ $.Domain }};

		location / {
			proxy_pass http://localhost:{{ add . $.Offset }};
			proxy_http_version 1.1;
			proxy_set_header Upgrade $http_upgrade;
			proxy_set_header Connection 'upgrade';
			proxy_set_header Host $host;
			proxy_cache_bypass $http_upgrade;
		}

		ssl_certificate /etc/letsencrypt/live/{{ $.Domain }}/fullchain.pem;
		ssl_certificate_key /etc/letsencrypt/live/{{ $.Domain }}/privkey.pem;
		include /etc/letsencrypt/options-ssl-nginx.conf;
		ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;
	}
{{ end }}
`

	data := domain.config

	file, err := os.Create("/etc/nginx/sites-available/default")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	tmpl, err := template.New("nginxConfig").Funcs(template.FuncMap{"add": add}).Parse(serverTmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Configuration written to /etc/nginx/sites-available/default")
}

func add(a, b int) int {
	return a + b
}
