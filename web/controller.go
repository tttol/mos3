package web

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/tttol/mos3/core/amazon/awscli"
	"github.com/tttol/mos3/core/amazon/awssdk"
	"github.com/tttol/mos3/core/logging"
)

func CliSdkHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("IndexHandler is called.")
	logging.LogRequest(r)

	// AWS CLI request
	userAgent := r.Header.Get("User-Agent")
	if strings.Contains(userAgent, "command/s3.ls") {
		awscli.Ls(w)
		return
	} else if strings.Contains(userAgent, "command/s3.cp") {
		awscli.Cp(w, r)
	}

	// AWS SDK request
	if strings.Contains(userAgent, "aws-sdk") {
		if r.Method == "GET" {
			if r.URL.Query().Get("list-type") == "2" {
				awssdk.ListObjectsV2(w, r)
			} else {
				awssdk.Get(w, r)
			}
			return
		} else if r.Method == "PUT" {
			if r.ContentLength > 0 {
				awssdk.Put(w, r)
			} else {
				awssdk.Copy(w, r)
			}
			return
		} else if r.Method == "POST" && r.URL.Query().Get("delete") == "" {
			awssdk.Delete(w, r)
			return
		}
	}
}
