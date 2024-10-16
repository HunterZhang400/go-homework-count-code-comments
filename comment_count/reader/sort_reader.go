package reader

import (
	"container/list"
	"errors"
	"os"
	"sort"
	"strings"
)

func NewSortReader() *SortReader {
	return &SortReader{}
}

// SortReader return all files ordered by absolute path alphabetically
type SortReader struct {
}

func (p *SortReader) Read(dir string, filter Filter) ([]string, error) {
	result := make([]string, 0, 100)
	waitDirs := list.New()
	waitDirs.PushBack(dir)
	for {
		dirInfo := waitDirs.Front()
		if dirInfo == nil {
			break
		}
		waitDirs.Remove(dirInfo)
		dirPath, ok := dirInfo.Value.(string)
		if !ok {
			return result, errors.New("waitDirs found value not string")
		}
		childResult, err := p.readChild(dirPath, filter, waitDirs)
		if err != nil {
			return result, err
		}
		result = append(result, childResult...)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result, nil
}

func (p *SortReader) readChild(dir string, filter Filter, waitDirs *list.List) ([]string, error) {
	result := make([]string, 0, 100)
	fs, err := os.ReadDir(dir)
	if err != nil {
		return result, err
	}
	parent := strings.TrimRight(strings.TrimRight(dir, "/"), "\\") + "/"
	for _, f := range fs {
		filePath := parent + f.Name()
		if f.IsDir() {
			waitDirs.PushBack(filePath)
			continue
		}
		if filter != nil && filter.IsFilter(filePath) {
			continue
		}
		result = append(result, filePath)
	}
	return result, nil
}
