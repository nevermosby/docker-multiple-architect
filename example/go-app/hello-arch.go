package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello World from %s!\n", runtime.GOARCH)
}