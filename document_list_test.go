package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestCost(t *testing.T) {
	testCases := []testCase{
		{
			input: "list -- |                CriminalKeys.rb    713KB            35",
		},
		{
			input: "   list -- |             TransitBuilder.ppt   1980KB            58",
		},
	}

	docList := NewDocumentList()
	for _, tc := range testCases {
		d, err := ParseDocument(tc.input)
		if err != nil {
			t.Errorf("Should not error on input '%s'", tc.input)
		}
		docList = docList.Add(d)
	}
	if cost := docList.Cost(); cost != 713+1980 {
		t.Errorf("Expected the cost to be %d, not %d", 713+1980, cost)
	}
}

func TestSorting(t *testing.T) {
	docList := NewDocumentList()

	size := 10
	for i := 1; i <= size; i += 1 {
		value := rand.Intn(size)
		docList = docList.Add(Document{Name: fmt.Sprintf("Document %d", 10-i), Size: i, Value: value, SecrecyRatio: float64(value) / float64(i)})
	}

	fmt.Println("docList:")
	fmt.Println(docList)

	docList.SortByName()
	fmt.Println("docList (by Name):")
	fmt.Println(docList)

	docList.SortBySize()
	fmt.Println("docList (by Size):")
	fmt.Println(docList)

	docList.SortByDescendingValue()
	fmt.Println("docList (by Descending Value):")
	fmt.Println(docList)

	docList.SortByRatio()
	fmt.Println("docList: (by Ratio)")
	fmt.Println(docList)

	docList.SortByDescendingRatio()
	fmt.Println("docList: (by Descending Ratio)")
	fmt.Println(docList)
}
