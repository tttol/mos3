package web

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/tttol/mos3/core/aws"
	"github.com/tttol/mos3/core/logging"
)

const uploadDir = "./upload"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)

	userAgent := r.Header.Get("User-Agent")
	if strings.Contains(userAgent, "command/s3.ls") {
		aws.Exec(w)
		return
	}

	files, err := os.ReadDir(uploadDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		slog.Error("template file error", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, files)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join(uploadDir, header.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	filename := r.FormValue("filename")
	err := os.Remove(filepath.Join(uploadDir, filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func RenameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	oldFilename := r.FormValue("oldFilename")
	newFilename := r.FormValue("newFilename")
	err := os.Rename(filepath.Join(uploadDir, oldFilename), filepath.Join(uploadDir, newFilename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func MkdirHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	dirname := r.FormValue("dirname")
	err := os.Mkdir(filepath.Join(uploadDir, dirname), os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func RmdirHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	dirname := r.FormValue("dirname")
	err := os.Remove(filepath.Join(uploadDir, dirname))
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
	err := os.Rename(filepath.Join(uploadDir, oldDirname), filepath.Join(uploadDir, newDirname))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
