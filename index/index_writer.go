package index

import (
	"strings"
	"regexp"
)

type IndexWriter struct {
	index *Index
}

func textToWordSequence(text string) []string {
	text = strings.ToLower(text)
	rep := regexp.MustCompile("[!,?.:]")
	text = rep.ReplaceAllString(text, "")
	return strings.Split(text, " ")
}

func (iw *IndexWriter) getDocumentId() int {
	return iw.index.docCount + 1
}

func (iw *IndexWriter) AddDocument(doc string) {

	docId := iw.getDocumentId()

	for j, term := range textToWordSequence(doc) {

		if pl, ok := iw.index.dictionary[term]; ok {

			if posting := pl.getByDocId(docId); posting != nil {
				posting.offsets = append(posting.offsets, j+1)
				posting.termFrequency++
			} else {
				iw.index.dictionary[term] = append(pl, &Posting{
					docId,
					[]int{j + 1},
					1,
				})
			}

		} else {
			iw.index.dictionary[term] = []*Posting{
				{docId, []int{j + 1}, 1},
			}
		}
	}

	iw.index.docCount++
}

func (iw *IndexWriter) Commit() *Index {
	return iw.index
}

// NewIndexWriter return new index writer
func NewIndexWriter() *IndexWriter {
	dict := make(map[string]PostingsList)
	index := &Index{
		dictionary: dict,
		docCount:   0,
	}
	return &IndexWriter{index}
}

func NewIndex(dictionary map[string]PostingsList) *Index {
	index := new(Index)
	index.dictionary = dictionary
	index.docCount = 0
	return index
}
