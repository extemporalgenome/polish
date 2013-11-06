// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package polish

import (
	"fmt"
	"strconv"
)

type Program struct {
	steps    []Step
	maxstack int
}

func (p Program) Execute() []float64 {
	stack := make([]float64, 0, p.maxstack)
	for _, s := range p.steps {
		stack = s.Step(stack)
	}
	return stack
}

type Step interface {
	Step([]float64) []float64
}

type BinOp func(float64, float64) float64

func (f BinOp) Step(stack []float64) []float64 {
	l := len(stack)
	stack[l-2] = f(stack[l-2], stack[l-1])
	return stack[:l-1]
}

func Add(x, y float64) float64 { return x + y }
func Sub(x, y float64) float64 { return x - y }
func Mul(x, y float64) float64 { return x * y }
func Div(x, y float64) float64 { return x / y }

type Constant float64

func (c Constant) Step(stack []float64) []float64 {
	return append(stack, float64(c))
}

type ErrStackUnderrun struct {
	ArgNum   int
	Arg      string
	Overflow int
}

func (e ErrStackUnderrun) Error() string {
	return fmt.Sprintf("buffer underrun: argument [%d] %q underran the stack to %d", e.ArgNum, e.Arg, e.Overflow)
}

func Parse(args []string) (p Program, err error) {
	var (
		step     Step
		size     int
		maxstack int
	)
	steps := make([]Step, len(args))
	for i, str := range args {
		take, give := 2, 1
		switch str {
		case "+":
			step = BinOp(Add)
		case "-":
			step = BinOp(Sub)
		case "*":
			step = BinOp(Mul)
		case "/":
			step = BinOp(Div)
		default:
			n, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return p, err
			} else {
				step, take, give = Constant(n), 0, 1
			}
		}
		size -= take
		if size < 0 {
			return p, ErrStackUnderrun{i, str, size}
		}
		size += give
		if size > maxstack {
			maxstack = size
		}
		steps[i] = step
	}
	return Program{steps, maxstack}, nil
}
