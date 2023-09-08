package packagego

import (
	"fmt"
	"time"
)

func MainPackageTime() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	// to parse string time, the layout number should be different between others, for example month is 01, the day should not be 01, instead it should be 02
	parse, _ := time.Parse("2006-01-02", "2023-09-04")
	fmt.Println(parse)
}
