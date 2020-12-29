package main

import (
	"flag"
	"fmt"
)

func main() {
	var flagvar int
	flag.IntVar(&flagvar, "p", 42, "port of the service to connect")
	flag.Parse()
	fmt.Println("flag has value", flagvar)
	deffer()
	b()
	fmt.Println("")
	fmt.Println(c())
}

func deffer() {
	i := 0
	defer fmt.Println("i at the time of defer:", i)
	i++
	fmt.Println(i)
	return
}

func b() {
	fmt.Print("Defer is reverse order: ")
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}
func c() (output int) {
	defer func() { output++ }()
	// should increment output after value is returned
	return 1
}
