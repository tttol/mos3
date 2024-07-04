package web

import (
	"net/http"
	"os"
	"path/filepath"
)

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

	http.Redirect(w, r, "/s3", http.StatusFound)
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
