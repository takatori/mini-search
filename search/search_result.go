package search

import (
	"fmt"
	"sort"
)

// result is used to store a search result.
type result struct {
	docId int
	score float64
}

// String print search result info
func (r *result) String() string {
	return fmt.Sprintf("{docId: %v, score: %v}", r.docId, r.score)
}

type SearchResults []*result

func (results SearchResults) Sort() SearchResults {
	sort.Slice(results, func(i, j int) bool {
		return results[i].score > results[j].score
	})
	return results
}

func (results SearchResults) DocIds() []int {

	docIds := make([]int, len(results))
	for i, r := range results {
		docIds[i] = r.docId
	}
	return docIds
}

func (results SearchResults) AddResult(d int, score float64) SearchResults {
	return append(results, &result{
		d,
		score,
	})
}

