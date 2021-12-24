package main

import "fmt"

// Golang by default when passing data using pass by value (duplicate value, so one object change not affect the source value), not pass by reference,
// Pointer can mae reference to data location in same memory, without duplicate already data
// so using pointer we can do pass by reference
// to use pointer using '&' operator

// if we want to create pointer with empty data (data initial not exist) use function new
func MainPointer() {
	// pass by value
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println()

	// pass by reference
	address3 := &address1
	address3.City = "Surabaya"
	fmt.Println(address1)
	fmt.Println(address3)
	fmt.Println()

	// Change data pointer, using * operator to make address 1 follow the changes
	*address3 = Address{"Malanga", "Jawa Timur", "Indonesia"}
	fmt.Println(address3)
	fmt.Println(address1)
	fmt.Println()

	// if using this way it will not change address 1 also
	address3 = &Address{"Bandung", "Jawa Barat", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address3)
	fmt.Println()

	// create pointer with empty data
	var address4 *Address = new(Address)
	fmt.Println(address4)
	address4.City = "Medan"
	fmt.Println(address4)

	// Pointer as function parameter, use * to mark we need pointer, and pass parameter with & operator
	address5 := Address{"Jakarta", "Jawa Barat", "Bangladesh"}
	changeCountryToIndonesia(&address5)
	fmt.Println(address5)
	fmt.Println()

	// method using pointer in the struct
	suha := Man{Name: "Suha"}
	suha.Married()
	fmt.Println(suha)
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

type Address struct {
	City, Province, Country string
}

// if you create method in struct, recommended to use pointer, so the memory will not be wasteful, by default if you create function for struct ,it is pass by value, so if you make any changes of data in struct it will not affect the source data struct outside the function

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println(man.Name)
}
