package main

import "fmt"

func main() { // tutorial on Go routines, by Manu
	channel := make(chan string)
	chDone := make(chan bool)

	go abcGenerator(channel, chDone)
	go printer(channel)

	fmt.Println("I am called first")

	<-chDone
}

func abcGenerator(ch chan string, chDone chan bool) {
	for i := byte('a'); i <= byte('z'); i++ {
		ch <- string(i) + "."
	}
	chDone <- true
}

func printer(ch chan string) {
	for i := range ch {
		print(i)
	}
}
