package main

// Use Blank identifier if you want to access init package without access any fucntion inside the package, e.g  _ "fmt"
import (
	"fmt"
	"golang-tutorial/helper"
)

func mainImport() {
	helper.SayHelloFromBot()
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // can't access private variable

	fmt.Println(helper.GetDatabase())
}
