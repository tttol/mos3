package awssdk

import (
	"bufio"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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

	reader := bufio.NewReader(r.Body)
	for {
		// Read the chunk size
		chunkSizeStr, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			http.Error(w, "Failed to read chunk size", http.StatusInternalServerError)
			return
		}

		// Parse the chunk size
		chunkSizeStr = strings.TrimSpace(chunkSizeStr)
		chunkSize, err := strconv.ParseInt(chunkSizeStr, 16, 64)
		if err != nil {
			http.Error(w, "Invalid chunk size", http.StatusInternalServerError)
			return
		}

		if chunkSize == 0 {
			// End of chunks
			break
		}

		// Read the chunk data
		chunkData := make([]byte, chunkSize)
		if _, err := io.ReadFull(reader, chunkData); err != nil {
			http.Error(w, "Failed to read chunk data", http.StatusInternalServerError)
			return
		}

		// Skip the chunk signature
		if _, err := reader.ReadString('\n'); err != nil {
			http.Error(w, "Failed to read chunk signature", http.StatusInternalServerError)
			return
		}

		// Write the data to the file
		if _, err := file.Write(chunkData); err != nil {
			http.Error(w, "Failed to write to file", http.StatusInternalServerError)
			return
		}
	}

	slog.Info("File saved to " + savePath)
	w.Write([]byte("File uploaded successfully"))
}
