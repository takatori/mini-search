package index

import (
	"math"
)

// Position represents a position of term in a document.
type Position struct {
	docId  int // document Id
	offset int // offset from the beginning of the document
}

const (
	BeginningOfFile = -math.MaxInt32
	EndOfFile       = math.MaxInt32
)

var BOF = &Position{
	BeginningOfFile,
	BeginningOfFile,
}

var EOF = &Position{
	EndOfFile,
	EndOfFile,
}

// Distance(p) returns a distance between two position.
func (p *Position) Distance(p2 *Position) int {

	if p.docId != p2.docId {
		return math.MaxInt32
	}

	distance := p.offset - p2.offset

	if distance < 0 {
		return -distance
	}

	return distance
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

func DocId(position *Position) int {
	return position.docId
}
func Offset(position *Position) int {
	return position.offset
}

// NewPosition return a new position.
func NewPosition(docId, offset int) *Position {
	return &Position{docId, offset}
}
