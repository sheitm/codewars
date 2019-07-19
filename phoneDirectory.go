package main

import (
	"fmt"
	"regexp"
	"strings"
)

// https://www.codewars.com/kata/56baeae7022c16dd7400086e/train/go

var re = regexp.MustCompile("\\d{1,2}-\\d{3}-\\d{3}-\\d{4}")

type phoneDirEntry struct {
	phone   string
	name    string
	address string
}

func (p phoneDirEntry) toString() string {
	return fmt.Sprintf("Phone => %s, Name => %s, Address => %s", p.phone, p.name, p.address)
}

// Phone return phone entry
func Phone(dir, num string) string {
	parsed := parsePhoneDir(dir)
	if arr, ok := parsed[num]; ok {
		if len(arr) == 1 {
			return arr[0]
		}
		return fmt.Sprintf("Error => Too many people: %s", num)
	}
	return fmt.Sprintf("Error => Not found: %s", num)
}

func parsePhoneDir(dir string) map[string][]string {
	lines := strings.Split(dir, "\n")
	m := map[string][]string{}
	for _, line := range lines {
		entry := toEntry(line)
		arr := []string{}
		if a, ok := m[entry.phone]; ok {
			arr = a
		}
		arr = append(arr, entry.toString())
		m[entry.phone] = arr
	}
	return m
}

func toEntry(s string) phoneDirEntry {
	m := re.Find([]byte(s))
	return phoneDirEntry{
		phone: string(m),
	}
}
