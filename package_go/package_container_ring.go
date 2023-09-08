package packagego

import (
	"container/ring"
	"fmt"
	"strconv"
)

func MainPackageContainerRing() {
	// this is data structure which implement ring list, there is no end for the iteration, if the pointer reach at the end of the ring, it will go to the first again
	data := ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	// using built in function Do, you can iterate the ring without infinite loop
	// if you are using the for loop, it will loop the data infinately
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
