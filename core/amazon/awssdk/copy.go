package awssdk

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

// Handle CopyObject requests from AWS SDK
func Copy(w http.ResponseWriter, r *http.Request) {
	slog.Info("awssdk.Copy is called.")
	from, to := r.Header.Get("X-Amz-Copy-Source"), r.URL.Path
	if from == "" || from == "/" || to == "" || to == "/" {
		slog.Error("Invalid filepath.", "from", from, "to", to)
		http.Error(w, "File path cannot be empty", http.StatusBadRequest)
		return
	}

	fromFile, err := os.Open(filepath.Join("upload", from))
	if err != nil {
		slog.Error("Failed to open source file.", "error", err)
		http.Error(w, "Failed to open source file", http.StatusInternalServerError)
		return
	}
	defer fromFile.Close()

	toFile, err := os.Create(filepath.Join("upload", to))
	if err != nil {
		slog.Error("Failed to create destination file.", "error", err)
		http.Error(w, "Failed to create destination file", http.StatusInternalServerError)
		return
	}
	defer toFile.Close()

	// Copy the contents from source to destination
	_, err = io.Copy(toFile, fromFile)
	if err != nil {
		slog.Error("Failed to copy file contents.", "from", fromFile, "to", toFile, "error", err)
		http.Error(w, "Failed to copy file contents", http.StatusInternalServerError)
		return
	}
	slog.Info("File copied successfully.", "from", fromFile.Name(), "to", toFile.Name())

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File copied successfully"))
}
