package main

import "fmt"

type Document struct {
	Name         string
	Size         int
	Value        int
	SecrecyRatio float64
}

func NewDocument(name string, size, value int) (d Document) {
	d.Name = name
	d.Size = size
	d.Value = value
	d.SecrecyRatio = float64(value) / float64(size)
	return
}

func (d Document) String() string {
	return fmt.Sprintf("%s %dKB %dS %.2fS/KB", d.Name, d.Size, d.Value, d.SecrecyRatio)
}

func (d Document) FormattedString() string {
	return fmt.Sprintf("\n%32s %5dKB %3dS (%.2fS/KB)", d.Name, d.Size, d.Value, d.SecrecyRatio)
}
