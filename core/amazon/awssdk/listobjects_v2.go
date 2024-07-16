package awssdk

import (
	"encoding/xml"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ListObjectsResult struct {
	XMLName     xml.Name `xml:"ListBucketResult"`
	Name        string   `xml:"Name"`
	Prefix      string   `xml:"Prefix"`
	Marker      string   `xml:"Marker"`
	Items       []Item   `xml:"Contents"`
	IsTruncated bool     `xml:"IsTruncated"`
}

type Item struct {
	Key  string `xml:"Key"`
	Size int64  `xml:"Size"`
}

func ListObjectsV2(w http.ResponseWriter, r *http.Request, uploadDirName string) {
	slog.Info("ListObjectsV2 is called.")

	path := strings.Split(r.URL.Path, "?list-type=2")[0] // It has been confirmed in the previous process controller.go that `?list-type=2` is included.
	dir := strings.TrimPrefix(path, "/")
	if dir == "" {
		slog.Error("No directory specified in the query parameter")
		http.Error(w, "No directory specified", http.StatusBadRequest)
		return
	}

	rootDir := filepath.Join(uploadDirName, dir)

	items, err := ListObjects(rootDir, uploadDirName)
	if err != nil {
		slog.Error("Failed to list objects", "error", err)
		http.Error(w, "Failed to list objects", http.StatusInternalServerError)
		return
	}

	slog.Info("Files are below", "files", items)

	isTruncated, items := IsTruncated(items)

	response := ListObjectsResult{
		Name:        dir,
		Items:       items,
		IsTruncated: isTruncated,
	}

	w.Header().Set("Content-Type", "application/xml")
	if err := xml.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func ListObjects(rootDir string, uploadDirName string) ([]Item, error) {
	var items []Item
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			items = append(items, Item{
				Key:  filepath.ToSlash(path[len(uploadDirName+"/"):]),
				Size: info.Size(),
			})
		}
		return nil
	})

	return items, err
}

func IsTruncated(items []Item) (bool, []Item) {
	if len(items) > 1000 {
		return true, items[:1000]
	} else {
		return false, items
	}
}
