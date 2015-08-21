package main

import "testing"

func TestOptimization(t *testing.T) {
	docList := DocumentList{
		{
			Name:         "ONSPasswords.cer",
			Size:         2406,
			Value:        86,
			SecrecyRatio: float64(86) / float64(2406),
		},
		{
			Name:         "CriminalBuilder.rb",
			Size:         2454,
			Value:        41,
			SecrecyRatio: float64(41) / float64(2454),
		},
		{
			Name:         "LowlandsReport.txt",
			Size:         2014,
			Value:        12,
			SecrecyRatio: float64(12) / float64(2014),
		},
		{
			Name:         "Store!IDReport.doc",
			Size:         843,
			Value:        72,
			SecrecyRatio: float64(72) / float64(843),
		},
		{
			Name:         "PowerRouterReport.rb",
			Size:         2200,
			Value:        83,
			SecrecyRatio: float64(83) / float64(2200),
		},
		{
			Name:         "KittenKeys.ppt",
			Size:         2018,
			Value:        94,
			SecrecyRatio: float64(94) / float64(2018),
		},
		{
			Name:         "ShadowKeys.txt",
			Size:         1878,
			Value:        81,
			SecrecyRatio: float64(81) / float64(1878),
		},
		{
			Name:         "HolidayRumors.txt",
			Size:         1655,
			Value:        66,
			SecrecyRatio: float64(66) / float64(1655),
		},
		{
			Name:         "SalesKeys.cer",
			Size:         1295,
			Value:        38,
			SecrecyRatio: float64(38) / float64(1295),
		},
		{
			Name:         "AEOSLaptops.cer",
			Size:         1196,
			Value:        63,
			SecrecyRatio: float64(63) / float64(1196),
		},
	}

	solution := docList.FindBestSolution(4408)

	//fmt.Printf("FINAL GUESS with Secrecy value %d: %+v", solution.Value(), solution)
	if solution.Value() != 229 {
		t.Errorf("Expected solution to have value 229, found %d", solution.Value())
	}
}
