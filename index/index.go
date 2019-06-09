package index

type Index struct {
	dictionary map[string]*PostingsList
	cache      map[string]int // TODO: move
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

func (idx *Index) FirstDoc(t string) int {
	return docId(idx.First(t))
}

func (idx *Index) LastDoc(t string) int {
	return docId(idx.Last(t))
}

// next(t, current) returns the position of t's first occurrence after the current position
func (idx *Index) Next(t string, current *Position) *Position {

	var postingList *PostingsList
	var ok bool
	var low, high, jump int

	if postingList, ok = idx.dictionary[t]; !ok {
		return EOF
	}

	length := postingList.length()

	if length == 0 || ComparePosition(current, postingList.LastPosition()) >= 0 {
		return EOF
	}

	if ComparePosition(postingList.FirstPosition(), current) > 0 {
		//idx.cache[t] = 0
		//return postingList.get(idx.cache[t])
		return postingList.get(0)
	}

	/*
	if idx.cache[t] > 0 && ComparePosition(postingList.get(idx.cache[t]), current) < 1 {
		low = idx.cache[t]
	} else {
		low = 0
	}*/

	low = 0

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

	//idx.cache[t] = idx.binarySearch(t, low, high, current)

	//return postingList.get(idx.cache[t])

	return postingList.get(idx.binarySearch(t, low, high, current))

}

// prev(t, current) returns the position of t's last occurrence before the current position
func (idx *Index) Prev(t string, current *Position) *Position {

	var postingList *PostingsList
	var ok bool

	if postingList, ok = idx.dictionary[t]; !ok {
		return BOF
	}

	length := postingList.length()

	if length == 0 || ComparePosition(current, postingList.FirstPosition()) <= 0 {
		return BOF
	}

	if ComparePosition(postingList.LastPosition(), current) < 0 {
		//idx.cache[t] = length - 1
		//return postingList.get(idx.cache[t])
		return postingList.LastPosition()
	}

	/*
	if idx.cache[t] > 0 && ComparePosition(current, postingList.get(idx.cache[t]+1)) < 0 {
		high = idx.cache[t] + 1
	} else {
		high = length - 1
	}*/

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

	// idx.cache[t] = idx.binarySearchPrev(t, low, high, current)

	// return postingList.get(idx.cache[t])

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
