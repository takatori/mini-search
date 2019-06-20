package index

import "math"

type Posting struct {
	docId         int
	offsets       []int
	termFrequency int
}

type PostingsList struct {
	list []*Posting
}

func (l *PostingsList) length() int {

	n := 0
	for _, p := range l.list {
		n += len(p.offsets)
	}
	return n
}

func (l *PostingsList) get(i int) *Position {

	var sum int

	for j, p := range l.list {
		if sum+p.termFrequency > i {
			return &Position{
				l.list[j].docId,
				l.list[j].offsets[i-sum],
			}
		}

		sum += p.termFrequency
	}

	return nil
}

func (l *PostingsList) getByDocId(docId int) *Posting {
	// TODO: binary search
	for _, posting := range l.list {
		if posting.docId == docId {
			return posting
		}
	}
	return nil
}

func (l *PostingsList) FirstPosition() *Position {

	if len(l.list) == 0 {
		return nil
	}

	return &Position{
		l.list[0].docId,
		l.list[0].offsets[0],
	}
}

func (l *PostingsList) LastPosition() *Position {

	length := len(l.list)

	if length == 0 {
		return nil
	}

	lastPosting := l.list[length-1]

	return &Position{
		lastPosting.docId,
		lastPosting.offsets[len(lastPosting.offsets)-1],
	}

}
func (l *PostingsList) tf(docId int) float64 {

	if p := l.getByDocId(docId); p != nil && p.termFrequency > 0 {
		return math.Log2(float64(p.termFrequency)) + 1
	} else {
		return 0
	}
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
