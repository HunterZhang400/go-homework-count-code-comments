package matcher

// Matcher every type of programme matcher
type Matcher interface {

	//match start and return end index, not match return -1
	MatchStart(line []byte, from int) int

	GetStartSeparatorLength() int

	// MatchEnd scan all line and try to match the state end and return end index, if not match return the last index scanned,
	// and the second bool incicate whether current state end (because string and inline comment can cross multiple lines).
	MatchEnd(line []byte, from int) (int, bool)
}
