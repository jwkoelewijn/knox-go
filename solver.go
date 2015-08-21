package main

import (
	"fmt"
	"log"
	"sync"
)

func (d DocumentList) FindBestSolution(bandwidth int) Solution {
	channel := make(chan Solution, 10)
	bestSolutionChannel := make(chan Solution)
	var wg sync.WaitGroup
	wg.Add(len(d))

	go func(inputChannel <-chan Solution, outputChannel chan<- Solution) {
		var solution Solution
		maxValue := 0
		for s := range inputChannel {
			if val := s.Value(); val > maxValue {
				// no need for all the other stuff
				solution = NewSolution(s.Documents, DocumentList{}, s.Bandwidth)
				maxValue = val
			}
		}
		outputChannel <- solution
	}(channel, bestSolutionChannel)

	for i, _ := range d {
		go combinations(d, i, bandwidth, channel, &wg)
	}

	wg.Wait()
	close(channel)
	solution := <-bestSolutionChannel
	log.Printf("Best solution %+v", solution)
	return solution
}

func combinations(d DocumentList, k, bandwidth int, solutionChannel chan<- Solution, wg *sync.WaitGroup) error {
	defer wg.Done()
	pool := d
	n := len(pool)

	if k > n {
		return fmt.Errorf("Cannot create combinations of length %d when the slice has only length %d", k, n)
	}

	indices := make([]int, k)
	for i := range indices {
		indices[i] = i
	}

	result := make([]Document, k)
	for i, el := range indices {
		result[i] = pool[el]
	}
	docList := DocumentList(result)

	if docList.Cost() < bandwidth {
		handleResult(docList, d, bandwidth, solutionChannel)
	}

	for {
		i := k - 1
		for ; i >= 0 && indices[i] == i+n-k; i -= 1 {
		}

		if i < 0 {
			handleResult([]Document{}, d, bandwidth, solutionChannel)
			return nil
		}

		indices[i] += 1
		for j := i + 1; j < k; j += 1 {
			indices[j] = indices[j-1] + 1
		}

		for ; i < len(indices); i += 1 {
			result[i] = pool[indices[i]]
		}
		docList := DocumentList(result)

		if docList.Cost() < bandwidth {
			handleResult(docList, d, bandwidth, solutionChannel)
		}
	}
	return nil
}

func handleResult(result, d DocumentList, bandwidth int, solutionChannel chan<- Solution) {
	sol := NewSolution(result, d, bandwidth)
	maximizeSolution(&sol)
	solutionChannel <- sol
}

func maximizeSolution(sol *Solution) {
	descendingSize := func(d1, d2 *Document) bool {
		return d1.Size > d2.Size
	}
	By(descendingSize).Sort(sol.Pool)

	for sol.Cost < sol.Bandwidth {
		newDoc, err := findFittingDocument(sol.Documents, sol.Pool, sol.Bandwidth-sol.Cost)
		if err != nil {
			break
		}
		sol.Documents = append(sol.Documents, newDoc)
		sol.Cost += newDoc.Size
	}
}

func findFittingDocument(docs, pool DocumentList, maxSize int) (d Document, err error) {
	for _, d := range pool {
		if d.Size < maxSize && !docs.Contains(d) {
			return d, nil
		}
	}
	return d, fmt.Errorf("Could not find fitting document")
}
