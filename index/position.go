package index

import (
	"math"
)

type Position struct {
	docId  int
	offset int
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

func (p *Position) Distance(p2 *Position) int {
	if p.docId != p2.docId {
		return math.MaxInt32 // TODO: fix
	}
	if p.offset - p2.offset > 0 {
		return p.offset - p2.offset
	} else {
		return p2.offset - p.offset
	}
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

func NewPosition(docId, offset int) *Position {
	return &Position{docId, offset}
}
