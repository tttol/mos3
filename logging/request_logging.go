package logging

import (
	"log/slog"
	"net/http"
)

func LogRequestHeaders(r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			slog.Info("%v: %v", name, h)
		}
	}
}
