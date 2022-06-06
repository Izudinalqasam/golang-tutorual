package basic

import "fmt"

func MainMathOperation() {
	var result = 10 + 10
	fmt.Println(result)

	var (
		a = 10
		b = 10
	)
	c := a * b
	fmt.Println(c)

	// Augmented assignment
	var i = 10
	i += 10
	fmt.Println(i)

	// Unary Opearion
	var negative = -100
	var positive = +100
	fmt.Println(negative, positive)

	// Comparison Operation
	var name1 = "Suha"
	var name2 = "Suha"
	var value1 = 100
	var value2 = 200

	var resultComparision bool = name1 == name2
	fmt.Println(resultComparision)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
