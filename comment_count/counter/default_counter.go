package counter

import (
	"compass.com/go-homework/comment_count/counter/matcher"
	"errors"
	"fmt"
	"os"
	"strings"
)

func NewCounter() *DefaultCounter {
	return &DefaultCounter{machers: map[State]matcher.Matcher{
		CodeStringState:  matcher.CodeStringMatcher{},
		CodeRStringState: matcher.CodeRStringMatcher{},
		InlineState:      matcher.InlineCommentMatcher{},
		BlockState:       matcher.BlockCommentMatcher{}}}
}

type DefaultCounter struct {
	machers     map[State]matcher.Matcher
	enableDebug bool
}

func (c *DefaultCounter) SetDebug(b bool) {
	c.enableDebug = b
}

func (c *DefaultCounter) Count(file string) (*Result, error) {
	if len(c.machers) == 0 {
		return nil, errors.New("no machers")
	}
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var (
		total, inlineCount, blockCount int
	)
	lines := strings.Split(string(content), "\n")
	lastIndex := len(lines) - 1
	currentState := CodeState
	for i, line := range lines {
		// start from non-blank character
		lineBytes := []byte(strings.TrimSpace(line))
		if len(lineBytes) == 0 {
			if c.enableDebug {
				c.printDebugInfo(&lineResult{
					EndState: currentState,
					StateSet: map[State]struct{}{currentState: {}},
				}, i, line)
			}
			// the requirement point 4, rows not include the last empty row
			if i != lastIndex {
				total++
				if currentState == InlineState {
					inlineCount++
				}
				if currentState == BlockState {
					blockCount++
				}
			}
			continue
		}
		total++
		result, err := c.processLine(currentState, lineBytes)
		if err != nil {
			return nil, err
		}
		currentState = result.EndState
		if _, ok := result.StateSet[InlineState]; ok {
			inlineCount++
		}
		if _, ok := result.StateSet[BlockState]; ok {
			blockCount++
		}
		if c.enableDebug {
			c.printDebugInfo(result, i, line)
		}
	}
	result := &Result{
		FilePath: file,
		Total:    total,
		Inline:   inlineCount,
		Block:    blockCount,
	}
	return result, nil
}

func (c *DefaultCounter) printDebugInfo(result *lineResult, i int, line string) {
	inlineTag := "     "
	blockTag := "     "
	if _, ok := result.StateSet[InlineState]; ok {
		inlineTag = "INLIN"
	}
	if _, ok := result.StateSet[BlockState]; ok {
		blockTag = "BLOCK"
	}
	fmt.Println(fmt.Sprintf("%-5d [%s,%s] %s", i+1, inlineTag, blockTag, line))
}

func (c *DefaultCounter) processLine(currentState State, lineBytes []byte) (*lineResult, error) {
	var (
		lastIndex = len(lineBytes) - 1
		idx       = 0
		isEnd     bool
		result    = &lineResult{
			EndState: currentState,
			StateSet: make(map[State]struct{}),
		}
	)
	result.StateSet[currentState] = struct{}{}
	for idx < lastIndex {
		switch result.EndState {

		case CodeState: //check is exist other state
			//because R string can contain inline and block comment so need to check
			currentState, idx = c.findFirstState(lineBytes, idx)
			result.EndState = currentState
			result.StateSet[currentState] = struct{}{}
			idx++

		case CodeStringState, CodeRStringState, InlineState, BlockState: //find current state end index and record
			matcher := c.getMatcher(result.EndState)
			if matcher == nil {
				return nil, errors.New(fmt.Sprintf("processLine invalid state:%v", result.EndState))
			}
			idx, isEnd = matcher.MatchEnd(lineBytes, idx)
			//if previous state not continues to the last character, reset to programme code state.
			//So, the remains can be scan next loop
			if idx < lastIndex {
				result.EndState = CodeState
			} else if idx == lastIndex {
				if isEnd {
					result.EndState = CodeState
				} else {
					//not end, keep to use current state
				}
			}
			idx++
		}
	}
	return result, nil
}

func (c *DefaultCounter) getMatcher(key State) matcher.Matcher {
	return c.machers[key]
}

// findFirstState from programme code to find the first other State and start index. (inline comment, code string/code R string, block comment etc.)
func (c *DefaultCounter) findFirstState(line []byte, from int) (State, int) {
	var (
		state          = CodeState
		index          = len(line) - 1
		nearestMatcher matcher.Matcher
	)
	for k, m := range c.machers {
		startIndex := m.MatchStart(line, from)
		if startIndex >= 0 && startIndex < index {
			state = k
			index = startIndex
			nearestMatcher = m
		}
	}
	if nearestMatcher != nil {
		index += nearestMatcher.GetStartSeparatorLength() - 1
	}
	return state, index
}

type lineResult struct {
	EndState State
	StateSet map[State]struct{}
}
