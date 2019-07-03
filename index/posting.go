package index

import (
	"math"
	"fmt"
	"strings"
)

type Posting struct {
	docId         int
	offsets       []int
	termFrequency int
}

func (p Posting) String() string {
	return fmt.Sprintf("<DocId: %d, offsets:%v>", p.docId, p.offsets)
}

func NewPosting(docId int, offsets []int) *Posting {
	return &Posting{
		docId:         docId,
		termFrequency: len(offsets),
		offsets:       offsets,
	}
}

type PostingsList []*Posting

func (pl PostingsList) String() string {
	str := make([]string, len(pl))
	for i, p := range pl {
		str[i] = p.String()
	}
	return strings.Join(str, " ")
}

func (pl PostingsList) length() int {

	n := 0
	for _, p := range pl {
		n += len(p.offsets)
	}
	return n
}

func (pl PostingsList) get(i int) *Position {

	var sum int

	for j, p := range pl {
		if sum+p.termFrequency > i {
			return &Position{
				pl[j].docId,
				pl[j].offsets[i-sum],
			}
		}

		sum += p.termFrequency
	}

	return nil
}

func (pl PostingsList) getByDocId(docId int) *Posting {
	// TODO: binary search
	for _, posting := range pl {
		if posting.docId == docId {
			return posting
		}
	}
	return nil
}

func (pl PostingsList) FirstPosition() *Position {

	if len(pl) == 0 {
		return nil
	}

	return &Position{
		pl[0].docId,
		pl[0].offsets[0],
	}
}

func (pl PostingsList) LastPosition() *Position {

	length := len(pl)

	if length == 0 {
		return nil
	}

	lastPosting := pl[length-1]

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