// Copyright © 2017 Tobin C. Harding <me@tobin.cc>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package search searches Go files for constants.
package search

import (
	"bufio"
	"os"
	"strings"
)

// ExtractConsants extracts all constants from f in the form
// 	foo = "abc"
func ExtractConsants(f *os.File) []string {
	constants := []string{}
	inConstBlock := false
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if inConstBlock {
			if line == ")" {
				inConstBlock = false
			} else {
				verifyAndAdd(&constants, line)
			}
		} else {
			if strings.HasPrefix(line, "const") {
				if strings.HasPrefix(line, "const (") {
					inConstBlock = true
				} else {
					verifyAndAdd(&constants, line[6:]) // skip "const "
				}
			}
		}
	}
	return constants
}

func verifyAndAdd(slice *[]string, s string) {
	if isWellFormed(s) {
		*slice = append(*slice, s)
	}
}

func isWellFormed(s string) bool {
	index := strings.Index(s, "=")

	if index == -1 {
		return false
	}

	if index == 0 {
		return false
	}

	if s[index-1] == ':' {
		return false
	}

	return true
}
