package util

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"github.com/tttol/mos3/core/model"
)

func GenerateS3Objects(r *http.Request, dir string, dirPath string) ([]model.S3Object, error) {
	dirEntry, err := os.ReadDir(filepath.Join(dir, dirPath))
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

		s3Objects = append(s3Objects, obj)
	}
	sorted := sortObjects(s3Objects)

	return sorted, nil
}

func sortObjects(s3Objects []model.S3Object) []model.S3Object {
	sort.Slice(s3Objects, func(i, j int) bool {
		// sort by IsDir asc
		if s3Objects[i].IsDir != s3Objects[j].IsDir {
			return s3Objects[i].IsDir
		}
		// sort by Name asc
		return s3Objects[i].Name < s3Objects[j].Name
	})
	return s3Objects
}
