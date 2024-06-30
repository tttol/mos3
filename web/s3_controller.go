package web

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

const uploadDir = "./upload"

type S3Object struct {
	FullPath string
	Name     string
	IsDir    bool
}

func S3Handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("S3Handler is called.")
	path := r.URL.Path[len("/s3/"):]
	if r.URL.Query().Get("ation") == "dl" {
		download(path)
		return
	}

	dirEntry, err := os.ReadDir(filepath.Join(uploadDir, path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var s3Objects []S3Object
	for _, entry := range dirEntry {
		var obj S3Object
		obj.Name = entry.Name()
		obj.FullPath = filepath.Join(r.URL.Path, entry.Name())
		obj.IsDir = entry.IsDir()
		slog.Info("S3Object", "data", obj)

		s3Objects = append(s3Objects, obj)
	}
	slog.Info("S3Objects", "data", s3Objects)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		slog.Error("template file error", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, s3Objects)
}
func download(path string) {
	slog.Info("Download file", "path", path)
}
