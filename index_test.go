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
		"thunder":  {36898, 137236, 745397, 745419, 1247139},
		"witch": {1598, 27555, 745407, 745429, 745451, 745467, 1245276},
		"witching": {265197},
	})

	type test struct{
		t string
		current int
		expected int
	}

	testCases := []test{
		{"witch", 745429, 745451},
		{"hurlyburly", 345678, 745434},
		{"witch", 1245276, END_OF_FILE},
		{"witch", BEGINNING_OF_FILE, 1598},
		{"witch", END_OF_FILE, END_OF_FILE},
	}


	for _, textCase := range testCases {
		actual := index.next(textCase.t, textCase.current)
		if actual != textCase.expected {
			t.Errorf("\n got: %v\n want: %v", actual, textCase.expected)
		}
	}

}

func TestPrev(t *testing.T) {

	index := NewIndex(map[string][]int{
		"hurlyburly": {316669, 745434},
		"thunder":  {36898, 137236, 745397, 745419, 1247139},
		"witch": {1598, 27555, 745407, 745429, 745451, 745467, 1245276},
		"witching": {265197},
	})

	type test struct{
		t string
		current int
		expected int
	}

	testCases := []test{
		{"witch", 745451,745429},
		{"hurlyburly", 456789, 316669},
		{"witch", 1598, BEGINNING_OF_FILE},
		{"witch", END_OF_FILE, 1245276},
		{"witch", BEGINNING_OF_FILE, BEGINNING_OF_FILE},
	}


	for _, textCase := range testCases {
		actual := index.prev(textCase.t, textCase.current)
		if actual != textCase.expected {
			t.Errorf("\n got: %v\n want: %v", actual, textCase.expected)
		}
	}

}

