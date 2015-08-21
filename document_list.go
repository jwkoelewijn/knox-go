package main

import "strings"

type DocumentList []Document

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
	byCost := func(doc Document) int {
		return doc.Size
	}
	return Summer(byCost).Sum(d)
}

func (d *DocumentList) Value() int {
	byValue := func(doc Document) int {
		return doc.Value
	}
	return Summer(byValue).Sum(d)
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
