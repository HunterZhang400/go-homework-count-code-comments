package matcher

const codeRStringStartPrefix = byte('R')
const codeRStringStartSuffix = byte('"')
const codeRStringEnd = byte('"')

// CodeRStringMatcher  R"" in C++  matcher
type CodeRStringMatcher struct {
}

func (e CodeRStringMatcher) MatchStart(line []byte, from int) int {
	for i := from; i < len(line); i++ {
		if line[i] == codeRStringStartPrefix && i < len(line)-1 && line[i+1] == codeRStringStartSuffix {
			return i
		}
	}
	return -1
}

func (e CodeRStringMatcher) GetStartSeparatorLength() int {
	return 2
}

func (e CodeRStringMatcher) MatchEnd(line []byte, from int) (int, bool) {
	for i := from; i < len(line); i++ {
		if line[i] == codeRStringEnd {
			return i, true
		}
	}
	return len(line) - 1, false
}
