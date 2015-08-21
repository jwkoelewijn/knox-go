package main

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
	sum := 0
	for _, doc := range s.Documents {
		sum += doc.Value
	}
	return sum
}
