package web

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/tttol/mos3/core/util"
)

func S3Handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("S3Handler is called.")
	path := r.URL.Path[len("/s3/"):]

	s3Objects, err := util.GenerateS3Objects(r, UPLOAD_DIR_PATH, GetDirPath(path))
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

	if r.URL.Query().Get("action") == "dl" {
		_, status, err := download(w, path)
		if err != nil {
			http.Error(w, err.Error(), status)
			return
		}
	}

	slog.Info("dataMap", "dataMap", dataMap)
	tmpl.Execute(w, dataMap)
}

// It returns the number of bytes copied and the first error encountered while copying, if any.
func download(w http.ResponseWriter, path string) (n int64, httpStatus int, err error) {
	slog.Info("Start downloading file", "path", path)

	file, err := os.Open(filepath.Join(UPLOAD_DIR_PATH, path))
	if err != nil {
		slog.Error("File open error", "error", err)
		return 0, http.StatusNotFound, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		slog.Error("File stat error", "error", err)
		return 0, http.StatusNotFound, err
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+stat.Name())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", string(rune(stat.Size())))

	writtenByteSize, err := io.Copy(w, file)
	if err != nil {
		slog.Error("File copy error", "error", err)
		return 0, http.StatusInternalServerError, err
	}

	slog.Info("Successful to dowload file", "path", path)
	return writtenByteSize, http.StatusOK, nil
}

func GetDirPath(path string) string {
	re := regexp.MustCompile(`\.\w+$`)
	if re.MatchString(path) {
		return filepath.Dir(path)
	} else {
		return path
	}
}
