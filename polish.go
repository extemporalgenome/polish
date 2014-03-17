// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package polish

import "strconv"

type Runner interface {
	Run([]float64) []float64
}

type Program []Runner

func (p Program) Run(stack []float64) []float64 {
	for _, s := range p {
		stack = s.Run(stack)
	}
	return stack
}

type Constant float64

func (c Constant) Run(stack []float64) []float64 {
	return append(stack, float64(c))
}

type BinOp func(float64, float64) float64

func (f BinOp) Run(stack []float64) []float64 {
	l := len(stack)
	stack, x, y := stack[:l-2], stack[l-2], stack[l-1]
	return append(stack, f(x, y))
}

func Add(x, y float64) float64 { return x + y }
func Sub(x, y float64) float64 { return x - y }
func Mul(x, y float64) float64 { return x * y }
func Div(x, y float64) float64 { return x / y }

var Dict = map[string]Runner{
	"+": BinOp(Add),
	"-": BinOp(Sub),
	"*": BinOp(Mul),
	"/": BinOp(Div),
}

func Parse(args []string) (p Program, err error) {
	p = make(Program, len(args))
	for i, arg := range args {
		word, ok := Dict[arg]
		if !ok {
			n, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				return nil, err
			}
			word = Constant(n)
		}
		p[i] = word
	}
	return p, nil
}
