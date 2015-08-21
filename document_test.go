package main

import "testing"

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
		d, err := ParseDocument(testCase.input)
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
