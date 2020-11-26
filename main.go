//Package decalaration
package main

// Importing packages
// import "fmt"

// Also importing packages
import (
	"fmt"
	"math"
	"math/cmplx"
)

// Example of a function
func add(x int, y int) int {
	return x + y
}

// Another examaple of a func
func multiplu(x, y int) int {
	return x * y
}

// Also a function with name return so u dont need to specify them in the function with keyword return
func sub(a int, b int) (x, y int) {
	x = a - b
	y = b - a
	return
}

func main() {
	// type 1 initalization
	var a, b, c int = 1, 2, 3

	// type 2 initalization
	var ok, notOk bool
	ok = true
	notOk = false

	// type 3 initialization without var keyword
	k := "hello_there"

	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))
	fmt.Println(a, b, c, ok, notOk, k)

	// Data types
	// bool

	// string

	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr

	// int uint uintptr are 32bit on 32bit systems and 64 on 64bit systems

	// byte - alias for uint8

	// rune - alias for int32
	//        represents a Unicode code point

	// float32 float64

	// complex64 complex128

	// type 4 initialization
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("type: %T Value: %v\n", z, z)

	// conversion
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var p uint = uint(f)
	fmt.Println(x, f, p)

	// const declaration
	const pi = 3.14
	fmt.Println(pi)

}
