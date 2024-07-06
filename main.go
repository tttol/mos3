package main

import (
	"log/slog"
	"net/http"

	"github.com/tttol/mos3/web"
)

func main() {
	http.HandleFunc("/", web.CliSdkHandler)
	http.HandleFunc("/s3/", web.S3Handler)
	http.HandleFunc("/uploadpage", web.UploadIndexHandler)
	http.HandleFunc("/upload", web.UploadHandler)
	http.HandleFunc("/remove", web.RemoveHandler)
	http.HandleFunc("/rename", web.RenameHandler)
	http.HandleFunc("/mkdir", web.MkdirHandler)
	http.HandleFunc("/rmdir", web.RmdirHandler)
	http.HandleFunc("/renamedir", web.RenamedirHandler)

	slog.Info("Starting server at :3333")
	http.ListenAndServe(":3333", nil)
}
