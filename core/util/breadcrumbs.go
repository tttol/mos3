package util

import (
	"regexp"
	"strings"
)

func GenerateBreadcrumbs(path string) map[string]interface{} {
	breadcrumbs := make(map[string]interface{})
	parts := strings.Split(path, "/")
	fullPath := ""
	for i, part := range parts {
		if part == "" {
			continue
		}
		fullPath += "/" + part

		r, _ := regexp.Compile(`.*\..*`)
		if i == len(parts)-1 && r.Match([]byte(part)) {
			breadcrumbs[part] = fullPath + "?action=dl"
		} else {
			breadcrumbs[part] = fullPath
		}
	}
	return breadcrumbs
}
