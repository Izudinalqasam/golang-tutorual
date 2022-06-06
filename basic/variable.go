package basic

// Golang is case sensitive, uppercase and lowercase are different
import (
	"fmt"
	"time"

	"github.com/Izudinalqasam/golang-tutorial/helper"
)

// if we declare variable but not initializing it, it will be default value of the data type (int -> 0), if you don't want it you can give nil

// Private variable or function start with lowercase name, e.g sayHello()
// Public variable or function start with uppercase name, e.g SayHello()
func MainVariable() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga koma lima = ", 3.5)

	fmt.Println("trues = ", true)
	fmt.Println("falses = ", false)

	fmt.Println("Hello, World!")
	fmt.Println("Current time is", time.Now())

	fmt.Println("length of string ", len("Suha"))
	fmt.Println("get char in the string start from 0 index ", string("Suha"[2]))

	var class string
	class = "Go"
	fmt.Println(class)

	var class2 = "Golang"
	println(class2)

	class3 := "Go version"
	fmt.Println(class3)

	// Declaring multiple variable, you should declare variable each on new line
	var (
		firstClass = "kelas 1"
		lastClass  = "kelas terakhir"
	)
	fmt.Println(firstClass)
	fmt.Println(lastClass)

	const thirdClass = "Kelas 3"
	const thirdNumber = 3
	fmt.Println(thirdClass)
	fmt.Println(thirdNumber)

	// Declaring multiple constant, you should declare constant each on new line
	const (
		fourthName   = "Kelas 4"
		fourthNumber = 4
	)
	fmt.Println(fourthName)
	fmt.Println(fourthNumber)

	// Converting data
	nilai32 := 32768
	fmt.Println(int64(nilai32))
	fmt.Println(string("Hello"[3]))

	// Type Declaration
	type NoKTP string
	type Married bool

	var noKtpSuha NoKTP = "182931849128412"
	var marriedStatus Married = true
	fmt.Println(noKtpSuha)
	fmt.Println(marriedStatus)

	helper.SayHelloFromBot()
}

// Number type data
// Go only has two types of Number data (int and floating point)
// int has int8, int16, int32, int64 which not start from zero (start from negative number)
// for int type which start from zero you can use uint8. uint16, uint32, uint64
// for floating point has float32, float64, complex64, complex128

// Alias
// in golang there is alias
// byte = int8
// rune = int32
// int = int32 (based on OS. OS 64bit will be set int64)
// uint = uint32

// Bool type data
// you can declare boolean variable with keyword bool

// String type data
// you can declare string variable with keyword string
// you can use double quotes to declare variable string

// Variable
// in Go, variable is not dynamic type data but static type data
// variable is declared with keyword var <name variable> <type data>
// golang will notify you if there is a variable has been declared but never be used
// var keyword is optional, but you should initiate variable directly and use :=
// remember := only used when initializing variable, to assign variable with new value you can use =

// Constants
// constant is variable will not change after first time assign value,
// you can declare constant with keyword const <name constant> <type data>,
// but you can omit type data with assign value directly

// Type Declaration
// you can make custom type data from another type
// you can make type declaration by using keyword type <name type> <type data>
