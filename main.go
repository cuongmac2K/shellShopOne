package main

import (
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/generator-ssl", generatorSsl)
	http.ListenAndServe(":11111", nil)
}

func generatorSsl(writer http.ResponseWriter, request *http.Request) {
	func() {
		cmd := exec.Command("cp", "/shopone/config/nginx-ssl-cloudflare-by-docker/setup-ssl/200/default80.conf", "/shopone/config/nginx-ssl-cloudflare-by-docker/setup-ssl/conf.d/default80.conf")
		cmd.Run()
	}()
	func() {
		cmd := exec.Command("docker", "restart", "cloudflarecertbot_ssl_setup", "nginx_ssl_setup")
		cmd.Run()
	}()
	func() {
		domain := request.FormValue("domain")
		cmd := exec.Command("docker", "exec", "cloudflarecertbot_ssl_setup", "certbot", "certonly", "--webroot", "--webroot-path", "/var/www/html", "-m", "'manhbv.giaminh@gmail.com'", "-d", domain)
		cmd.Run()
	}()
	func() {
		cmd := exec.Command("cp", "/shopone/config/nginx-ssl-cloudflare-by-docker/setup-ssl/301/default80.conf", "/shopone/config/nginx-ssl-cloudflare-by-docker/setup-ssl/conf.d/default80.conf")
		cmd.Run()
	}()
	func() {
		cmd := exec.Command("docker", "restart", "cloudflarecertbot_ssl_setup", "nginx_ssl_setup", "nginx")
		cmd.Run()
	}()
}
