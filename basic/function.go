package basic

import "fmt"

// Function is first class citizen, means it can be data type and be saved in variable
func MainFucntion() {
	sayHelloTo("Bambang", "Suha")

	result := getHello("Bambang")
	fmt.Println("Result function", result)

	firstName, lastName := getFullName()
	fmt.Println("First name", firstName, " <==>", "Last Name", lastName)

	// ignore return value
	nickName, _, _ := getFullNameWithNick()
	fmt.Println("Nick name", nickName)

	// Named return value
	concileName, _ := getFullNameWithNamedReturn()
	fmt.Println("Concile name", concileName)

	// variadic Function
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Total", total)

	slice := []int{10, 20, 30, 40, 50}
	totalInSlice := sumAll(slice...)
	fmt.Println("Total in Slice", totalInSlice)

	// Save function in variable
	goodBye := getGoodBye
	fmt.Println(goodBye("Bambang"))

	// Function as parameter
	sayHelloWithFilter("asuk", spamFilter)

	// Anonymous function (2)
	blackList := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blackList)

	// Anonymous function (3)
	registerUser("admin", func(name string) bool {
		return name == "admin"
	})
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

// Function return multiple value
func getFullName() (string, string) {
	return "Bambang", "Suha"
}

func getFullNameWithNick() (string, string, string) {
	return "Bambang", "Suha", "Bamsu"
}

// Named Return value
func getFullNameWithNamedReturn() (firstName, lastName string) {
	firstName = "Bambang"
	lastName = "suha"
	return
}

// Variadic Function
// Use last parameter as varags (can receive multiple input value)
// using varargs only work in the last position, can't be used in the middle or front of function parameter
func sumAll(numbers ...int) (total int) {
	for _, number := range numbers {
		total += number
	}
	return
}

func getGoodBye(name string) string {
	return "Good bye " + name
}

// Function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func sayHelloWithFilterType(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "asuk" {
		return "..."
	} else {
		return name
	}
}

type Filter func(string) string

type Filter2 func(string) string

// Anonymous function (1)
type Backlist func(string) bool

func registerUser(name string, blacklist Backlist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("wellcome ", name)
	}
}

func doingFilter(input string, filter Filter2) string {
	return filter(input)
}
