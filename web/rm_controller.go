package web

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("RemoveHandler is called.")
	if r.Method != "POST" {
		slog.Error("Invalid request method", "method", r.Method)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	path := r.FormValue("path")[len("/s3/"):]
	err := os.Remove(filepath.Join(UPLOAD_DIR, path))
	if err != nil {
		slog.Error("Failed to remove.", "error", err, "path", path)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slog.Info("Remove file success.", "path", path)

	http.Redirect(w, r, "/s3", http.StatusFound)
}
