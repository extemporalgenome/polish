// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/extemporalgenome/polish"
)

type ArbitraryOp func([]float64) []float64

func (f ArbitraryOp) Run(stack []float64) []float64 {
	return f(stack)
}

func Dup(stack []float64) []float64 {
	return append(stack, stack[len(stack)-1])
}

func Rot(stack []float64) []float64 {
	i := len(stack) - 3
	v := stack[i]
	return append(append(stack[:i], stack[i+1:]...), v)
}

func Swap(stack []float64) []float64 {
	j := len(stack) - 1
	i := j - 1
	stack[i], stack[j] = stack[j], stack[i]
	return stack
}

func init() {
	polish.Dict["dup"] = ArbitraryOp(Dup)
	polish.Dict["rot"] = ArbitraryOp(Rot)
	polish.Dict["swap"] = ArbitraryOp(Swap)
}

func main() {
	input := "1 2 3 swap"
	program, err := polish.Parse(strings.Fields(input))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}
	fmt.Println(program.Run(nil))
}
