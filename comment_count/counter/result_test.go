package counter

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestResultOutput(t *testing.T) {
	resultChan := make(chan Result, 20)
	endSingal := false
	//simulate add result
	go func() {
		for i := 0; i < 50; i++ {
			total := rand.Intn(1000) + 100
			inline := total - rand.Intn(total)
			block := (total - inline) / 8
			resultChan <- Result{
				FilePath: fmt.Sprintf("testing/cpp%s/%d.cpp", createDeepFolder(rand.Intn(4)), i*total),
				Total:    total,
				Inline:   inline,
				Block:    block,
			}
			time.Sleep(time.Millisecond * time.Duration(20))
		}
		endSingal = true
	}()
	// print out
	for endSingal == false || len(resultChan) > 0 {
		select {
		case result := <-resultChan:
			fmt.Println(result.ToPrintString())
		default:
			time.Sleep(time.Second)
		}
	}
}

func createDeepFolder(n int) string {
	buf := bytes.Buffer{}
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("/children_%d", i))
	}
	return buf.String()
}

func TestHideLongPath(t *testing.T) {
	assert.Equal(t, "testing/2/1334.cpp", hideMiddlePath("testing/2/1334.cpp", 2))
	assert.Equal(t, "testing/2/1334.cpp", hideMiddlePath("testing/2/1334.cpp", 20))
	assert.Equal(t, "/folder1/fold...folder4/aaaa.cpp", hideMiddlePath("/folder1/folder2/folder3/folder4/aaaa.cpp", 32))
}

func TestResultFormat(t *testing.T) {
	result := Result{
		FilePath: "aaa",
		Total:    200,
		Inline:   50,
		Block:    10,
	}
	assert.Equal(t, "aaa                                           total:  200    inline:   50    block:   10", result.ToPrintString())
	fmt.Println(result.ToPrintString())
}
