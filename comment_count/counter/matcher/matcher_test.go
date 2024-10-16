package matcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringMatcher(t *testing.T) {
	matcher := CodeStringMatcher{}
	index := matcher.MatchStart([]byte("x := \"xxx\";"), 0)
	assert.Equal(t, 5, index)
	endIndex, isEnd := matcher.MatchEnd([]byte("x := \"xxx\";"), index+matcher.GetStartSeparatorLength())
	assert.Equal(t, 9, endIndex)
	assert.Equal(t, true, isEnd)

	endIndex, isEnd = matcher.MatchEnd([]byte("x := \"x\\\""), index+1)
	assert.Equal(t, 8, endIndex, "CodeStringMatcher have escape ")
	assert.Equal(t, true, isEnd, "CodeStringMatcher have escape ")

	endIndex, isEnd = matcher.MatchEnd([]byte("x := \"x\\"), index+1)
	assert.Equal(t, 7, endIndex, "CodeStringMatcher have line break ")
	index = matcher.MatchStart([]byte("x :=//;"), 0)
	assert.Equal(t, -1, index, "not found")
	endIndex, isEnd = matcher.MatchEnd([]byte("x := 9x\\"), index+1)
	assert.Equal(t, 7, endIndex, " not found end")
}

func TestInlineMatcherHaveEscape(t *testing.T) {
	matcher := CodeStringMatcher{}
	index, end := matcher.MatchEnd([]byte(`xyz\"999"`), 0)
	assert.Equal(t, 8, index)
	assert.Equal(t, true, end)
}

func TestInlineComment(t *testing.T) {
	matcher := InlineCommentMatcher{}
	index := matcher.MatchStart([]byte(`xyz//999"`), 0)
	assert.Equal(t, 3, index)
	index, end := matcher.MatchEnd([]byte(`xyz//999`), index+matcher.GetStartSeparatorLength())
	assert.Equal(t, 7, index)
	index, end = matcher.MatchEnd([]byte("xyz//999\"\\\n000\\"), index+matcher.GetStartSeparatorLength())
	assert.Equal(t, 14, index)
	assert.Equal(t, end, false)
	index = matcher.MatchStart([]byte(`xyz999"`), 0)
	assert.Equal(t, -1, index)
	endIndex, _ := matcher.MatchEnd([]byte("x := 9x\\"), index+1)
	assert.Equal(t, 7, endIndex, " not found end")
}

func TestBlockComment(t *testing.T) {
	matcher := BlockCommentMatcher{}
	index := matcher.MatchStart([]byte(`xyz//9**9/*123456*/"`), 0)
	assert.Equal(t, 9, index)
	endIndex, isEnd := matcher.MatchEnd([]byte(`xyz//9**9/*123456*/"`), index+matcher.GetStartSeparatorLength())
	assert.Equal(t, 18, endIndex)
	assert.Equal(t, true, isEnd)
	index = matcher.MatchStart([]byte(`xyz//9**9123456*/"`), 0)
	assert.Equal(t, -1, index)
	endIndex, _ = matcher.MatchEnd([]byte("x := 9x\\"), index+1)
	assert.Equal(t, 7, endIndex, " not found end")
}

func TestRStringComment(t *testing.T) {
	matcher := CodeRStringMatcher{}
	index := matcher.MatchStart([]byte(`xyzR"123\n"`), 0)
	assert.Equal(t, 3, index)
	endIndex, isEnd := matcher.MatchEnd([]byte(`xyzR"123"`), index+matcher.GetStartSeparatorLength())
	assert.Equal(t, 8, endIndex)
	assert.Equal(t, true, isEnd)
	index = matcher.MatchStart([]byte(`xRyz"123R\n"`), 0)
	assert.Equal(t, -1, index, "not found")
	endIndex, _ = matcher.MatchEnd([]byte("x := 9x\\"), index+1)
	assert.Equal(t, 7, endIndex, " not found end")
}
