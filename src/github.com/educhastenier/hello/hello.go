package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("42+13=", add(42, 13))
	a, b, c := swap("hello", "entire", "world")
	fmt.Println(a, b, c)
	fmt.Println(split(72))
}

func split(sum int) (dizaines, unites int) {
	dizaines = sum / 10
	unites = sum - (dizaines * 10)
	return
}

func swap(x, y, z string) (string, string, string) {
	return z, y, x
}

func add(x, y int) int {
	return x + y
}
