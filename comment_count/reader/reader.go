package reader

// Reader read all files in dir recursively
type Reader interface {
	Read(dir string, filter Filter) ([]string, error)
}
