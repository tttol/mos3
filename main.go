package main

import (
	"log/slog"
	"net/http"

	"github.com/tttol/mos3/web"
)

func main() {
	http.HandleFunc("/", web.IndexHandler)
	http.HandleFunc("/s3/", web.S3Handler)
	http.HandleFunc("/upload", web.UploadHandler)
	http.HandleFunc("/delete", web.DeleteHandler)
	http.HandleFunc("/rename", web.RenameHandler)
	http.HandleFunc("/mkdir", web.MkdirHandler)
	http.HandleFunc("/rmdir", web.RmdirHandler)
	http.HandleFunc("/renamedir", web.RenamedirHandler)

	slog.Info("Starting server at :3333")
	http.ListenAndServe(":3333", nil)
}
