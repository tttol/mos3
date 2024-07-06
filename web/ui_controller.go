package web

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/tttol/mos3/core/util"
)

func S3Handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("S3Handler is called.")
	path := r.URL.Path[len("/s3/"):]
	if r.URL.Query().Get("action") == "dl" {
		download(path)
		return
	}

	s3Objects, err := util.GenerateS3Objects(r, UPLOAD_DIR, path)
	if err != nil {
		slog.Error("GenerateS3Objects error", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var currentPath string
	if path == "" {
		currentPath = "/"
	} else {
		currentPath = path
	}

	dataMap := map[string]interface{}{
		"S3Objects":   s3Objects,
		"Breadcrumbs": util.GenerateBreadcrumbs(path),
		"CurrentPath": currentPath,
	}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		slog.Error("template file error", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	slog.Info("dataMap", "dataMap", dataMap)
	tmpl.Execute(w, dataMap)
}



func RenameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	oldFilename := r.FormValue("oldFilename")
	newFilename := r.FormValue("newFilename")
	err := os.Rename(filepath.Join(UPLOAD_DIR, oldFilename), filepath.Join(UPLOAD_DIR, newFilename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func RenamedirHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	oldDirname := r.FormValue("oldDirname")
	newDirname := r.FormValue("newDirname")
	err := os.Rename(filepath.Join(UPLOAD_DIR, oldDirname), filepath.Join(UPLOAD_DIR, newDirname))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func download(path string) {
	slog.Info("Download file", "path", path)
}
