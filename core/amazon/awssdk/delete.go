package awssdk

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path"
	"strings"
)

type DeleteRequest struct {
	XMLName xml.Name `xml:"Delete"`
	Object  []Object `xml:"Object"`
}

type Object struct {
	Key string `xml:"Key"`
}

type DeleteResponse struct {
	XMLName xml.Name  `xml:"DeleteResult"`
	Deleted []Deleted `xml:"Deleted"`
}

type Deleted struct {
	Key string `xml:"Key"`
}

// Handle deletion request from AWS SDK
func Delete(w http.ResponseWriter, r *http.Request) {
	slog.Info("awssdk.Delete is called.")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var deleteRequest DeleteRequest
	err = xml.Unmarshal(body, &deleteRequest)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	dirPath := strings.TrimPrefix(r.URL.Path, "/")
	dirPath = path.Join("upload", dirPath)

	var deleteResponse DeleteResponse

	for _, object := range deleteRequest.Object {
		filePath := path.Join(dirPath, object.Key)
		err := os.Remove(filePath)
		if err != nil {
			slog.Info("Failed to delete", "filepath", filePath, "error", err)
			http.Error(w, fmt.Sprintf("Failed to delete file: %s", object.Key), http.StatusInternalServerError)
			return
		}
		slog.Info("File deleted. ", "filePath", filePath)
		deleteResponse.Deleted = append(deleteResponse.Deleted, Deleted(object))
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	xml.NewEncoder(w).Encode(deleteResponse)
}
