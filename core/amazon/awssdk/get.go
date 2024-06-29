package awssdk

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func Get(w http.ResponseWriter, r *http.Request) {
	slog.Info("awssdk.Get is called.")
	if r.URL.Path == "" || r.URL.Path == "/" {
		http.Error(w, "File path cannot be empty", http.StatusBadRequest)
		return
	}

	filePath := fmt.Sprintf("upload%s", r.URL.Path) // Relative path from the directory where main.go is executed
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		slog.Error("File not found.", "filePath", filePath)
		http.Error(w, fmt.Sprintf("File %s not found", filePath), http.StatusNotFound)
		return
	}

	slog.Info("Serving file", "filePath", filePath)
	http.ServeFile(w, r, filePath)
	w.Write([]byte("File downloaded successfully"))
}
