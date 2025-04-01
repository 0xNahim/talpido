package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("IP")
	if ip == "" {
		http.Error(w, "IP no proporcionada en la cabecera", http.StatusBadRequest)
		return
	}

	fileData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fileName := fmt.Sprintf("%s.zip", ip)
	err = os.WriteFile(fileName, fileData, 0644)
	if err != nil {
		http.Error(w, "Error al guardar el archivo", http.StatusInternalServerError)
		return
	}
	fmt.Println("Information received from the IP: "+ ip)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Archivo guardado con éxito"))
}

func serveTopoFiles(w http.ResponseWriter, r *http.Request) {
	dirPath := "./topo"

	
	if r.URL.Path != "/topo" {
	
		filePath := filepath.Join(dirPath, r.URL.Path[len("/topo/"):]) 

		_, err := os.Stat(filePath)
		if err != nil {
			http.Error(w, "Archivo no encontrado", http.StatusNotFound)
			return
		}

	
		w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
		http.ServeFile(w, r, filePath)
		return
	}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		http.Error(w, "No se pudo leer la carpeta", http.StatusInternalServerError)
		return
	}

	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Archivos en ./topo</h1><ul>"))
	for _, file := range files {
		if file.IsDir() {
		}
		w.Write([]byte(fmt.Sprintf("<li><a href=\"/topo/%s\">%s</a></li>", file.Name(), file.Name())))
	}
	w.Write([]byte("</ul>"))
}



func Start() {
	http.HandleFunc("/topo", serveTopoFiles)
	http.Handle("/topo/", http.StripPrefix("/topo/", http.FileServer(http.Dir("./topo"))))
	http.HandleFunc("/upload", handleUpload)

	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
