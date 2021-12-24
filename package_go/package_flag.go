package packagego

import (
	"flag"
	"fmt"
)

// Package flag consist of functionality for to parsing Command line argument (case from os.ARG())
func MainPackageFlag() {
	// Example of argument "-host=localhost -username=root -password=root"
	// first argument is flag name
	// second argument is default value
	// third argument is description
	host := flag.String("host", "localhost", "Put your host")
	username := flag.String("username", "root", "Put your username")
	password := flag.String("password", "", "Put your password")

	// Parsing argument process
	flag.Parse()

	// flag parsing above return pointer (location in memory if you try print that)
	// to access value of the pointer using *
	fmt.Println(*host, *username, *password)
}
