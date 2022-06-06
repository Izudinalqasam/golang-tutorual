package basic

import "fmt"

func MainDataStructurMap() {
	person := map[string]string{
		"name":    "Suha",
		"address": "subang",
	}

	// add key value
	person["age"] = "20"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(len(person))

	delete(person, "age")
	fmt.Println(person)
}

// Map
// Map is a collection of key-value pairs. key is unique
// for key type data you can use any key (number, string or floating point)
