package logging

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

func LogRequest(r *http.Request) {
	slog.Info("[Request URL&method]", "url", r.URL.String(), "method", r.Method)
	for name, headers := range r.Header {
		for _, h := range headers {
			if strings.EqualFold(name, "Authorization") {
				// skip Authorization header
				continue
			}
			slog.Info("[Request Header]", name, h)
		}
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("[Request Body] Error reading body", "error", err)
		return
	}
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	slog.Info("[Request Body]" + string(bodyBytes))
}
