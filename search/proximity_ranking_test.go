package search

import (
	"testing"
	"github.com/takatori/mini-search/index"
	"reflect"
)

func TestProxymityRanking(t *testing.T) {


	collection := []string{
		"Do you quarrel, sir?",
		"Quarrel sir! no, sir!",
		"If you do, sir, I am for you: I serve as good a man as you.",
		"No better.",
		"Well, sir",
	}

	writer := index.NewIndexWriter()

	for _, c := range collection {
		writer.AddDocument(c)
	}

	idx := writer.Commit()

	terms := []string{"quarrel", "sir"}

	actual := RankProximity(idx, terms, 10)

	expected := []int{3,1}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got: %v, want: %v", actual, expected)
	}
}