package util

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tttol/mos3/core/model"
)

func GetS3Objects(r *http.Request, path string) ([]model.S3Object, error) {
	dirEntry, err := os.ReadDir(filepath.Join("./upload", path))
	if err != nil {
		slog.Error("ReadDir error", "error", err)
		return nil, err
	}

	var s3Objects []model.S3Object
	for _, entry := range dirEntry {
		var obj model.S3Object
		obj.Name = entry.Name()
		obj.FullPath = filepath.Join(r.URL.Path, entry.Name())
		obj.IsDir = entry.IsDir()
		slog.Info("S3Object", "data", obj)

		s3Objects = append(s3Objects, obj)
	}
	slog.Info("S3Objects", "data", s3Objects)

	return s3Objects, nil
}
