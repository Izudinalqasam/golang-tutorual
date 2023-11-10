package packagego

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation error")
	ErrNotFound   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ErrValidation
	}

	if id != "suha" {
		return ErrNotFound
	}

	// Sukses
	return nil
}

func MainPackageError() {
	err := GetById("kurni")
	if err != nil {
		if errors.Is(err, ErrValidation) {
			fmt.Println(err.Error())
		} else if errors.Is(err, ErrNotFound) {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Unknown error")
		}
	}
}
