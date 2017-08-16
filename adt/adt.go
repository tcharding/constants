// Package adt provides an abstract data type to manipulate and view constants.
package adt

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type constant string
type value string
type path string

type ADT struct {
	m map[path]map[constant]value
}

func NewADT() *ADT {
	return &ADT{m: make(map[path]map[constant]value)}
}

// AddRawConstants adds constants to ADT, raw[i] is of form
//	"foo = \"abcd\""
//	"b   = 10"
func (adt *ADT) AddRawConstants(path string, raw []string) error {
	m := make(map[constant]value)
	for _, s := range raw {
		c, v, err := splitRawIntoConstantAndValue(s)
		if err != nil {
			log.Fatal(err)
			continue
		}
		_, ok := m[c]
		if ok {
			fmt.Errorf("constant already exists: %s\n", c)
		}
		m[c] = v
	}
	return adt.addFileConstants(path, m)
}

func (adt *ADT) addFileConstants(p string, m map[constant]value) error {
	_, ok := adt.m[path(p)]
	if ok {
		return fmt.Errorf("constants for %s already exist", p)
	}

	adt.m[path(p)] = m
	return nil
}

// Dump dumps all constants to stdout.
func (adt *ADT) Dump() {
	for _, m := range adt.m {
		for k, v := range m {
			fmt.Printf("%s = %s\n", k, v)
		}
	}
}

// Duplicates prints duplicate constants to stdout.
func (adt *ADT) Duplicates() {
	seen := make(map[value]int)

	for _, m := range adt.m {
		for _, v := range m {
			seen[v]++
		}
	}

	for v, num := range seen {
		if num > 1 {
			fmt.Printf("constant: %s\n", v)
			for p, m := range adt.m {
				for c, val := range m {
					if v == val {
						prettyPrint(c, p)
					}
				}
			}
			fmt.Println()
		}
	}
}

func splitRawIntoConstantAndValue(s string) (constant, value, error) {
	index := strings.Index(s, "=")
	if index == -1 {
		return "", "", fmt.Errorf("raw string ill formed: %s", s)
	}

	c := constant(strings.TrimSpace(s[:index]))
	v := value(strings.TrimSpace(s[index+1:]))
	return c, v, nil
}

func prettyPrint(c constant, p path) {
	dir, _ := os.Getwd()
	fmt.Printf("\t%s: %s\n", c, p[len(dir)+1:])
}
