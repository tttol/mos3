package util

import (
	"regexp"
	"strings"
)

type Breadcrumb struct {
	Name string
	Path string
}

func GenerateBreadcrumbs(path string) []Breadcrumb {
	var breadcrumbs []Breadcrumb
	splitted := strings.Split(path, "/")
	fullPath := "/s3"
	for i, s := range splitted {
		if s == "" {
			continue
		}
		fullPath += "/" + s

		r, _ := regexp.Compile(`.*\..*`)
		if i == len(splitted)-1 && r.Match([]byte(s)) {
			breadcrumbs = append(breadcrumbs, Breadcrumb{Name: s, Path: fullPath + "?action=dl"})
		} else {
			breadcrumbs = append(breadcrumbs, Breadcrumb{Name: s, Path: fullPath})
		}
	}
	return breadcrumbs
}
