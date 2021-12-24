package packagego

import (
	"fmt"
	"os"
)

// Package OS consist of functionality for accessing OS feature independently (can be used in every os)
// argument in index 0 if path compilation, so to access argument pass by you start from index 1
func MainPackageOS() {
	// Get arguments passed to the program from command line
	args := os.Args
	fmt.Println(args[1])
	fmt.Println(args[2])

	// Get current hostname of device
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", hostname)
	} else {
		fmt.Println("Error :", err)
	}
}
