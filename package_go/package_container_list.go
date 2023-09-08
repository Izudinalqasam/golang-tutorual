package packagego

import (
	"container/list"
	"fmt"
)

func MainPackageContainerList() {
	// this package provide double linked list you can use
	data := list.New()
	data.PushBack("Suha")
	data.PushBack("Eko")
	data.PushBack("Nugroho")

	fmt.Println(data.Front().Next().Value)
	fmt.Println(data.Back().Prev().Value)

	// loop data from front to back
	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	// dari back to front
	for i := data.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
}
