package main

import "fmt"

var (
	StringCh = make(chan string)
)

func ReadAndPrint(ch chan string) {
	for msg := range ch {
		fmt.Println("Message received: ", msg)
	}
}

func main() {
	go ReadAndPrint(StringCh)

	StringCh <- "First"
	StringCh <- "Second"
	StringCh <- "Third"
	StringCh <- "Forth"
}
