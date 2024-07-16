package awssdk

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

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
