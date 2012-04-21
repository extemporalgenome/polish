// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package polish

import (
	"fmt"
	"strconv"
)

type Program struct {
	maxstack int
	steps    []Step
}

func (p Program) Execute() []float64 {
	stack := make([]float64, 0, p.maxstack)
	for _, s := range p.steps {
		stack = s.Execute(stack)
	}
	return stack
}

type Step interface {
	Execute([]float64) []float64
}

type BinOp func(float64, float64) float64

func (f BinOp) Execute(stack []float64) []float64 {
	l := len(stack)
	stack[l-2] = f(stack[l-2], stack[l-1])
	return stack[:l-1]
}

func Add(x, y float64) float64 { return x + y }
func Sub(x, y float64) float64 { return x - y }
func Mul(x, y float64) float64 { return x * y }
func Div(x, y float64) float64 { return x / y }

type Constant float64

func (c Constant) Execute(stack []float64) []float64 {
	return append(stack, float64(c))
}

// 3 2 4 + 1 - 6 * /

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
		step          Step
		size, maxsize int
		n             float64
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
			n, err = strconv.ParseFloat(str, 64)
			if err != nil {
				return
			} else {
				step, take, give = Constant(n), 0, 1
			}
		}
		size -= take
		if size < 0 {
			err = ErrStackUnderrun{i, str, size}
			return
		}
		size += give
		if size > maxsize {
			maxsize = size
		}
		steps[i] = step
	}
	p.maxstack = maxsize
	p.steps = steps
	return
}
