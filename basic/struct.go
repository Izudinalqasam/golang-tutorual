package basic

import "fmt"

// Struct is a template data or prototype which is used to unified zero or more other data type in one unitied data
// Data in struct is saved in the field, so struct is a collection of fields
// in OOP is almost like Class
// Struct can't be used directly, but we can make object from it (like abstract class or interface in java)
// Struct can be used as parameter of the function,
// Struct can has a function, but in the code it is doesn't look like function inside a struct.
func MainStruct() {
	var suha Customer
	suha.name = "Suha"
	suha.address = "Indonesia"
	suha.age = 30

	fmt.Println(suha.name)
	fmt.Println(suha.address)
	fmt.Println(suha.age)

	// Another way to create struct instance instead of above
	joko := Customer{
		name:    "Joko",
		address: "Indoneisa",
		age:     50,
	}
	// or (below way is the same but if order, total of field is different, it will be error)
	budi := Customer{"budi", "Indonesia", 20}

	fmt.Println(joko)
	fmt.Println(budi)

	// call function Struct
	rully := Customer{name: "Rully"}
	rully.sayHello()
}

type Customer struct {
	name, address string
	age           int
}

// Example create a function for Struct
func (customer Customer) sayHello() {
	fmt.Println("Hello My Name is", customer.name)
}
