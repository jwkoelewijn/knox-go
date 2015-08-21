package main

import (
	"sort"
	"strings"
)

type DocumentList []Document

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

func NewDocumentList() *DocumentList {
	var docs DocumentList
	return &docs
}

func (d *DocumentList) Add(doc Document) *DocumentList {
	docs := append(*d, doc)
	return &docs
}

func (d DocumentList) Contains(doc Document) bool {
	for _, curDoc := range d {
		if curDoc == doc {
			return true
		}
	}
	return false
}

func (d *DocumentList) Cost() int {
	sum := 0
	for _, doc := range *d {
		sum += doc.Size
	}
	return sum
}

func (d DocumentList) String() string {
	var docStrings []string
	for _, doc := range d {
		docStrings = append(docStrings, doc.FormattedString())
	}
	return strings.Join(docStrings, "")
}

func (d DocumentList) Length() int {
	return len(d)
}
