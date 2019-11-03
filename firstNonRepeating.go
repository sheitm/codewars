package main

import "strings"

func firstNonRepeating(str string) string {
	hits := map[string]encounteredChar{}
	for i := 0; i < len(str); i++ {
		c := string(str[i])
		cl := strings.ToLower(c)
		if _, ok := hits[cl]; ok {
			delete(hits, cl)
			hits[cl] = encounteredChar{hits:2}
			continue
		}
		ec := encounteredChar{
			pos:i,
			representation: c,
			hits:1,
		}
		hits[cl] = ec
	}

	lowest := encounteredChar{pos: len(str), representation:""}
	for _, ec := range hits  {
		if ec.hits < 2 && ec.pos < lowest.pos {
			lowest = ec
		}
	}
	return lowest.representation
}

type encounteredChar struct {
	pos int
	representation string
	hits int
}