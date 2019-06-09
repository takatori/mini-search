package search

import (
	"github.com/takatori/mini-search/index"
)

// function to locate the first occurrence of a phrase after a given position
func NextPhrase(idx *index.Index, phrase []string, current *index.Position) (*index.Position, *index.Position) {

	v := current

	for _, t := range phrase {
		v = idx.Next(t, v)
	}
	if v == index.EOF {
		return index.EOF, index.EOF
	}
	u := v

	for i := len(phrase) - 2; i >= 0; i-- {
		u = idx.Prev(phrase[i], u)
	}

	if v.Distance(u) == len(phrase)-1 {
		return u, v
	}

	return NextPhrase(idx, phrase, u)

}

func PhraseSearch(idx *index.Index, phrase []string) [][2]*index.Position {

	var results [][2]*index.Position
	var u, v *index.Position

	u = index.BOF

	for index.ComparePosition(u, index.EOF) < 0 {
		u, v = NextPhrase(idx, phrase, u)
		if u != index.EOF {
			results = append(results, [2]*index.Position{u, v})
		}
	}
	return results
}
