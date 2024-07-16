package web

import (
	"net/http"
	"os"
	"path/filepath"
)

func RenameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	oldFilename := r.FormValue("oldFilename")
	newFilename := r.FormValue("newFilename")
	err := os.Rename(filepath.Join(UPLOAD_DIR_PATH, oldFilename), filepath.Join(UPLOAD_DIR_PATH, newFilename))
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
	err := os.Rename(filepath.Join(UPLOAD_DIR_PATH, oldDirname), filepath.Join(UPLOAD_DIR_PATH, newDirname))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
