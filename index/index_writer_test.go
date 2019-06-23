package index

import (
	"testing"
	"reflect"
)

func TestIndexWriter(t *testing.T) {

	collection := []string{
		"Do you quarrel, sir?",
		"Quarrel sir! no, sir!",
		"If you do, sir, I am for you: I serve as good a man as you.",
		"No better.",
		"Well, sir",
	}

	dictionary := map[string]*PostingsList{
		"a":       {[]*Posting{{3, []int{13}, 1}}},
		"am":      {[]*Posting{{3, []int{6}, 1}}},
		"as":      {[]*Posting{{3, []int{11, 15}, 2}}},
		"better":  {[]*Posting{{4, []int{2}, 1}}},
		"do":      {[]*Posting{{1, []int{1}, 1}, {3, []int{3}, 1}}},
		"for":     {[]*Posting{{3, []int{7}, 1}}},
		"good":    {[]*Posting{{3, []int{12}, 1}}},
		"i":       {[]*Posting{{3, []int{5, 9}, 2}}},
		"if":      {[]*Posting{{3, []int{1}, 1}}},
		"man":     {[]*Posting{{3, []int{14}, 1}}},
		"no":      {[]*Posting{{2, []int{3}, 1}, {4, []int{1}, 1}}},
		"quarrel": {[]*Posting{{1, []int{3}, 1}, {2, []int{1}, 1}}},
		"serve":   {[]*Posting{{3, []int{10}, 1}}},
		"sir":     {[]*Posting{{1, []int{4}, 1}, {2, []int{2, 4}, 2}, {3, []int{4}, 1}, {5, []int{2}, 1}}},
		"well":    {[]*Posting{{5, []int{1}, 1}}},
		"you":     {[]*Posting{{1, []int{2}, 1}, {3, []int{2, 8, 16}, 3}}},
	}

	expected := &Index{
		dictionary: dictionary,
		docCount:   5,
	}

	writer := NewIndexWriter()

	for _, c := range collection {
		writer.AddDocument(c)
	}

	actual := writer.Commit()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got:\n%v\n ================\nwant:\n%v\n", actual, expected)
	}
}
