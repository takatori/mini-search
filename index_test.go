package index

import (
	"testing"
)

func TestDocId(t *testing.T) {

	testCases := map[*Position]int{
		&Position{123, 456}: 123,
		&Position{123, 789}: 123,
	}

	for param, expected := range testCases {
		actual := docId(param)
		if actual != expected {
			t.Errorf("\n got: %v\n want: %v", actual, expected)
		}
	}
}

func TestOffset(t *testing.T) {

	testCases := map[*Position]int{
		&Position{123, 456}: 456,
		&Position{123, 789}: 789,
	}

	for param, expected := range testCases {
		actual := offset(param)
		if actual != expected {
			t.Errorf("\n got: %v\n want: %v", actual, expected)
		}
	}
}

func TestFirst(t *testing.T) {

	index := NewIndex(map[string]*PostingsList{
		"hurlybufly": NewPostingsList(
			[]*Posting{
				NewPosting(9, []int{30963}),
				NewPosting(22, []int{293}),
			}),
		"witching": NewPostingsList(
			[]*Posting{
				NewPosting(8, []int{25805}),
			}),
	})

	testCases := map[string]Position{
		"hurlybufly": {9, 30963},
		"witching":   {8, 25805},
	}

	for param, expected := range testCases {
		actual := index.First(param)
		if *actual != expected {
			t.Errorf("\n got: %v\n want: %v", actual, expected)
		}
	}
}

func TestFirstDoc(t *testing.T) {

	index := NewIndex(map[string]*PostingsList{
		"hurlybufly": NewPostingsList(
			[]*Posting{
				NewPosting(9, []int{30963}),
				NewPosting(22, []int{293}),
			}),
		"witching": NewPostingsList(
			[]*Posting{
				NewPosting(8, []int{25805}),
			}),
	})

	testCases := map[string]int{
		"hurlybufly": 9,
		"witching":   8,
	}

	for param, expected := range testCases {
		actual := index.FirstDoc(param)
		if actual != expected {
			t.Errorf("\n got: %v\n want: %v", actual, expected)
		}
	}
}

func TestLast(t *testing.T) {

	index := NewIndex(map[string]*PostingsList{
		"thunder": NewPostingsList(
			[]*Posting{
				NewPosting(1, []int{36898}),
				NewPosting(5, []int{6402}),
				NewPosting(22, []int{256, 278}),
				NewPosting(37, []int{12538, 40000}),
			}),
		"witching": NewPostingsList(
			[]*Posting{
				NewPosting(8, []int{25805}),
			}),
	})



	testCases := map[string]Position{
		"thunder":  {37, 40000},
		"witching": {8, 25805},
	}

	for param, expected := range testCases {
		actual := index.Last(param)
		if *actual != expected {
			t.Errorf("\n got: %v\n want: %v", actual, expected)
		}
	}
}

func TestLastDoc(t *testing.T) {

	index := NewIndex(map[string]*PostingsList{
		"thunder": NewPostingsList(
			[]*Posting{
				NewPosting(1, []int{36898}),
				NewPosting(5, []int{6402}),
				NewPosting(22, []int{256, 278}),
				NewPosting(37, []int{12538}),
			}),
		"witching": NewPostingsList(
			[]*Posting{
				NewPosting(8, []int{25805}),
			}),
	})

	testCases := map[string]int{
		"thunder":  37,
		"witching": 8,
	}

	for param, expected := range testCases {
		actual := index.LastDoc(param)
		if actual != expected {
			t.Errorf("\n got: %v\n want: %v", actual, expected)
		}
	}
}

func TestNext(t *testing.T) {

	index := NewIndex(map[string]*PostingsList{
		"hurlybufly": NewPostingsList(
			[]*Posting{
				NewPosting(9, []int{30963}),
				NewPosting(22, []int{293}),
			}),
		"thunder": NewPostingsList(
			[]*Posting{
				NewPosting(1, []int{36898}),
				NewPosting(5, []int{6402}),
				NewPosting(22, []int{256, 278}),
				NewPosting(37, []int{12538}),
			}),
		"witching": NewPostingsList(
			[]*Posting{
				NewPosting(8, []int{25805}),
			}),
	})


	type test struct {
		t        string
		current  Position
		expected Position
	}

	testCases := []test{
		//{"witch", &Position{22, 288}, &Position{22, 310}},
		{"hurlybufly", Position{9, 30963}, Position{22, 293}},
		{"witch", Position{37, 10675}, Position{END_OF_FILE, END_OF_FILE}},
	}

	for _, testCase := range testCases {
		actual := index.Next(testCase.t, &testCase.current)
		if *actual != testCase.expected {
			t.Errorf("\n got: %v\n want: %v", actual, testCase.expected)
		}
	}

}

/*
func TestPrev(t *testing.T) {

	index := NewIndex(map[string][]int{
		"hurlyburly": {316669, 745434},
		"thunder":    {36898, 137236, 745397, 745419, 1247139},
		"witch":      {1598, 27555, 745407, 745429, 745451, 745467, 1245276},
		"witching":   {265197},
	})

	type test struct {
		t        string
		current  int
		expected int
	}

	testCases := []test{
		{"witch", 745451, 745429},
		{"hurlyburly", 456789, 316669},
		{"witch", 1598, BEGINNING_OF_FILE},
		{"witch", END_OF_FILE, 1245276},
		{"witch", BEGINNING_OF_FILE, BEGINNING_OF_FILE},
	}

	for _, testCase := range testCases {
		actual := index.prev(testCase.t, testCase.current)
		if actual != testCase.expected {
			t.Errorf("\n got: %v\n want: %v", actual, testCase.expected)
		}
	}

}

func TestNextPhrase(t *testing.T) {

	index := NewIndex(map[string][]int{
		"first": {2205, 2268, 745406, 745466, 745501, 1271487},
		"witch": {1598, 27555, 745407, 745429, 745451, 745467, 745502, 1245276},
	})

	type test struct {
		t             []string
		current       int
		expectedStart int
		expectedEnd   int
	}

	testCases := []test{
		{[]string{"first", "witch"}, BEGINNING_OF_FILE, 745406, 745407},
		{[]string{"first", "witch"}, 745500, 745501, 745502},
	}

	for _, testCase := range testCases {
		u, v := index.nextPhrase(testCase.t, testCase.current)
		if u != testCase.expectedStart || v != testCase.expectedEnd {
			t.Errorf("\n got: [%v, %v]\n want: [%v, %v]", u, v, testCase.expectedStart, testCase.expectedEnd)
		}
	}

}

func TestAllPhrase(t *testing.T) {

	index := NewIndex(map[string][]int{
		"first": {2205, 2268, 745406, 745466, 745501, 1271487},
		"witch": {1598, 27555, 745407, 745429, 745451, 745467, 745502, 1245276},
	})

	type test struct {
		t        []string
		current  int
		expected [][]int
	}

	testCases := []test{
		{[]string{"first", "witch"},
			BEGINNING_OF_FILE,
			[][]int{
				{745406, 745407},
				{745466, 745467},
				{745501, 745502},
			},
		},
	}
	for _, testCase := range testCases {
		results := index.allPhrase(testCase.t, testCase.current)
		for i, result := range results {
			if result[0] != testCase.expected[i][0] || result[1] != testCase.expected[i][1] {
				t.Errorf("\n got: %v\n want: %v", result, testCase.expected[i])
			}
		}
	}
}
*/