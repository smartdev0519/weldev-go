// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// fastlog2Table contains log2 approximations for 5 binary digits.
// This is used to implement fastlog2, which is used for heap sampling.

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

func main() {
	var buf bytes.Buffer

	fmt.Fprintln(&buf, "// Code generated by mkfastlog2table.go; DO NOT EDIT.")
	fmt.Fprintln(&buf, "// Run go generate from src/runtime to update.")
	fmt.Fprintln(&buf, "// See mkfastlog2table.go for comments.")
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, "package runtime")
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, "const fastlogNumBits =", fastlogNumBits)
	fmt.Fprintln(&buf)

	fmt.Fprintln(&buf, "var fastlog2Table = [1<<fastlogNumBits + 1]float64{")
	table := computeTable()
	for _, t := range table {
		fmt.Fprintf(&buf, "\t%v,\n", t)
	}
	fmt.Fprintln(&buf, "}")

	if err := ioutil.WriteFile("fastlog2table.go", buf.Bytes(), 0644); err != nil {
		log.Fatalln(err)
	}
}

const fastlogNumBits = 5

func computeTable() []float64 {
	fastlog2Table := make([]float64, 1<<fastlogNumBits+1)
	for i := 0; i <= (1 << fastlogNumBits); i++ {
		fastlog2Table[i] = math.Log2(1.0 + float64(i)/(1<<fastlogNumBits))
	}
	return fastlog2Table
}
