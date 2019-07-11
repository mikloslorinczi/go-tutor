package main

import "fmt"

const Num = 13

const (
	Word string = "Sajt"
	Bool bool   = false
)

var NumVar int

var (
	WordVar string
	BoolVar bool
)

func main() {
	fmt.Printf("Value of Num: %v, Word: %v, Bool: %v\n", Num, Word, Bool)

	fmt.Printf("Type of Num: %T, Word: %T, Bool: %T\n", Num, Word, Bool)

	fmt.Printf("Value of NumVar: %v, WordVar: %v, BoolVar: %v\n", NumVar, WordVar, BoolVar)

	NumVar = 3

	AnotherNumVar := Num + NumVar

	fmt.Printf("NumVar now: %v AnotherNumVar: %v\n", NumVar, AnotherNumVar)

	WordVar = "Torta"

	fmt.Println(Word + WordVar)
}
