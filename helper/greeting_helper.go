package helper

import "fmt"

// This file is an example for usecase access another package in golang
var version = "1.0.0"      // private variable, can't be accesssed from outside module
var Application = "Golang" // public variable, can be accesssed from outside module

func SayHelloFromBot() {
	fmt.Println("Hello guys")
}

func sayGoodByeBot() {
	fmt.Println("Good bye guys")
}
