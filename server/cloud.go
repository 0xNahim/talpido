package server

import (
	"os/exec"
	"regexp"
	"io/ioutil"
	"time"
)

func StartTunnel() string{
	cmd := exec.Command("cloudflared", "tunnel", "--url", "http://localhost:8080", "--logfile", "./tunnel.txt")
	cmd.Start()
	time.Sleep(5 * time.Second)
	url := extractTunnelURL()
	return url
}

func extractTunnelURL() (string) {
	data, _ := ioutil.ReadFile("./tunnel.txt")
	re := regexp.MustCompile(`https://[^ ]+\.trycloudflare\.com`)
	match := re.FindString(string(data))
	cmd1 := exec.Command("rm", "-rf", "tunnel.txt")
	cmd1.Run()
	return match
}