package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

// =======================================
//
//	Slices
func Pic(dx, dy int) [][]uint8 {
	a := make([]uint8, dx)
	b := make([][]uint8, dy)
	for i1 := range b {
		for i2 := range a {
			a[i2] = uint8((i1 + i2) / 2)
		}
		b[i1] = a
	}
	return b
}

// =======================================
//
//	Maps
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	x := strings.Fields(s)
	for i := 0; i < len(x); i++ {
		m[x[i]] += 1
	}
	return m
}

// =======================================
//
//	Functions
func fibonacci() func(int) int {
	result := 0
	return func(x int) int {
		result += x
		return result
	}
}
func main() {
	//=======================================
	//					Slices
	pic.Show(Pic)

	//=======================================
	//					Maps

	wc.Test(WordCount)

	//=======================================
	//					Functions
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
