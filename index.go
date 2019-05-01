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
	cache map[string]int
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

	var p []int
	var ok bool
	var low, high, jump int

	if p, ok = index.Dictionary[t]; !ok {
		return END_OF_FILE
	}

	length := len(p)

	if length == 0 || p[length-1] <= current {
		return END_OF_FILE
	}

	if p[0] > current {
		index.cache[t] = 0
		return p[index.cache[t]]
	}

	if index.cache[t] > 0 && p[index.cache[t] - 1] <= current {
		low = index.cache[t] - 1
	} else {
		low = 0
	}

	jump = 1
	high = low + jump

	for high < length && p[high] <= current {
		low = high
		jump = 2 * jump
		high = low + jump
	}

	if high > length {
		high = length
	}

	index.cache[t] = index.binarySearch(t, low, high, current)

	return p[index.cache[t]]

}

func (index *Index) binarySearch(t string, low, high, current int) int {

	for high-low > 1 {
		mid := (low + high) / 2
		if p, ok := index.Dictionary[t]; !ok {
			return END_OF_FILE
		} else {
			if p[mid] <= current {
				low = mid
			} else {
				high = mid
			}
		}
	}
	return high
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

func (index *Index) allPhrase(phrase []string, current int) [][2]int {

	var results [][2]int
	var u int
	var v int

	u = current

	for u < END_OF_FILE {
		u, v = index.nextPhrase(phrase, u)
		if u != END_OF_FILE {
			results = append(results, [2]int{u, v})
		}
	}
	return results
}

func NewIndex(dictionary map[string][]int) *Index {
	index := new(Index)
	index.Dictionary = dictionary
	index.cache = make(map[string]int)
	return index
}
