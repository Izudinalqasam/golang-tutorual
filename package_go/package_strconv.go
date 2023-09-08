package packagego

import (
	"fmt"
	"strconv"
)

func MainPackageStrconv() {
	// Package to convert from and to string of another data type
	boolean, err := strconv.ParseBool("True")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(boolean)
	}

	// 10 is bilangan biner, can be 2, 8, 16. and 64 is int64
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(number)
	}

	// this is same with above
	numberAtoi, err := strconv.Atoi("1000000")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(numberAtoi)
	}

	// convert number to string
	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)
}
