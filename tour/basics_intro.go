package main

// factored import
import (
	"fmt"
	"math"
	"math/cmplx" // imported must be used
	"math/rand"
	"time"
)

// this func can be above or below main
// https://go.dev/blog/declaration-syntax
func add(x int, y int) int {
	return x + y
}

// if same type for args then omit all except the last
func multiply(x, y int) int {
	return x * y
}

// swap and return tuple
func swap(x, y string) (string, string) {
	return y, x
}

// named returns example
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// naked return. to be only used in short functions
	return
}

// package level vars. type towards the end. (similar to args)\
// fine if not used
var dummy int // default is 0 (zero values)
var i, j int = 1, 2

// var blocks
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// constants
const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const text = "welcome"

	fmt.Println("Hello World")
	fmt.Println("The time is:", time.Now())
	// use rand.seed to see different values
	fmt.Println("My fav num is:", rand.Intn(10))

	// Pi is an exported name - notice the uppercase
	fmt.Println(math.Pi)

	fmt.Println("Add result:", add(42, 12))
	fmt.Println("Multiply result:", multiply(2, 4))

	// declared vars must be used
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	// local variable. need to be used.
	// initializers example
	// type omitted since variable will take the type of initializer
	var c, python, java = true, false, "no!"

	// short variable declaration (only inside a function)
	// value on right hand side is inferred (int in this case)
	k := 3

	fmt.Println(i, j, k, c, python, java)

	// print type and val
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// type conversion (explicit unlike C)
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	fmt.Println(text, Pi)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// will cause overflow
	// fmt.Println(needInt(Big))

}
