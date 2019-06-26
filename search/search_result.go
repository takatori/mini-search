package search

import "fmt"

// result is used to store a search result.
type result struct {
	docId int
	score float64
}

// String print search result info
func (r *result) String() string {
	return fmt.Sprintf("{docId: %v, score: %v}", r.docId, r.score)
}
