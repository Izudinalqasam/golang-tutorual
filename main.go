package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	packagego "github.com/Izudinalqasam/golang-tutorial/package_go"
)

func main() {
	packagego.MainPackageMath()

	arr := []string{"coba1", "coba2", "coba3"}
	arr2 := arr[1:]
	fmt.Println(arr[:2])
	fmt.Println(arr2)
	fmt.Println(arr[0])

	// packagego.MainPackageRegexp()
	// goroutine.MainGoroutineDownloadSize()
}

func handler() {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)

	go operation1(ctx)

	time.After(2 * time.Second)
	cancel()

	time.Sleep(2 * time.Second)
}

func operation1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err().Error())
			return
		default:
			fmt.Println("Do Something")
		}
	}
}

func stringBuilderExample() {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("Hello, ")
	strBuilder.WriteString("world")
	result := strBuilder.String()
	fmt.Println(result)
}
