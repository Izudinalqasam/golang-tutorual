package basic

import "fmt"

func MainIfExpression() {
	class := "Kelas1"

	if class == "kelas2" {
		fmt.Println("kelas 2")
	} else if class == "kelas3" {
		fmt.Println("Kelas 3")
	} else {
		fmt.Println(class)
	}

	// if with short statement
	if length := len(class); length > 5 {
		fmt.Println("length of class is greater than 5")
	} else {
		fmt.Println("length of class is less than 5")
	}
}
