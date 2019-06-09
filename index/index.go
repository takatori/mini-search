package index

type Index struct {
	dictionary map[string]*PostingsList
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
	return docId(idx.First(t))
}

// return the docid of the last document containing the term t
func (idx *Index) LastDoc(t string) int {
	return docId(idx.Last(t))
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
		return postingList.get(0)
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
