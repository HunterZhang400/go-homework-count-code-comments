package matcher

const inlineCommentStartPrefix = byte('/')
const inlineCommentStartSuffix = byte('/')
const escape = byte('\\')

// InlineCommentMatcher inline comment  matcher
type InlineCommentMatcher struct {
}

func (e InlineCommentMatcher) MatchStart(line []byte, from int) int {
	for i := from; i < len(line); i++ {
		if line[i] == inlineCommentStartPrefix && i < len(line)-1 && line[i+1] == inlineCommentStartSuffix {
			return i
		}
	}
	return -1
}

func (e InlineCommentMatcher) GetStartSeparatorLength() int {
	return 2
}

func (e InlineCommentMatcher) MatchEnd(line []byte, from int) (int, bool) {
	//not end
	if line[len(line)-1] == lineBreak {
		return len(line) - 1, false
	}
	//inline comment ended by new line
	return len(line) - 1, true
}
