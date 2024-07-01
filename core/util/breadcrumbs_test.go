package util

import (
	"reflect"
	"testing"
)

func TestGenerateBreadcrumbs(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]interface{}
	}{
		{
			input: "/fuga/fugafuga/aaa.txt",
			expected: map[string]interface{}{
				"fuga":     "/s3/fuga",
				"fugafuga": "/s3/fuga/fugafuga",
				"aaa.txt":  "/s3/fuga/fugafuga/aaa.txt?action=dl",
			},
		},
		{
			input: "/hoge/fuga/piyo",
			expected: map[string]interface{}{
				"hoge": "/s3/hoge",
				"fuga": "/s3/hoge/fuga",
				"piyo": "/s3/hoge/fuga/piyo",
			},
		},
		{
			input: "/s3/hoge",
			expected: map[string]interface{}{
				"s3":   "/s3/s3",
				"hoge": "/s3/s3/hoge",
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
