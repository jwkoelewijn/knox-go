package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Document struct {
	Name         string
	Size         int
	Value        int
	SecrecyRatio float64
}

const rePattern = `\A\s*list -- \|\s*(?P<name>((\w|!)+)\.(\w+))\s*((?P<size>\d*)KB)\s*(?P<value>\d*)\z`

type lineRegExp struct {
	*regexp.Regexp
}

var re = lineRegExp{regexp.MustCompile(rePattern)}

func (r *lineRegExp) matchMap(s string) map[string]string {
	captures := make(map[string]string)
	match := r.FindStringSubmatch(s)
	if match == nil {
		return captures
	}
	for i, name := range r.SubexpNames() {
		if i == 0 {
			continue
		}
		captures[name] = match[i]
	}
	return captures
}

func (d Document) String() string {
	return fmt.Sprintf("%s %dKB %dS %.2fS/KB", d.Name, d.Size, d.Value, d.SecrecyRatio)
}

func (d Document) FormattedString() string {
	return fmt.Sprintf("\n%32s %5dKB %3dS (%.2fS/KB)", d.Name, d.Size, d.Value, d.SecrecyRatio)
}

func (d *Document) Unmarshal(line string) error {
	matchMap := re.matchMap(line)
	if len(matchMap) == 0 {
		return errors.New(fmt.Sprintf("Line %s is not a Document line", line))
	}
	d.Name = matchMap["name"]
	var err error
	d.Size, err = strconv.Atoi(matchMap["size"])
	if err != nil {
		return err
	}
	d.Value, err = strconv.Atoi(matchMap["value"])
	if err != nil {
		return err
	}
	d.SecrecyRatio = float64(d.Value) / float64(d.Size)
	return nil
}
