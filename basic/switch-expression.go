package main

import "fmt"

func MainSwithExpression() {
	name := "aldo"

	switch name {
	case "Bambang":
		fmt.Println("Hello Bambang!")
	case "Suha":
		fmt.Println("Hello Suha!")
	default:
		fmt.Println(name, "tidak ditemukan")
	}

	// Switch with short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("length of name is greater than 5")
	case false:
		fmt.Println("length of name is less than 5")
	}

	// Switch without condition predefined
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("length of name is greater than 5")
	case length < 5:
		fmt.Println("length of name is less than 5")
	default:
		fmt.Println("length is right")
	}
}
