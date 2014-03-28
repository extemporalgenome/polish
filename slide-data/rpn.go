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

func main() {
	input := "2 3 4 + *"
	program, err := polish.Parse(strings.Fields(input))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}
	fmt.Println(program.Run(nil))
}
