package util

import (
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/tttol/mos3/core/model"
)

func setupTestDir(t *testing.T) {
	err := os.MkdirAll("./testdir/hoge/", os.ModePerm)
	if err != nil {
		t.Fatalf("setupTestDir error: %v", err)
	}
	_, err = os.Create("./testdir/hoge/file1.txt")
	if err != nil {
		t.Fatalf("setupTestDir error: %v", err)
	}
	_, err = os.Create("./testdir/hoge/file2.txt")
	if err != nil {
		t.Fatalf("setupTestDir error: %v", err)
	}
	err = os.Mkdir("./testdir/hoge/dir1", os.ModePerm)
	if err != nil {
		t.Fatalf("setupTestDir error: %v", err)
	}
}

func removeTestDir() {
	os.RemoveAll("./testdir")
}

func TestGenerateS3Objects(t *testing.T) {
	setupTestDir(t)
	defer removeTestDir()

	req := httptest.NewRequest("GET", "/hoge", nil)

	expected := []model.S3Object{
		{
			Name:     "dir1",
			FullPath: "/hoge/dir1",
			IsDir:    true,
		},
		{
			Name:     "file1.txt",
			FullPath: "/hoge/file1.txt",
			IsDir:    false,
		},
		{
			Name:     "file2.txt",
			FullPath: "/hoge/file2.txt",
			IsDir:    false,
		},
	}

	actual, err := GenerateS3Objects(req, "./testdir", "hoge")
	if err != nil {
		t.Fatalf("GenerateS3Objects error: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestSortObjects(t *testing.T) {
	input := []model.S3Object{
		{
			Name:  "ccc.txt",
			IsDir: false,
		},
		{
			Name:  "bbb.txt",
			IsDir: false,
		},
		{
			Name:  "aaa.txt",
			IsDir: false,
		},
		{
			Name:  "dir2",
			IsDir: true,
		},
		{
			Name:  "dir1",
			IsDir: true,
		},
	}

	expected := []model.S3Object{
		{
			Name:  "dir1",
			IsDir: true,
		},
		{
			Name:  "dir2",
			IsDir: true,
		},
		{
			Name:  "aaa.txt",
			IsDir: false,
		},
		{
			Name:  "bbb.txt",
			IsDir: false,
		},
		{
			Name:  "ccc.txt",
			IsDir: false,
		},
	}

	actual := sortObjects(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
