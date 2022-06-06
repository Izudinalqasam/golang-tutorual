package basic

import "fmt"

func MainForLoop() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Counter is ", counter)
		counter++
	}

	// Another option 1
	for tapping := 1; tapping <= 5; tapping++ {
		fmt.Println("Tapping is", tapping)
	}

	// Another option 2
	mySlice := []string{"Eko", "nana", "sugeng"}
	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}

	// Another option 3, For Range
	for i, value := range mySlice {
		fmt.Println("Index ", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "suha"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, " = ", value)
	}

	// Break in for loop
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Event for loop break ", i)
	}

	// Continue in for loop
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Event for loop continue ", i)
		}
	}
}

// Underscore is used for variable not used in golang
