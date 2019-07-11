package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	StringCh = make(chan string)
)

func ReadAndPrint(ch chan string) {
	for msg := range ch {
		fmt.Println("Message received: ", msg)
		time.Sleep(time.Millisecond * 300)
	}
}

func SpammChannel(msg string, ch chan string) {
	for {
		n := rand.Intn(3) // n will be between 0 and 3
		time.Sleep(time.Duration(n) * time.Second)
		ch <- msg
	}
}

func FastSpamm(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "Spammm fast"
	}
}

func main() {

	boolCh := make(chan bool, 2)

	boolCh <- false
	boolCh <- true

	fmt.Println("Bool chanel says ", <-boolCh)
	fmt.Println("Bool chanel says ", <-boolCh)

	go ReadAndPrint(StringCh)

	// go SpammChannel("Spamm this", StringCh)

	// go SpammChannel("Spamm that", StringCh)

	FastSpamm(StringCh)

	close(StringCh)

	time.Sleep(time.Second * 1)

	time.Sleep(time.Second * 15)
}
