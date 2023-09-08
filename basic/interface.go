package basic

import (
	"errors"
	"fmt"
)

// Interface only contains function declaration, it doesn't contain any implementation
// in golang a data can be interpreted as interface if the data contain same declaration of the interface
// like below struct has getName() like interface has. so the struct can be interpreted as interface

// Interface empty is interface doesn't contain any function declaration, this make all data type be this empty interface type
func MainInterface() {
	suha := Person{"Suha"}
	sayHello(suha)

	cat := Animal{"pus"}
	sayHello(cat)

	whatEver := EmptyInterfaceFunction(2)
	fmt.Println(whatEver)

	// Error interface
	result, error := divideByZero(100, 0)
	if error != nil {
		fmt.Println("Error", error.Error())
	} else {
		fmt.Println("Result", result)
	}
}

type HasName interface {
	getName() string
}

type GetName interface {
	getName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

// Struct Person
type Person struct {
	name string
}

func (person Person) getName() string {
	return person.name
}

// Struct Animal
type Animal struct {
	name string
}

func (animal Animal) getName() string {
	return animal.name
}

// Empty Interface
func EmptyInterfaceFunction(i int) interface{} {
	if i == 1 {
		return 1
	} else {
		return "Whatever"
	}
}

// Error interface
func divideByZero(nilai int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Divider cannot be zero")
	} else {
		return nilai / divider, nil
	}
}
