package web

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"testing"
// )

// func TestRemoveHandler(t *testing.T) {
// 	// TEST_UPLOAD_DIR := "/test"

// 	tempDir, err := os.MkdirTemp("", "test-upload-dir")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer os.RemoveAll(tempDir)

// 	// Override TEST_UPLOAD_DIR to use the temporary directory
// 	// uploadDir := TEST_UPLOAD_DIR
// 	// defer func() { TEST_UPLOAD_DIR = uploadDir }()
// 	// TEST_UPLOAD_DIR = tempDir

// 	// Create a test file to be removed
// 	testFilePath := filepath.Join(UPLOAD_DIR, "testfile.txt")
// 	dst, err := os.Create(testFilePath)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer dst.Close()

// 	err = os.WriteFile(testFilePath, []byte("test content"), 0644)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Prepare the request
// 	formData := strings.NewReader("path=/s3/testfile.txt")
// 	req, err := http.NewRequest("POST", "/s3/remove", formData)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// Record the response
// 	responseRecorder := httptest.NewRecorder()
// 	handler := http.HandlerFunc(RemoveHandler)

// 	// Call the handler
// 	handler.ServeHTTP(responseRecorder, req)

// 	// Check the status code
// 	if status := responseRecorder.Code; status != http.StatusFound {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusFound)
// 	}

// 	// Check the file is removed
// 	if _, err := os.Stat(testFilePath); !os.IsNotExist(err) {
// 		t.Errorf("expected file to be removed, but it still exists: %v", err)
// 	}
// }

// func TestRemoveHandler_InvalidMethod(t *testing.T) {
// 	// Prepare the request with invalid method
// 	req, err := http.NewRequest("GET", "/s3/remove", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Record the response
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(RemoveHandler)

// 	// Call the handler
// 	handler.ServeHTTP(rr, req)

// 	// Check the status code
// 	if status := rr.Code; status != http.StatusMethodNotAllowed {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusMethodNotAllowed)
// 	}
// }

// func TestRemoveHandler_FileNotFound(t *testing.T) {
// 	// Prepare the request
// 	formData := strings.NewReader("path=/s3/nonexistentfile.txt")
// 	req, err := http.NewRequest("POST", "/s3/remove", formData)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// Record the response
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(RemoveHandler)

// 	// Call the handler
// 	handler.ServeHTTP(rr, req)

// 	// Check the status code
// 	if status := rr.Code; status != http.StatusInternalServerError {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusInternalServerError)
// 	}
// }
