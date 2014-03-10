// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package polish

import (
	"strings"
	"testing"
)

func TestCorrectness(t *testing.T) {
	prog := "3 7 + 2 * 4 / 1.5 -"
	const expect = 3.5
	p, err := Parse(strings.Fields(prog))
	if err != nil {
		t.Fatal(err)
	}
	stack := p.Run(nil)
	if len(stack) != 1 || stack[0] != expect {
		t.Fatal("Expected", []float64{expect}, "instead of", stack)
	}
}

func TestCorrectness_prestack(t *testing.T) {
	progs := []string{"3 7", "+ 2", "*"}
	const expect = 20
	var stack []float64
	for _, prog := range progs {
		p, err := Parse(strings.Fields(prog))
		if err != nil {
			t.Fatal(err)
		}
		stack = p.Run(stack)
	}
	if len(stack) != 1 || stack[0] != expect {
		t.Fatal("Expected", []float64{expect}, "instead of", stack)
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
		p.Run(nil)
	}
}

func BenchmarkPureGo(b *testing.B) {
	var x float64 = 5.0
	for i := 0; i < b.N; i++ {
		_ = 1 / (2 - (3 * (4 + x)))
	}
}
