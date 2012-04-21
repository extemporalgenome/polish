// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package polish

import (
	"strings"
	"testing"
)

func TestCorrectness(t *testing.T) {
	instructions := "3 7 + 2 * 4 / 1.5 -"
	var expect float64 = 3.5
	t.Log(instructions)
	if p, err := Parse(strings.Fields(instructions)); err != nil {
		t.Fatal(err)
	} else if stack := p.Execute(); len(stack) != 1 {
		t.Fatal("Expected a stack of size 1")
	} else if stack[0] != expect {
		t.Fatal("Expected result to be", expect)
	}
}

func BenchmarkParse(b *testing.B) {
	in := strings.Fields("1 2 3 4 5 + * - /")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(in)
	}
}

func BenchmarkExec(b *testing.B) {
	instructions := "1 2 3 4 5 + * - /"
	p, _ := Parse(strings.Fields(instructions))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Execute()
	}
}
