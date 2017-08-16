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

package search

import (
	"os"
	"path/filepath"
	"testing"
)

//
// Test constants
//

// Globals

const one = "foo is duplicated"

const (
	two   = 10
	three = "baz constant"
)

// single local
func dummy() {
	const four = 10
}

// multiple locals
func another() {
	const (
		five = "snth"
		six  = "foo is duplicated"
	)
}

var dontIncludeThis = `
blah
const blah blah
`

var want = []string{
	"one = \"foo is duplicated\"",
	"two   = 10",
	"three = \"baz constant\"",
	"four = 10",
	"five = \"snth\"",
	"six  = \"foo is duplicated\"",
}

func TestSearch(t *testing.T) {
	path := pathForThisFile()

	f, err := os.Open(path)
	if err != nil {
		t.Errorf("failed to open %s", path)
	}

	got := ExtractConsants(f)

	if len(got) != len(want) {
		t.Errorf("Fail: slice lengths do not match\n")
	} else {
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Fail.\n\t Got: %s\n\tWant: %s", got[i], want[i])
			}
		}
	}
}

// pathForThisFile returns the full path of this file.
func pathForThisFile() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(dir, "search_test.go")
}
