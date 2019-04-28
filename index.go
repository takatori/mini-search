package index

import (
	"math"
)

const (
	BEGINNING_OF_FILE = -math.MaxInt32
	END_OF_FILE       = math.MaxInt32
)

// schema independent index
type Index struct {
	Dictionary map[string][]int
}

// first(t) returns the first position at which the term  t occurs in the collection
func (index *Index) first(t string) int {

	if postingList, ok := index.Dictionary[t]; !ok {
		return 0
	} else {
		return postingList[0]
	}
}

// last(t) returns the last position at which t occurs in collection
func (index *Index) last(t string) int {

	if postingList, ok := index.Dictionary[t]; !ok {
		return 0
	} else {
		return postingList[len(postingList)-1]
	}
}

// next(t, current) returns the position of t's first occurrence after the current position
func (index *Index) next(t string, current int) int {

	if postingList, ok := index.Dictionary[t]; !ok {
		return 0
	} else {
		for _, p := range postingList {
			if p > current {
				return p
			}
		}
		return END_OF_FILE
	}
}

// prev(t, current) returns the position of t's last occurrence before the current position
func (index *Index) prev(t string, current int) int {

	if postingList, ok := index.Dictionary[t]; !ok {
		return 0
	} else {
		for i := len(postingList) - 1; i >= 0; i-- {
			p := postingList[i]
			if p < current {
				return p
			}
		}
		return BEGINNING_OF_FILE
	}

}

// function to locate the first occurrence of a phrase after a given position
func (index *Index) nextPhrase(phrase []string, current int) (int, int) {

	length := len(phrase)

	v := current
	for _, t := range phrase {
		v = index.next(t, v)
	}
	if v == END_OF_FILE {
		return END_OF_FILE, END_OF_FILE
	}
	u := v
	for i := length - 2; i >= 0; i-- {
		u = index.prev(phrase[i], u)
	}
	if v-u == length-1 {
		return u, v
	} else {
		return index.nextPhrase(phrase, u)
	}
}

func NewIndex(dictionary map[string][]int) *Index {
	index := new(Index)
	index.Dictionary = dictionary
	return index
}
