package web

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetDirPath(t *testing.T) {
	assert.Equal(t, GetDirPath("/hoge/fuga"), "/hoge/fuga")
	assert.Equal(t, GetDirPath("/hoge/fuga/piyo.txt"), "/hoge/fuga")
}
