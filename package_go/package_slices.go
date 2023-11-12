package packagego

import (
	"fmt"
	"slices"
)

func MainPackageSlices() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "suha"))
	fmt.Println(slices.Index(names, "suha"))
	fmt.Println(slices.Index(names, "Paul"))
}
