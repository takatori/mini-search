package search

import (
	"testing"
	"github.com/takatori/mini-search/index"
)

func TestNextPhrase(t *testing.T) {

	idx := index.NewIndex(map[string]index.PostingsList{
		"first": []*index.Posting{
			index.NewPosting(1, []int{2205, 2268, 745406, 745466, 745501, 1271487}),
			index.NewPosting(22, []int{265, 235, 360}),
			index.NewPosting(37, []int{36886}),
		},
		"witch": []*index.Posting{
			index.NewPosting(1, []int{1598, 27555, 745407, 745429, 745451, 745467, 745502, 1274527}),
			index.NewPosting(22, []int{266, 288, 310, 326}),
			index.NewPosting(37, []int{10675}),
		},
	})

	type test struct {
		t             []string
		current       *index.Position
		expectedStart *index.Position
		expectedEnd   *index.Position
	}

	testCases := []test{
		{[]string{"first", "witch"},
			index.BOF,
			index.NewPosition(1, 745406),
			index.NewPosition(1, 745407),
		},
		{[]string{"first", "witch"},
			index.NewPosition(1, 745500),
			index.NewPosition(1, 745501),
			index.NewPosition(1, 745502),
		},
	}

	for _, testCase := range testCases {
		u, v := NextPhrase(idx, testCase.t, testCase.current)
		if *u != *testCase.expectedStart || *v != *testCase.expectedEnd {
			t.Errorf("\n got : [%v, %v]\n want: [%v, %v]", u, v, testCase.expectedStart, testCase.expectedEnd)
		}
	}

}

func TestPhraseSearch(t *testing.T) {

	idx := index.NewIndex(map[string]index.PostingsList{
		"first": []*index.Posting{
			index.NewPosting(1, []int{2205, 2268, 745406, 745466, 745501, 1271487}),
			index.NewPosting(22, []int{265, 235, 360}),
			index.NewPosting(37, []int{36886}),
		},
		"witch": []*index.Posting{
			index.NewPosting(1, []int{1598, 27555, 745407, 745429, 745451, 745467, 745502, 1274527}),
			index.NewPosting(22, []int{266, 288, 310, 326}),
			index.NewPosting(37, []int{10675}),
		},
	})

	type test struct {
		t        []string
		expected [][2]*index.Position
	}

	testCases := []test{
		{[]string{"first", "witch"},
			[][2]*index.Position{
				{index.NewPosition(1, 745406), index.NewPosition(1, 745407)},
				{index.NewPosition(1, 745466), index.NewPosition(1, 745467)},
				{index.NewPosition(1, 745501), index.NewPosition(1, 745502)},
			},
		},
	}
	for _, testCase := range testCases {
		results := PhraseSearch(idx, testCase.t)
		for i, result := range results {
			if *result[0] != *testCase.expected[i][0] || *result[1] != *testCase.expected[i][1] {
				t.Errorf("\n got : %v\n want: %v", result, testCase.expected[i])
			}
		}
	}
}
