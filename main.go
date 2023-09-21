package main

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type Sizer interface {
	Size() int64
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(renderIndexHTML())
}

func serveFile(w http.ResponseWriter, r *http.Request, filePath string) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		// Handle file not found or other errors
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Get the file's content type based on its extension
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		// If the content type is not recognized, set a default
		contentType = "application/octet-stream"
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", contentType)

	// Copy the file's content to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		// Handle copy error, e.g., if the client disconnects
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 32<<20)

	file, multipartFileHeader, err := r.FormFile("upload_file")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	uniqueFilename := uuid.New().String()

	// Extract the file extension from the original filename
	fileExtension := path.Ext(multipartFileHeader.Filename)

	// Specify the directory where you want to save the uploaded file
	saveDirectory := "file/" // Replace with your desired directory

	// Create a new file with the same extension in the specified directory
	newFilePath := filepath.Join(saveDirectory, uniqueFilename+fileExtension)

	newFile, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Copy the contents of the uploaded file to the new file
	_, err = io.Copy(newFile, file)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// // Log relevant information
	// log.Printf("Uploaded file saved as: %s\n", newFilePath)
	// log.Printf("Original Filename: %#v\n", multipartFileHeader.Filename)
	// log.Printf("File Size: %#v bytes\n", file.(Sizer).Size())
	// log.Printf("MIME Type: %#v\n", http.DetectContentType(fileHeader))

	// Respond with a success message or redirect as needed
	http.Redirect(w, r, fmt.Sprintf("/?file=%s", uniqueFilename+fileExtension), http.StatusSeeOther)
}

func handleFile(w http.ResponseWriter, r *http.Request) {
	var url = r.URL.String()[6:]
	println(url)

	if strings.Contains(url, "/") || strings.Contains(url, "..") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serveFile(w, r, "file/"+url)

	// w.WriteHeader(http.StatusOK)
}

func main() {
	_ = os.Mkdir("file", os.ModePerm)

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/file/", handleFile)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8090", nil)
}
