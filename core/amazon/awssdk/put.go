package awssdk

import (
	"bufio"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Put(w http.ResponseWriter, r *http.Request) {
	slog.Info("awssdk.Put is called.")
	filePath := r.URL.Path
	if filePath == "" || filePath == "/" {
		http.Error(w, "File path cannot be empty", http.StatusBadRequest)
		return
	}

	savePath := filepath.Join("upload", filePath)
	dir := filepath.Dir(savePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		http.Error(w, "Failed to create directory", http.StatusInternalServerError)
		return
	}

	file, err := os.Create(savePath)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	write(r.Body, file, w)

	slog.Info("File saved to " + savePath)
	w.Write([]byte("File uploaded successfully"))
}

func write(data io.ReadCloser, file *os.File, w http.ResponseWriter) {
	reader := bufio.NewReader(data)
	for {
		chunk, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				slog.Info("End of file.")
				break
			}
			slog.Error("Failed to read chunk", "error", err)
			http.Error(w, "Failed to read chunk", http.StatusInternalServerError)
			return
		}
		slog.Info("Reading chunk...", "chunk", chunk)

		if strings.Contains(chunk, "chunk-signature=") {
			slog.Info("Skipping chunk signature... " + chunk)
			continue
		}

		if _, err := file.Write([]byte(chunk)); err != nil {
			slog.Error("Failed to write to file", "error", err)
			http.Error(w, "Failed to write to file", http.StatusInternalServerError)
			return
		}
		slog.Info("Write to file...", "chunk", chunk)
	}
}
