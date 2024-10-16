package counter

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCounterV1(t *testing.T) {
	defaultCounter := NewCounter()
	defaultCounter.SetDebug(true)
	result, err := defaultCounter.Count("../../testing/cpp/special_cases.cpp")
	if err != nil {
		log.Fatalf("counterV2.Count err:%s", err.Error())
		return
	}
	assert.Equal(t, 62, result.Total, "wrong Total line")
	assert.Equal(t, 6, result.Inline, "wrong Inline comment line")
	assert.Equal(t, 34, result.Block, "wrong Block comment line")
}

func TestCounterNoMatcher(t *testing.T) {
	counter := DefaultCounter{}
	_, err := counter.Count("../../testing/cpp/special_cases.cpp")
	assert.Equal(t, "no machers", err.Error())
}
