package basic

import "fmt"

// Array is a collection of values of the same type.
// Array is fixed size
func MainArray() {
	var names [3]string
	class2 := [3]string{"A", "B", "C"}

	names[1] = "Suha"
	fmt.Println(names[1])

	fmt.Println(class2)
	fmt.Println("Lenght of array 2 is ", len(class2))
}
