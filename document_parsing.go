package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

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

func ParseDocument(line string) (d Document, err error) {
	matchMap := re.matchMap(line)
	if len(matchMap) == 0 {
		return d, errors.New(fmt.Sprintf("Line %s is not a Document line", line))
	}
	name := matchMap["name"]
	size, err := strconv.Atoi(matchMap["size"])
	if err != nil {
		return d, err
	}
	value, err := strconv.Atoi(matchMap["value"])
	if err != nil {
		return d, err
	}
	return NewDocument(name, size, value), nil
}
