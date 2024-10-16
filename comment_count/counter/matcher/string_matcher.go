package matcher

const lineBreak = byte('\\')
const stringStart = '"'
const stringEnd = '"'

// CodeStringMatcher  "" in C++  matcher
type CodeStringMatcher struct {
}

func (e CodeStringMatcher) MatchStart(line []byte, from int) int {
	for i := from; i < len(line); i++ {
		if line[i] == stringStart {
			return i
		}
	}
	return -1
}

func (e CodeStringMatcher) GetStartSeparatorLength() int {
	return 1
}

func (e CodeStringMatcher) MatchEnd(line []byte, from int) (int, bool) {
	for i := from; i < len(line); i++ {
		if line[i] == stringEnd {
			if i > 0 && line[i-1] == escape {
				continue
			}
			return i, true
		}
	}
	//not end
	if line[len(line)-1] == lineBreak {
		return len(line) - 1, false
	}
	//string  ended by new line
	return len(line) - 1, true
}
