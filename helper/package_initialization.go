package helper

import "fmt"

// Package Initialization
// use init() function to let function run when package is imported (initialization)

var connection string

func init() {
	fmt.Println("Package initialization")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
