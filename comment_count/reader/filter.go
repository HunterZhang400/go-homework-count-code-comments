package reader

// Filter a filter for Reader
type Filter interface {
	//IsFilter check whether the file  should be filtered.
	//Returning true means it should be filtered, and false means it should not be filtered
	IsFilter(entry string) bool
}
