package s3

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func Cp(w http.ResponseWriter, r *http.Request) {
	slog.Info("Executing `aws s3 cp`")
	// filePath := r.URL.Path
	filePath := "../../upload/hoge/aaa.txt"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// HTTP 404
		slog.Error("File not found.", "filePath", filePath)
		http.Error(w, fmt.Sprintf("File %s not found", filePath), http.StatusNotFound)
		return
	}

	if r.Method != http.MethodHead {
		// HTTP 405
		slog.Error("Method not allowed", "method", r.Method)
		http.Error(w, fmt.Sprintf("Method %s not allowed. Only HEAD is allowed. ", r.Method), http.StatusMethodNotAllowed)
	}

	// download file
	http.ServeFile(w, r, filePath)
}
