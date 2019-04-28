package index

import (
	"testing"
)

func TestFirst(t *testing.T) {

	index := NewIndex(map[string][]int{
		"hurlybufly": {316669, 745434},
		"witching":   {265197},
	})

	testCases := map[string]int{
		"hurlybufly": 316669,
		"witching":   265197,
	}

	for param, expected := range testCases {
		actual := index.first(param)
		if actual != expected {
			t.Errorf("\n got: %v\n want: %v", actual, expected)
		}
	}
}

func TestLast(t *testing.T) {

	index := NewIndex(map[string][]int{
		"thunder":  {36898, 137236, 745397, 745419, 1247139},
		"witching": {265197},
	})

	testCases := map[string]int{
		"thunder":  1247139,
		"witching": 265197,
	}

	for param, expected := range testCases {
		actual := index.last(param)
		if actual != expected {
			t.Errorf("\n got: %v\n want: %v", actual, expected)
		}
	}
}

func TestNext(t *testing.T) {

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
		{"witch", 745429, 745451},
		{"hurlyburly", 345678, 745434},
		{"witch", 1245276, END_OF_FILE},
		{"witch", BEGINNING_OF_FILE, 1598},
		{"witch", END_OF_FILE, END_OF_FILE},
	}

	for _, testCase := range testCases {
		actual := index.next(testCase.t, testCase.current)
		if actual != testCase.expected {
			t.Errorf("\n got: %v\n want: %v", actual, testCase.expected)
		}
	}

}

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
