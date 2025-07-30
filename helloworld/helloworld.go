package main

import (
	"fmt"
	"runtime"

	"golang.kevinchen-verkada/helloworld/formatting"
)

func main() {
	fmt.Println(formatting.FormatOutput(fmt.Sprintf("Go version: %s", runtime.Version()), true /* useBox */))
}
