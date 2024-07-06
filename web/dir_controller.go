package web

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

func MkdirHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	currentPath := r.FormValue("currentPath")
	dirname := r.FormValue("dirname")
	dir := filepath.Join(UPLOAD_DIR, currentPath, dirname)
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		slog.Error("failed to mkdir", "target directory", dir, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slog.Info("Success mkdir", "target directory", dir)

	http.Redirect(w, r, "/s3/"+currentPath, http.StatusFound)
}

func RmdirHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	dirname := r.FormValue("dirname")
	err := os.Remove(filepath.Join(UPLOAD_DIR, dirname))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
