package awscli

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func Cp(w http.ResponseWriter, r *http.Request) {
	slog.Info("Executing `aws s3 cp`")
	filePath := fmt.Sprintf("upload%s", r.URL.Path) // Relative path from the directory where main.go is executed
	slog.Info("Checking file path", "filePath", filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// HTTP 404
		slog.Error("File not found.", "filePath", filePath)
		http.Error(w, fmt.Sprintf("File %s not found", filePath), http.StatusNotFound)
		return
	}

	// TODO HTTP method check

	// if r.Method != http.MethodHead {
	// 	// HTTP 405
	// 	slog.Error("Method not allowed", "method", r.Method)
	// 	http.Error(w, fmt.Sprintf("Method %s not allowed. Only HEAD is allowed. ", r.Method), http.StatusMethodNotAllowed)
	// 	return
	// }

	// download file
	slog.Info("Serving file", "filePath", filePath)
	http.ServeFile(w, r, filePath)
}
