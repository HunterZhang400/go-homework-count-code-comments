// provider  the implements of count comment
package counter

// Counter define the counter specification
type Counter interface {
	//Count the total lines, inline comment and block comment separately
	Count(file string) (*Result, error)
}

type State uint16

const InlineState = State(1)
const BlockState = State(2)
const CodeState = State(3)
const CodeStringState = State(4)
const CodeRStringState = State(5)
