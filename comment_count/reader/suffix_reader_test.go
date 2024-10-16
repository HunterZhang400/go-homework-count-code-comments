package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanFilter(t *testing.T) {
	filter := NewSuffixFilter([]string{".Cpp", ".hpp"})
	assert.Equal(t, true, filter.IsFilter("xxx.jpg"))
	assert.Equal(t, false, filter.IsFilter("xxx.cpp"))
	assert.Equal(t, false, filter.IsFilter("xxx.HPP"))
	assert.Equal(t, true, filter.IsFilter("xxx.c"))
	assert.Equal(t, true, filter.IsFilter("xxx"))
}
