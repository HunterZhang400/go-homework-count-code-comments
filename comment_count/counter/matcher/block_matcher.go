package matcher

const blockCommentStartPrefix = byte('/')
const blockCommentStartSuffix = byte('*')
const blockCommentEndPrefix = byte('*')
const blockCommentEndSuffix = byte('/')

// BlockCommentMatcher block comment  matcher
type BlockCommentMatcher struct {
}

func (e BlockCommentMatcher) MatchStart(line []byte, from int) int {
	for i := from; i < len(line); i++ {
		// match start
		if line[i] == blockCommentStartPrefix && i < len(line)-1 && line[i+1] == blockCommentStartSuffix {
			return i
		}
	}
	return -1
}

func (e BlockCommentMatcher) GetStartSeparatorLength() int {
	return 2
}

func (e BlockCommentMatcher) MatchEnd(line []byte, from int) (int, bool) {
	for i := from; i < len(line); i++ {
		// match end
		if line[i] == blockCommentEndPrefix && i < len(line)-1 && line[i+1] == blockCommentEndSuffix {
			return i + 1, true
		}
	}
	return len(line) - 1, false
}
