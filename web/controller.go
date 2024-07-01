package web

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/tttol/mos3/core/amazon/awscli"
	"github.com/tttol/mos3/core/amazon/awssdk"
	"github.com/tttol/mos3/core/logging"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("IndexHandler is called.")
	logging.LogRequest(r)

	// AWS CLI request
	userAgent := r.Header.Get("User-Agent")
	if strings.Contains(userAgent, "command/s3.ls") {
		awscli.Ls(w)
		return
	} else if strings.Contains(userAgent, "command/s3.cp") {
		awscli.Cp(w, r)
	}

	// AWS SDK request
	if strings.Contains(userAgent, "aws-sdk") {
		if r.Method == "GET" {
			if r.URL.Query().Get("list-type") == "2" {
				awssdk.ListObjectsV2(w, r)
			} else {
				awssdk.Get(w, r)
			}
			return
		} else if r.Method == "PUT" {
			if r.ContentLength > 0 {
				awssdk.Put(w, r)
			} else {
				awssdk.Copy(w, r)
			}
			return
		} else if r.Method == "POST" && r.URL.Query().Get("delete") == "" {
			awssdk.Delete(w, r)
			return
		}
	}
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

	dst, err := os.Create(filepath.Join(UPLOAD_DIR, header.Filename))
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
	err := os.Remove(filepath.Join(UPLOAD_DIR, filename))
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
	err := os.Rename(filepath.Join(UPLOAD_DIR, oldFilename), filepath.Join(UPLOAD_DIR, newFilename))
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
	err := os.Mkdir(filepath.Join(UPLOAD_DIR, dirname), os.ModePerm)
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
	err := os.Remove(filepath.Join(UPLOAD_DIR, dirname))
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
