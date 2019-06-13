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
				NewPosting(9, []int{40963}),
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
		{"hurlybufly", Position{9, 30963}, Position{9, 40963}},
		{"hurlybufly", Position{9, 40000}, Position{9, 40963}},
		{"witch", Position{37, 10675}, Position{EndOfFile, EndOfFile}},
		{"witch", Position{37, 12538}, Position{EndOfFile, EndOfFile}},
	}

	for _, testCase := range testCases {
		actual := index.Next(testCase.t, &testCase.current)
		if *actual != testCase.expected {
			t.Errorf("\n got: %v\n want: %v", actual, testCase.expected)
		}
	}

}

func TestNextDoc(t *testing.T) {

	index := NewIndex(map[string]*PostingsList{
		"hurlybufly": NewPostingsList(
			[]*Posting{
				NewPosting(9, []int{30963}),
				NewPosting(9, []int{40963}),
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
		current  int
		expected int
	}

	testCases := []test{
		{"hurlybufly", 9, 22},
		{"hurlybufly", 22, EndOfFile},
		{"witching", BeginningOfFile, 8},
	}

	for _, testCase := range testCases {
		actual := index.NextDoc(testCase.t, testCase.current)
		if actual != testCase.expected {
			t.Errorf("\n got: %v\n want: %v", actual, testCase.expected)
		}
	}

}


func TestPrev(t *testing.T) {

	index := NewIndex(map[string]*PostingsList{
		"hurlyburly": NewPostingsList(
			[]*Posting{
				NewPosting(9, []int{30963}),
				NewPosting(22, []int{290}),
				NewPosting(22, []int{293}),
			}),
		"thunder": NewPostingsList(
			[]*Posting{
				NewPosting(1, []int{36898}),
				NewPosting(5, []int{6402}),
				NewPosting(22, []int{256, 278}),
				NewPosting(37, []int{12538}),
			}),
		"witch": NewPostingsList(
			[]*Posting{
				NewPosting(1, []int{1598, 27555}),
				NewPosting(22, []int{266, 288, 310, 326}),
				NewPosting(37, []int{10675}),
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
		{"hurlyburly", Position{22, 29309}, Position{22, 293}},
		{"witch", Position{22, 310}, Position{22, 288}},
		{"witch", Position{1, 1598}, Position{BeginningOfFile, BeginningOfFile}},
		{"witch", Position{EndOfFile, EndOfFile}, Position{37, 10675}},
	}
	for _, testCase := range testCases {
		actual := index.Prev(testCase.t, &testCase.current)
		if *actual != testCase.expected {
			t.Errorf("\n got: %v\n want: %v", *actual, testCase.expected)
		}
	}
}

func TestPrevDoc(t *testing.T) {

	index := NewIndex(map[string]*PostingsList{
		"hurlyburly": NewPostingsList(
			[]*Posting{
				NewPosting(9, []int{30963}),
				NewPosting(22, []int{290}),
				NewPosting(22, []int{293}),
			}),
		"thunder": NewPostingsList(
			[]*Posting{
				NewPosting(1, []int{36898}),
				NewPosting(5, []int{6402}),
				NewPosting(22, []int{256, 278}),
				NewPosting(37, []int{12538}),
			}),
		"witch": NewPostingsList(
			[]*Posting{
				NewPosting(1, []int{1598, 27555}),
				NewPosting(22, []int{266, 288, 310, 326}),
				NewPosting(37, []int{10675}),
			}),
		"witching": NewPostingsList(
			[]*Posting{
				NewPosting(8, []int{25805}),
			}),
	})

	type test struct {
		t        string
		current  int
		expected int
	}

	testCases := []test{
		{"hurlyburly", 22, 9},
		{"witch", 37, 22},
		{"witch", 1, BeginningOfFile},
		{"witch", EndOfFile, 37},
	}
	for _, testCase := range testCases {
		actual := index.PrevDoc(testCase.t, testCase.current)
		if actual != testCase.expected {
			t.Errorf("\n got: %v\n want: %v", actual, testCase.expected)
		}
	}
}
