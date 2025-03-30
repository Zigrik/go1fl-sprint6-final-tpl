package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// The handler that returns the index form
func HandleIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "The request was not received using the GET method", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	indexPath := filepath.Join("..", "index.html")

	file, err := os.ReadFile(indexPath)
	if err != nil {
		http.Error(res, "Couldn't open the Index file", http.StatusInternalServerError)
		return
	}

	res.Write(file)
}

// Handler converting incoming file from/to Morse code
func HandleUpload(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "The request was not received using the POST method", http.StatusInternalServerError)
		return
	}
	req.ParseMultipartForm(10 << 20)

	// getting the file from the form
	file, header, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "Error receiving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// read the entire file into a byte slice
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// translate into a string and replace the hyphenation characters with spaces
	fileContent := string(data)
	fileContent = strings.ReplaceAll(fileContent, "\n", " ")
	fileContent = strings.ReplaceAll(fileContent, "\r", " ")

	// сhecking the file for convertibility
	out, err := service.СheckMorseAndConvert(fileContent)
	if err != nil {
		http.Error(res, "Failed to read file", http.StatusInternalServerError)
		return
	}

	dirPath := "../results"

	// сhecking if there is a directory to save
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			http.Error(res, "Directory creation error", http.StatusInternalServerError)
			return
		}
	}

	// generate the file name based on the current time and the extension of the original file.
	ext := filepath.Ext(header.Filename)
	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("../results/result_%s%s", timestamp, ext)

	resFile, err := os.Create(filename)
	if err != nil {
		http.Error(res, "File creation error", http.StatusInternalServerError)
		return
	}
	defer resFile.Close()

	// writing the results to a file.
	_, err = fmt.Fprint(resFile, out)
	if err != nil {
		http.Error(res, "Data recording error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(res, "The file has been successfully uploaded to the server and converted")
}
