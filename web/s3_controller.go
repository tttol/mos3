package web

import (
	"log/slog"
	"net/http"
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
		slog.Error("GetS3Objects error", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dataMap := map[string]interface{}{
		"S3Objects":   s3Objects,
		"Breadcrumbs": util.GenerateBreadcrumbs(path),
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
func download(path string) {
	slog.Info("Download file", "path", path)
}
