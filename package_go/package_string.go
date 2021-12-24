package packagego

import (
	"fmt"
	"strings"
)

// this package consists of functionality to manipulate strings
func MainPackageString() {
	fmt.Println(strings.Contains("Suha Hidayat", "Hidayat"))
	fmt.Println(strings.Split("Suha Hidayat", " "))
	fmt.Println(strings.ToLower("Suha Hidayat"))
	fmt.Println(strings.ToUpper("Suha Hidayat"))
	fmt.Println(strings.Trim("    suha Hidayat    ", " "))
	fmt.Println(strings.ReplaceAll("suha suha suha hidayat", "suha", "budi"))
}
