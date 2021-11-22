package main

// factored import
import (
	"fmt"
	"math" // imported must be used
	"math/rand"
	"time"
)

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

	sum := 0
	// init and post are optional
	// for ; sum <100; {}
	// or simply, for sum < 100
	// infinite loop: for {}
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(pow(3, 2, 10))
	switcher()
	defer_greeter()

	pointer_demo()
}
