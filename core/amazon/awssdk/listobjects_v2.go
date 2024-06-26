package awssdk

import (
	"encoding/xml"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ListBucketResult struct {
	XMLName xml.Name `xml:"ListBucketResult"`
	Name    string   `xml:"Name"`
	Prefix  string   `xml:"Prefix"`
	Marker  string   `xml:"Marker"`
	Items   []Item   `xml:"Contents"`
}

type Item struct {
	Key  string `xml:"Key"`
	Size int64  `xml:"Size"`
}

func ListObjectsV2(w http.ResponseWriter, r *http.Request) {
	slog.Info("ListObjectsV2 is called.")

	path := strings.Split(r.URL.Path, "?list-type=2")[0] // It has been confirmed in the previous process controller.go that `?list-type=2` is included.
	dir := strings.TrimPrefix(path, "/")
	if dir == "" {
		slog.Error("No directory specified in the query parameter")
		http.Error(w, "No directory specified", http.StatusBadRequest)
		return
	}

	// Define the target directory
	rootDir := filepath.Join("upload", dir)

	var items []Item
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			items = append(items, Item{
				Key:  filepath.ToSlash(path[len("upload/"):]), // Convert to a relative path
				Size: info.Size(),
			})
		}
		return nil
	})
	if err != nil {
		slog.Error("Failed to list objects", "error", err)
		http.Error(w, "Failed to list objects", http.StatusInternalServerError)
		return
	}

	slog.Info("Files are below", "files", items)

	response := ListBucketResult{
		Name:  dir,
		Items: items,
	}

	w.Header().Set("Content-Type", "application/xml")
	if err := xml.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
