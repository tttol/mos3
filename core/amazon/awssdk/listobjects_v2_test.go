package awssdk

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestListObjects(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := "test-upload"
	err := os.Mkdir(tempDir, 0777)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files and directories
	err = os.MkdirAll(filepath.Join(tempDir, "dir1"), os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(tempDir, "file1.txt"), []byte("file1 content"), os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(tempDir, "dir1", "file2.txt"), []byte("file2 content"), os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	// Call the ListObjects function
	items, err := ListObjects(filepath.Join(tempDir), "test-upload")
	if err != nil {
		t.Fatalf("ListObjects returned an error: %v", err)
	}

	// Define expected results
	expectedItems := []Item{
		{Key: "dir1/file2.txt", Size: int64(len("file2 content"))},
		{Key: "file1.txt", Size: int64(len("file1 content"))},
	}

	// Check the number of items
	if len(items) != len(expectedItems) {
		t.Fatalf("Expected %d items, got %d", len(expectedItems), len(items))
	}

	// Check each item
	for i, item := range items {
		if item.Key != expectedItems[i].Key || item.Size != expectedItems[i].Size {
			t.Errorf("Expected item %v, got %v", expectedItems[i], item)
		}
	}
}

func TestIsTruncated(t *testing.T) {
	items1 := make([]Item, 1)
	items999 := make([]Item, 999)
	items1000 := make([]Item, 1000)
	items1001 := make([]Item, 1001)
	items1200 := make([]Item, 1200)

	isTruncated1, actualItems1 := IsTruncated(items1)
	assert.Equal(t, isTruncated1, false)
	assert.Equal(t, actualItems1, items1)

	isTruncated999, actualItems999 := IsTruncated(items999)
	assert.Equal(t, isTruncated999, false)
	assert.Equal(t, actualItems999, items999)

	isTruncated1000, actualItems1000 := IsTruncated(items1000)
	assert.Equal(t, isTruncated1000, false)
	assert.Equal(t, actualItems1000, items1000)

	isTruncated1001, actualItems1001 := IsTruncated(items1001)
	assert.Equal(t, isTruncated1001, true)
	assert.Equal(t, actualItems1001, items1001[:1000])

	isTruncated1200, actualItems1200 := IsTruncated(items1200)
	assert.Equal(t, isTruncated1200, true)
	assert.Equal(t, actualItems1200, items1200[:1000])

}
