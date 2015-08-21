package main

import "sort"

// By is the type of a "less" function that defines the ordering of its Document arguments
type By func(d1, d2 *Document) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function
func (by By) Sort(docs DocumentList) {
	ds := &documentSorter{
		documents: docs,
		by:        by, // The Sort method's receiver is the function (closure) that defined the sort order
	}
	sort.Sort(ds)
}

type documentSorter struct {
	documents DocumentList
	by        func(d1, d2 *Document) bool // function used in the Less method
}

// Len is part of the sort.Interface
func (d *documentSorter) Len() int {
	return d.documents.Length()
}

// Swap is part of the sort.Interface
func (d *documentSorter) Swap(i, j int) {
	d.documents[i], d.documents[j] = d.documents[j], d.documents[i]
}

// Less is part of the sort.Interface
func (d *documentSorter) Less(i, j int) bool {
	return d.by(&d.documents[i], &d.documents[j])
}

//
// Specialized sorting methods on DocumentList
//

func (d DocumentList) SortBySize() {
	bySize := func(d1, d2 *Document) bool {
		return d1.Size < d2.Size
	}
	By(bySize).Sort(d)
}

func (d DocumentList) SortByDescendingSize() {
	bySize := func(d1, d2 *Document) bool {
		return d1.Size > d2.Size
	}
	By(bySize).Sort(d)
}

func (d DocumentList) SortByValue() {
	byValue := func(d1, d2 *Document) bool {
		return d1.Value < d2.Value
	}
	By(byValue).Sort(d)
}

func (d DocumentList) SortByDescendingValue() {
	byValue := func(d1, d2 *Document) bool {
		return d1.Value > d2.Value
	}
	By(byValue).Sort(d)
}

func (d DocumentList) SortByRatio() {
	byRatio := func(d1, d2 *Document) bool {
		return d1.SecrecyRatio < d2.SecrecyRatio
	}
	By(byRatio).Sort(d)
}

func (d DocumentList) SortByDescendingRatio() {
	byRatio := func(d1, d2 *Document) bool {
		return d1.SecrecyRatio > d2.SecrecyRatio
	}
	By(byRatio).Sort(d)
}

func (d DocumentList) SortByName() {
	byName := func(d1, d2 *Document) bool {
		return d1.Name < d2.Name
	}
	By(byName).Sort(d)
}
