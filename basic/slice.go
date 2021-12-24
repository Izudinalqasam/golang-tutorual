package main

import "fmt"

// Slice is chunk of array, slice is data which accessing part or all array data
// Slice is same with array but the size or lenght of the slice can change
// Slice has 3 data type, namely pointer, length and capacity
// Pointer --> point first data in the array
// length --> how long data from slice
// capacity --> how capacity of the slice, where lenght can't more than capacity
// if you change data in array, data in the slice will change too, vice versa
// example slice --> array[low;high], array[4:7]
func MainSlice() {
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println("Lenght of slice is ", len(slice1))
	// Capacity will return total capacity from low index (pointer) of the slice to the last index of the array
	// so this will return 8 because month has 12, and slice low in the index 4 which is 5 th position, 12 - 5 = 7 + 1(low index included)
	fmt.Println("Capacity of slice is ", cap(slice1))

	months[5] = "diubah"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	// append will replace the high index of slice to new value
	// append will create new array if the capacity is full, if not it will use same array,
	// so in that case if we change value in slice or array, the value of both places will change,
	// if return new array using append function, we change value in the slice will not affect value in the array
	var slice3 = append(slice2, "suhah")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println((slice3))

	fmt.Println(slice2)

	// to create slice from scratch and it is saver, make(<tipedata>, <length>, <capacity>)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Suha"
	newSlice[1] = "Effendi"

	fmt.Println(newSlice, " and capacity is ", cap(newSlice))

	// Copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// Carefully to declare array, if we forget specify lenght or use ..., it will create slice not array, example below:
	intArray := [4]int{1, 2, 3, 4}
	intSlice := []int{1, 2, 3, 4}
	fmt.Println(intArray)
	fmt.Println(intSlice)
}
