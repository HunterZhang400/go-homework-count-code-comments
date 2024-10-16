package reader

import (
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

func TestSortReader(t *testing.T) {
	reader := NewSortReader()
	sortedList, err := reader.Read("../../testing", NewSuffixFilter([]string{".c", ".cpp", ".h", ".hpp"}))
	if err != nil {
		t.Error(err.Error())
		return
	}
	set := make(map[string]bool, len(sortedList))
	for _, file := range sortedList {
		set[file] = true
	}
	assert.NotContains(t, sortedList, "../../testing/cpp/lib_json/CMakeLists.txt")
	assert.Equal(t, "../../testing/cpp/lib_json/json_reader.cpp", sortedList[0])
	assert.Equal(t, "../../testing/cpp/special_cases.cpp", path.Join(sortedList[4]))
}
