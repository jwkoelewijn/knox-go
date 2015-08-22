package main

import "fmt"

type Solution struct {
	Documents DocumentList
	Cost      int
	Bandwidth int
	Pool      DocumentList
}

func NewSolution(documents, pool DocumentList, bandwidth int) Solution {
	sol := Solution{Documents: documents, Cost: documents.Cost(), Bandwidth: bandwidth, Pool: pool}
	return sol
}

func (s Solution) Value() int {
	return s.Documents.Value()
}

func (s Solution) LeftOver() int {
	return s.Bandwidth - s.Cost
}

func (sol *Solution) Maximize() {
	sol.Pool.SortByDescendingRatio()

	for sol.Cost < sol.Bandwidth {
		newDoc, err := sol.findFittingDocument()
		if err != nil {
			break
		}
		sol.Documents = append(sol.Documents, newDoc)
		sol.Cost += newDoc.Size
	}
}

func (sol Solution) findFittingDocument() (Document, error) {
	maxSize := sol.LeftOver()
	bestSecrecy := 0.0
	found := false
	var bestDoc Document
	for _, d := range sol.Pool {
		if d.Size < maxSize && !sol.Documents.Contains(d) {
			if d.SecrecyRatio > bestSecrecy {
				bestDoc = d
				found = true
			}
		}
	}
	if found {
		return bestDoc, nil
	} else {
		return bestDoc, fmt.Errorf("Could not find fitting document")
	}
}
