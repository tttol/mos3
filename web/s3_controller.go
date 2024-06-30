package web

import (
	"log/slog"
	"net/http"
	"text/template"

	"github.com/tttol/mos3/core/util"
)

const uploadDir = "./upload"

func S3Handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("S3Handler is called.")
	path := r.URL.Path[len("/s3/"):]
	if r.URL.Query().Get("ation") == "dl" {
		download(path)
		return
	}

	s3Objects, err := util.GetS3Objects(r, path)
	if err != nil {
		slog.Error("GetS3Objects error", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
