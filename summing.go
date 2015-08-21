package main

type Summer func(d Document) int

func (sum Summer) Sum(dl *DocumentList) int {
	s := 0
	for _, doc := range *dl {
		s += sum(doc)
	}
	return s
}
