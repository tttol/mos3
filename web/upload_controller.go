package web

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func UploadIndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/upload.html")
	if err != nil {
		slog.Error("template file error", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dataMap := map[string]interface{}{
		"CurrentPath": r.URL.Query().Get("currentPath"),
	}
	tmpl.Execute(w, dataMap)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	path := r.FormValue("currentPath")
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join(UPLOAD_DIR, path, header.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, filepath.Join("/s3", path), http.StatusFound)
}
