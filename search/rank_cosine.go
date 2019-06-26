package search

import (
	"github.com/takatori/mini-search/index"
	"sort"
	"math"
)

func cosineSim(v1, v2 []float64) float64 {

	var score float64

	//v1 = normalize(v1)
	//v2 = normalize(v2)
	//fmt.Printf("v1:%v, v2:%v\n", v1, v2)

	for i, v1s := range v1 {
		score += v1s * v2[i]
	}

	return score
}

func docVector(idx *index.Index, terms []string, docId int) []float64 {

	dV := make([]float64, len(terms))

	for i, term := range terms {
		dV[i] = idx.TF_IDF(term, docId)
	}

	return dV
}

func queryVector(idx *index.Index, terms []string) []float64 {

	qV := make([]float64, len(terms))

	for i, term := range terms {
		qV[i] = idx.IDF(term) // TODO
	}

	return qV
}

func normalize(v []float64) []float64 {

	var dist float64

	for _, e := range v {
		dist += e * e
	}

	dist = math.Sqrt(dist)

	results := make([]float64, len(v))

	for i, e := range v {
		results[i] = e / dist
	}

	return results
}

func nextMinDoc(idx *index.Index, terms []string, docId int) int {

	d := index.EndOfFile

	for _, t := range terms {
		tmp := idx.NextDoc(t, docId)
		if tmp < d {
			d = tmp
		}
	}

	return d
}

func RankCosine(idx *index.Index, terms []string, k int) []int {

	results := make([]*result, 0, k)
	d := nextMinDoc(idx, terms, index.BeginningOfFile)
	qV := queryVector(idx, terms)

	for i := 0; d < index.EndOfFile && i < k; i++ {
		results = append(results, &result{
			docId: d,
			score: cosineSim(qV, docVector(idx, terms, d)),
		})
		d = nextMinDoc(idx, terms, d)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].score > results[j].score
	})

	docIds := make([]int, len(results))
	for i, r := range results {
		docIds[i] = r.docId
	}
	return docIds
}
