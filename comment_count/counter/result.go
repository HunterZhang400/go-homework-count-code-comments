package counter

import "fmt"

const pathWidth = 42
const lineNumberWidth = 5
var printFormat = fmt.Sprintf("%%-%ds    total:%%%dd    inline:%%%dd    block:%%%dd", pathWidth, lineNumberWidth,
	lineNumberWidth, lineNumberWidth)

type Result struct {
	FilePath string
	Total    int
	Inline   int
	Block    int
}

func (r *Result) ToPrintString() string {
	return fmt.Sprintf(printFormat, hideMiddlePath(r.FilePath, pathWidth), r.Total, r.Inline, r.Block)
}

// hideMiddlePath hide the middle part of a path, maxWidth should greater than 10 and this function
// does not consider multi-language characters. For example:
//
// "/folder1/folder2/folder3/folder4/aaaa.cpp"
//
// Converted to a max width of 32 would be:
//
// "/folder1/fold...folder4/aaaa.cpp"
func hideMiddlePath(path string, maxWidth int) string {
	if maxWidth <= 10 {
		return path
	}
	if len(path) > maxWidth {
		prefixSize := maxWidth/2 - 3
		return path[:prefixSize] + "..." + path[len(path)-maxWidth/2:]
	}
	return path
}
