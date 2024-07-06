package util

import (
	"reflect"
	"testing"
)

func TestGenerateBreadcrumbs(t *testing.T) {
	tests := []struct {
		input    string
		expected []Breadcrumb
	}{
		{
			input: "/fuga/fugafuga/aaa.txt",
			expected: []Breadcrumb{
				{Name: "fuga", Path: "/s3/fuga"},
				{Name: "fugafuga", Path: "/s3/fuga/fugafuga"},
				{Name: "aaa.txt", Path: "/s3/fuga/fugafuga/aaa.txt?action=dl"},
			},
		},
		{
			input: "/hoge/fuga/piyo",
			expected: []Breadcrumb{
				{Name: "hoge", Path: "/s3/hoge"},
				{Name: "fuga", Path: "/s3/hoge/fuga"},
				{Name: "piyo", Path: "/s3/hoge/fuga/piyo"},
			},
		},
		{
			input: "/s3/hoge",
			expected: []Breadcrumb{
				{Name: "s3", Path: "/s3/s3"},
				{Name: "hoge", Path: "/s3/s3/hoge"},
			},
		},
		{
			input: "/files/2024/0706/1845/txt/1234.txt",
			expected: []Breadcrumb{
				{Name: "files", Path: "/s3/files"},
				{Name: "2024", Path: "/s3/files/2024"},
				{Name: "0706", Path: "/s3/files/2024/0706"},
				{Name: "1845", Path: "/s3/files/2024/0706/1845"},
				{Name: "txt", Path: "/s3/files/2024/0706/1845/txt"},
				{Name: "1234.txt", Path: "/s3/files/2024/0706/1845/txt/1234.txt?action=dl"},
			},
		},
	}

	for _, test := range tests {
		result := GenerateBreadcrumbs(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("GenerateBreadcrumbs(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
