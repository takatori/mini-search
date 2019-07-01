package search

import (
	"github.com/takatori/mini-search/index"
)

func nextCover(idx *index.Index, terms []string, p *index.Position) (*index.Position, *index.Position) {

	v := index.BOF
	u := index.EOF

	for _, term := range terms {
		p := idx.Next(term, p)
		if index.ComparePosition(p, v) > 0 {
			v = p
		}
	}

	if v == index.EOF {
		return index.EOF, index.EOF
	}

	for _, term := range terms {
		p := idx.Prev(term, index.NewPosition(index.DocId(v), index.Offset(v)+1))
		if index.ComparePosition(u, p) > 0 {
			u = p
		}
	}

	if index.DocId(u) == index.DocId(v) {
		return u, v
	} else {
		return nextCover(idx, terms, u)
	}

}

func RankProximity(idx *index.Index, terms []string, k int) []int {

	results := make(SearchResults, 0, k)
	v, u := nextCover(idx, terms, index.BOF)
	d := index.DocId(u)
	score := 0.0

	for index.ComparePosition(index.EOF, u) > 0 {
		if d < index.DocId(u) {
			results = results.AddResult(d, score)
			d = index.DocId(u)
			score = 0
		}
		score = score + 1/(float64(v.Distance(u))+1)
		u, v = nextCover(idx, terms, u)
	}
	if d < index.EndOfFile {
		results = results.AddResult(d, score)
	}

	return results.Sort().DocIds()
}
