package basic

import "fmt"

// it can change data type to another data type
func MainTypeAssertion() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	// panic if failed to change data type
	// resultInt := result.(int)
	// fmt.Println(resultInt)

	// elegant way to use switch statement
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}

func random() interface{} {
	return "OK"
}
