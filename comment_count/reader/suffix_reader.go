package reader

import (
	"path"
	"strings"
)

// NewSuffixFilter, includes used to appoint your suffixes(contains dot and ignore case), i.e. []string{".docx",".pptx"}
func NewSuffixFilter(includes []string) *SuffixFilter {
	filter := &SuffixFilter{includes: map[string]struct{}{}}
	for _, suffix := range includes {
		filter.includes[strings.ToLower(suffix)] = struct{}{}
	}
	return filter
}

// SuffixFilter a file filter by suffix, ignore case
type SuffixFilter struct {
	includes map[string]struct{}
}

func (f *SuffixFilter) IsFilter(name string) bool {
	suffix := strings.ToLower(path.Ext(name))
	if len(f.includes) > 0 {
		if _, ok := f.includes[suffix]; ok {
			return false
		}
	}
	return true
}
