package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	input   string
	name    string
	size    int
	value   int
	errored bool
}

func TestUnmarshal(t *testing.T) {
	testCases := []testCase{
		{
			input:   "list -- |                CriminalKeys.rb    713KB            35",
			name:    "CriminalKeys.rb",
			size:    713,
			value:   35,
			errored: false,
		},
		{
			input:   "   list -- | Remaining Bandwidth:  4371 KB",
			errored: true,
		},
		{
			input:   "   list -- |             TransitBuilder.ppt   1980KB            58",
			name:    "TransitBuilder.ppt",
			size:    1980,
			value:   58,
			errored: false,
		},
	}

	for _, testCase := range testCases {
		var d Document
		err := d.Unmarshal(testCase.input)
		if (err == nil) != !testCase.errored {
			t.Errorf("Expected Unmarshal('%s') to error where it didn't or the other way around", testCase.input)
		}
		if !testCase.errored {
			if d.Name != testCase.name {
				t.Errorf("Expected name to equal %s, found %s", testCase.name, d.Name)
			}
			if d.Size != testCase.size {
				t.Errorf("Expected size to equal %d, found %d", testCase.size, d.Size)
			}
			if d.Value != testCase.value {
				t.Errorf("Expected value to equal %d, found %d", testCase.value, d.Value)
			}
		}
	}
}

func TestCombinations(t *testing.T) {
	docList := NewDocumentList()

	for i := 1; i <= 10; i += 1 {
		docList = docList.Add(Document{Name: fmt.Sprintf("Document %d", 10-i), Size: i, Value: 10 - i, SecrecyRatio: float64(10-i) / float64(i)})
	}

	//name := func(d1, d2 *Document) bool {
	//	return d1.Name < d2.Name
	//}

	secrecyRatio := func(d1, d2 *Document) bool {
		return d1.SecrecyRatio < d2.SecrecyRatio
	}

	secrecyRatioReverse := func(d1, d2 *Document) bool {
		return !secrecyRatio(d1, d2)
	}

	fmt.Println("docList:")
	fmt.Println(docList)

	By(secrecyRatio).Sort(*docList)

	fmt.Println("docList:")
	fmt.Println(docList)
	By(secrecyRatioReverse).Sort(*docList)
	fmt.Println("docList:")
	fmt.Println(docList)
	fmt.Println("--------------- length 1")
	//docList.Combinations(1)
	fmt.Println("--------------- length 2")
	//docList.Combinations(2)
	fmt.Println("--------------- length 3")
	//docList.Combinations(3)
}
