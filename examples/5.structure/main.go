package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) Say() {
	fmt.Printf("Hello my name is: %s and I'm %d years old\n", p.Name, p.Age)
}

func (personCopy Person) Say2() {
	fmt.Printf("Hello my name is: %s and I'm %d years old\n", personCopy.Name, personCopy.Age)
}

func (p *Person) GetOlder() {
	p.Age++
}

func (personCopy Person) GetOlder2() {
	personCopy.Age++
}

func main() {
	Guy := new(Person)

	fmt.Println(Guy)

	fmt.Printf("%+v\n", Guy)

	AnotherGuy := Person{
		Name: "Sanyi",
		Age:  47,
	}

	fmt.Println(AnotherGuy)

	fmt.Printf("%+v\n", AnotherGuy)

	Guy.Age = 15
	Guy.Name = "Waldo"

	fmt.Println(Guy)

	fmt.Printf("%+v\n", Guy)

	PersonPointer := &Person{
		Name: "Kata",
		Age:  25,
	}

	AnotherPersonPointer := NewPerson("Lali", 33)

	Guy.Say2()
	AnotherGuy.Say2()
	PersonPointer.Say()
	AnotherPersonPointer.Say()

	Guy.GetOlder()
	Guy.Say()
}
