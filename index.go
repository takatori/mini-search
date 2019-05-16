package index

import (
	"math"
)

const (
	BEGINNING_OF_FILE = -math.MaxInt32
	END_OF_FILE       = math.MaxInt32
)

var beginningPosition = &Position{
	BEGINNING_OF_FILE,
	BEGINNING_OF_FILE,
}

var endPosition = &Position{
	END_OF_FILE,
	END_OF_FILE,
}

type Index struct {
	dictionary map[string]*PostingsList
	cache      map[string]int // TODO: move
}

type PostingsList struct {
	list []*Posting
}

type Posting struct {
	docId         int
	termFrequency int
	offsets       []int
}

type Position struct {
	docId  int
	offset int
}

func ComparePosition(p1, p2 *Position) int {
	if (p1.docId > p2.docId) || (p1.docId == p2.docId && p1.offset > p2.offset) {
		return 1
	} else if p1.docId == p2.docId && p1.offset == p2.offset {
		return 0
	} else {
		return -1
	}
}

func (postingList *PostingsList) get(i int) *Position {

	var sum int

	for j, p := range postingList.list {
		if sum+p.termFrequency > i {
			return &Position{
				postingList.list[j].docId,
				postingList.list[j].offsets[i-sum],
			}
		}

		sum += p.termFrequency
	}

	return nil
}

func (postingList *PostingsList) FirstPosition() *Position {

	length := len(postingList.list)

	if length == 0 {
		return nil
	}

	return &Position{
		postingList.list[0].docId,
		postingList.list[0].offsets[0],
	}
}

func (postingList *PostingsList) LastPosition() *Position {

	length := len(postingList.list)

	if length == 0 {
		return nil
	}

	lastPosting := postingList.list[length-1]

	return &Position{
		lastPosting.docId,
		lastPosting.offsets[len(lastPosting.offsets)-1],
	}

}

// First(t) returns the first position at which the term  t occurs in the collection
func (index *Index) First(t string) *Position {

	if postingsList, ok := index.dictionary[t]; !ok {
		return endPosition
	} else {
		return postingsList.FirstPosition()
	}
}

func (index *Index) FirstDoc(t string) int {
	return docId(index.First(t))
}

// last(t) returns the last position at which t occurs in collection
func (index *Index) Last(t string) *Position {

	if postingsList, ok := index.dictionary[t]; !ok {
		return beginningPosition
	} else {
		return postingsList.LastPosition()
	}
}

func (index *Index) LastDoc(t string) int {
	return docId(index.Last(t))
}

// next(t, current) returns the position of t's first occurrence after the current position
func (index *Index) Next(t string, current *Position) *Position {

	var postingList *PostingsList
	var ok bool
	var low, high, jump int

	if postingList, ok = index.dictionary[t]; !ok {
		return endPosition
	}

	length := len(postingList.list)

	if length == 0 || ComparePosition(current, postingList.LastPosition()) > 0 {
		return endPosition
	}

	if ComparePosition(postingList.FirstPosition(), current) > 0 {
		index.cache[t] = 0
		return postingList.get(index.cache[t])
	}

	if index.cache[t] > 0 && ComparePosition(postingList.get(index.cache[t]), current) < 1 {
		low = index.cache[t]
	} else {
		low = 0
	}

	jump = 1
	high = low + jump

	for high < length && ComparePosition(postingList.get(high), current) < 1 {
		low = high
		jump = 2 * jump
		high = low + jump
	}

	if high > length {
		high = length
	}

	index.cache[t] = index.binarySearch(t, low, high, current)

	return postingList.get(index.cache[t])

}

func (index *Index) binarySearch(t string, low, high int, current *Position) int {

	for high-low > 1 {
		mid := (low + high) / 2
		if p, ok := index.dictionary[t]; !ok {
			return END_OF_FILE
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

/*
func (index *Index) binarySearchPrev(t string, low, high, current int) int {

	for high-low > 1 {
		mid := (low + high) / 2
		if p, ok := index.Dictionary[t]; !ok {
			return END_OF_FILE
		} else {
			if p[mid] < current {
				low = mid
			} else {
				high = mid
			}
		}
	}
	return low
}

// prev(t, current) returns the position of t's last occurrence before the current position
func (index *Index) prev(t string, current int) int {

	var p []int
	var ok bool
	var low, high, jump int

	if p, ok = index.Dictionary[t]; !ok {
		return BEGINNING_OF_FILE
	}

	length := len(p)

	if length == 0 || p[0] >= current {
		return BEGINNING_OF_FILE
	}

	if p[length-1] < current {
		index.cache[t] = length - 1
		return p[index.cache[t]]
	}

	if index.cache[t] > 0 && p[index.cache[t]+1] >= current {
		high = index.cache[t] + 1
	} else {
		high = length - 1
	}

	jump = 1
	low = high - jump

	for low > 0 && p[low] >= current {
		high = low
		jump = 2 * jump
		low = high - jump
	}

	if low < 0 {
		low = 0
	}

	index.cache[t] = index.binarySearchPrev(t, low, high, current)
	return p[index.cache[t]]
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
	var u, v int

	u = current

	for u < END_OF_FILE {
		u, v = index.nextPhrase(phrase, u)
		if u != END_OF_FILE {
			results = append(results, [2]int{u, v})
		}
	}
	return results
}
*/

func docId(position *Position) int {
	return position.docId
}
func offset(position *Position) interface{} {
	return position.offset
}

func NewIndex(dictionary map[string]*PostingsList) *Index {
	index := new(Index)
	index.dictionary = dictionary
	index.cache = make(map[string]int) // TODO: move query struct
	return index
}

func NewPostingsList(list []*Posting) *PostingsList {
	return &PostingsList{
		list: list,
	}
}

func NewPosting(docId int, offsets []int) *Posting {
	return &Posting{
		docId:         docId,
		termFrequency: len(offsets),
		offsets:       offsets,
	}
}
