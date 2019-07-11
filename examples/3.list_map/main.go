package main

import "fmt"

func main() {
	fmt.Println("-- Arrays -------------------------------")
	Arrays()
	fmt.Println("-- Maps ---------------------------------")
	Maps()
}

func Arrays() {
	var Arr = [3]int{}

	fmt.Println("Empty Array:", Arr)

	Slice := make([]int, 5)

	fmt.Println("Empty Slice:", Slice)

	Slice[3] = 19

	fmt.Println("19 assigned to index 3:", Slice)

	Slice = append(Slice, 44)

	fmt.Println("Appeded with 44:", Slice)

	fmt.Printf("Length of slice %d\n", len(Slice))

	for index, value := range Slice {
		fmt.Printf("Array index: %d Value of element: %v\n", index, value)
	}
}

func Maps() {
	Intmap := map[string]int{}

	fmt.Printf("Empty intmap: %+v\n", Intmap)

	fmt.Printf("Length of intmap: %d\n", len(Intmap))

	Intmap["First"] = 27

	NewInt := Intmap["First"]

	NonExistantKey := Intmap["Second"]

	fmt.Printf("Empty intmap: %+v\n", Intmap)

	fmt.Printf("Length of intmap: %d\n", len(Intmap))

	fmt.Printf("Existing key: %v, Non existing key: %v\n", NewInt, NonExistantKey)

	val, ok := Intmap["First"]

	fmt.Printf("Is key OK: %t Value of key: %v\n", ok, val)

	Intmap["Third"] = 33
	Intmap["Cat"] = 34
	Intmap["Dog"] = 35

	delete(Intmap, "Cat")

	for key, value := range Intmap {
		fmt.Printf("Map Key: %s Map Value: %d\n", key, value)
	}
}
