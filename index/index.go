package index

import (
	"math"
	"sort"
	"strings"
	"fmt"
)

// Index is an inverted index.
type Index struct {
	dictionary map[string]PostingsList
	docCount   int
}

func (idx Index) String() string {

	keys := make([]string, 0, len(idx.dictionary))

	for k := range idx.dictionary {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	str := make([]string, len(keys))

	for i, k := range keys {
		if postingList, ok := idx.dictionary[k]; ok {
			str[i] = fmt.Sprintf("'%s'->[%s]", k, postingList.String())
		}
	}

	return strings.Join(str, "\n")

}

// DocCount() return the number of document count
func (idx *Index) DocCount() int {
	return idx.docCount
}

// DocFrequency(t) return the number of term document count
func (idx *Index) DocFrequency(t string) int {
	if postingsList, ok := idx.dictionary[t]; !ok {
		return 0
	} else {
		return len(postingsList)
	}
}

func (idx *Index) TF(term string, docId int) float64 {
	if postingsList, ok := idx.dictionary[term]; !ok {
		return 0
	} else {
		return postingsList.tf(docId)
	}
}

func (idx *Index) IDF(term string) float64 {
	N := idx.DocCount()
	Nt := idx.DocFrequency(term)
	return math.Log2(float64(N) / float64(Nt))
}

// TF_IDF(term, DocId) return tf-idf score
func (idx *Index) TF_IDF(term string, docId int) float64 {
	return idx.TF(term, docId) * idx.IDF(term)
}

// First(t) returns the first position at which the term  t occurs in the collection
func (idx *Index) First(t string) *Position {

	if postingsList, ok := idx.dictionary[t]; !ok {
		return EOF
	} else {
		return postingsList.FirstPosition()
	}
}

// last(t) returns the last position at which t occurs in collection
func (idx *Index) Last(t string) *Position {

	if postingsList, ok := idx.dictionary[t]; !ok {
		return BOF
	} else {
		return postingsList.LastPosition()
	}
}

// return the docid of the first document containing the term t
func (idx *Index) FirstDoc(t string) int {
	return DocId(idx.First(t))
}

// return the docid of the last document containing the term t
func (idx *Index) LastDoc(t string) int {
	return DocId(idx.Last(t))
}

// next(t, current) returns the position of t's first occurrence after the current position
func (idx *Index) Next(t string, current *Position) *Position {

	postingList, ok := idx.dictionary[t];
	if !ok {
		return EOF
	}

	length := postingList.length()

	if length == 0 || ComparePosition(current, postingList.LastPosition()) >= 0 {
		return EOF
	}

	if ComparePosition(postingList.FirstPosition(), current) > 0 {
		return postingList.FirstPosition()
	}

	low := 0
	jump := 1
	high := low + jump

	for high < length && ComparePosition(postingList.get(high), current) < 1 {
		low = high
		jump = 2 * jump
		high = low + jump
	}

	if high > length {
		high = length
	}

	return postingList.get(idx.binarySearch(t, low, high, current))

}

func (idx *Index) NextDoc(t string, current int) int {
	return idx.Next(t, NewPosition(current, EndOfFile)).docId
}

// prev(t, current) returns the position of t's last occurrence before the current position
func (idx *Index) Prev(t string, current *Position) *Position {

	postingList, ok := idx.dictionary[t];
	if !ok {
		return BOF
	}

	length := postingList.length()

	if length == 0 || ComparePosition(current, postingList.FirstPosition()) <= 0 {
		return BOF
	}

	if ComparePosition(postingList.LastPosition(), current) < 0 {
		return postingList.LastPosition()
	}

	high := length - 1
	jump := 1
	low := high - jump

	for low > 0 && ComparePosition(postingList.get(low), current) >= 0 {
		high = low
		jump = 2 * jump
		low = high - jump
	}

	if low < 0 {
		low = 0
	}

	return postingList.get(idx.binarySearchPrev(t, low, high, current))
}

func (idx *Index) PrevDoc(t string, current int) int {
	return idx.Prev(t, NewPosition(current, BeginningOfFile)).docId
}

func (idx *Index) binarySearch(t string, low, high int, current *Position) int {

	for high-low > 1 {
		mid := (low + high) / 2
		if p, ok := idx.dictionary[t]; !ok {
			return EndOfFile
		} else {
			if ComparePosition(p.get(mid), current) < 1 {
				low = mid
			} else {
				high = mid
			}
		}
	}
	return high
}

func (idx *Index) binarySearchPrev(t string, low, high int, current *Position) int {

	for high-low > 1 {
		mid := (low + high) / 2
		if p, ok := idx.dictionary[t]; !ok {
			return EndOfFile
		} else {
			if ComparePosition(p.get(mid), current) >= 0 {
				high = mid
			} else {
				low = mid
			}
		}
	}
	return low
}
