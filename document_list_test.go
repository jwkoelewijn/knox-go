package main

import "testing"

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
		var d Document
		err := d.Unmarshal(tc.input)
		if err != nil {
			t.Errorf("Should not error on input '%s'", tc.input)
		}
		docList = docList.Add(d)
	}
	if cost := docList.Cost(); cost != 713+1980 {
		t.Errorf("Expected the cost to be %d, not %d", 713+1980, cost)
	}
}
