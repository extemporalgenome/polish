// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/extemporalgenome/polish"
	"os"
)

func main() {
	program, err := polish.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}
	stack := program.Execute()
	if len(stack) == 1 {
		fmt.Println(stack[0])
	} else {
		fmt.Println(stack)
	}
}
