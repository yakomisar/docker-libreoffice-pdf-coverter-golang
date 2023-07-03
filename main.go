package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

func doc2pdfHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := path.Ext(header.Filename)

	tempFile, err := ioutil.TempFile("", "*"+ext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cmd := exec.Command(
		"libreoffice",
		"--invisible",
		"--headless",
		"--nologo",
		"--convert-to",
		"pdf",
		"--outdir",
		os.TempDir(),
		tempFile.Name(),
	)
	err = cmd.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	convertedFile := strings.TrimSuffix(tempFile.Name(), ext) + ".pdf"
	pdf, err := os.Open(convertedFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.Remove(convertedFile)

	w.Header().Set("Content-Type", "application/pdf")
	_, err = io.Copy(w, pdf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/doc2pdf", doc2pdfHandle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "6000"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
