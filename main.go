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
	var s []int
	return func(x int) int {
		var result int
		if x == 0 {
			result = 0
		} else if x == 1 {
			result = 1
		} else {
			result = s[x-1] + s[x-2]
		}

		s = append(s, result)
		return result
	}
}

func main() {

	fmt.Println("=======Slices========")
	pic.Show(Pic)

	fmt.Println("=======Maps========")
	wc.Test(WordCount)

	fmt.Println("=======Functions========")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
