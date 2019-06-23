package index

import (
	"strings"
	"regexp"
)

type IndexWriter struct {
	index *Index
}

func (i *IndexWriter) AddDocument(doc string) {
	i.index.docCount++
	docId := i.index.docCount

	doc = strings.ToLower(doc)
	rep := regexp.MustCompile("[!,?.:]")
	doc = rep.ReplaceAllString(doc, "")
	terms := strings.Split(doc, " ")

	for j, term := range terms {
		if postingsList, ok := i.index.dictionary[term]; ok {
			if posting := postingsList.getByDocId(docId); posting != nil {
				posting.offsets = append(posting.offsets, j+1)
				posting.termFrequency++
			} else {
				postingsList.list = append(postingsList.list, &Posting{
					docId,
					[]int{j+1},
					1,
				})
			}
		} else {
			i.index.dictionary[term] = &PostingsList{
				[]*Posting{{docId, []int{j+1}, 1}},
			}
		}
	}
}

func (i *IndexWriter) Commit() *Index {
	return i.index
}

func NewIndexWriter() *IndexWriter {
	dict := make(map[string]*PostingsList)
	index := &Index{
		dictionary: dict,
		docCount: 0,
	}
	return &IndexWriter{index}
}

func NewIndex(dictionary map[string]*PostingsList) *Index {
	index := new(Index)
	index.dictionary = dictionary
	index.docCount = 0
	return index
}
