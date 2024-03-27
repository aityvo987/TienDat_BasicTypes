package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
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

// Chapter Method and interface

type IPAddr [4]byte
type MyReader struct {
	b byte
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

// Error

type ErrorSqrt float64

func (e ErrorSqrt) Error() string {
	//return fmt.Sprintf("Your number %v cannot be squared root",e)
	// only e in the parameter make infinite loop becausse it call the ErrorSqrt again
	return fmt.Sprintf("Your number %v cannot be squared root", float64(e))
}

func SqrtE(x float64) (float64, error) {
	if x < 0 {
		t := ErrorSqrt(x)
		return x, t
	}
	z := 1.0
	for ((z*z - x) / (2 * z)) != 0 {
		if math.Abs((z*z-x)/(2*z)) < 0.0000001 {
			return z, nil
		}
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

// Reader
func (m MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

// ROt13
type rot13Reader struct {
	i io.Reader
}

func rot13(b byte) byte {
	switch {
	case (b >= 65 && b <= 77) || (b >= 97 && b <= 109):
		b += 13
	case (b >= 78 && b <= 90) || (b >= 110 && b <= 122):
		b -= 13
	default:
	}
	return b
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, error := r13.i.Read(b)
	for i := 0; i <= n; i++ {
		b[i] = rot13(b[i])
	}
	return n, error
}

// Image
type Image struct {
	width, height int
	color         uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.color + uint8(x), i.color + uint8(y), 255, 255}
}

func main() {
	defer fmt.Println("-----------End of File-----------")
	fmt.Println("=======Slices========")
	pic.Show(Pic)

	fmt.Println("=======Maps========")
	wc.Test(WordCount)

	fmt.Println("=======Functions========")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
	fmt.Println("=======Stringer========")
	hosts := map[string]IPAddr{
		"myPC":      {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v : %v\n", name, ip)
	}
	fmt.Println("=======Error========")
	fmt.Println(SqrtE(2))
	fmt.Println(SqrtE(-2))
	fmt.Println("=======Read========")
	reader.Validate(MyReader{})

	fmt.Println("=======Rot13Reader========")
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println()

	fmt.Println("=======Image========")
	m := Image{100, 55, 255}
	pic.ShowImage(m)
}
