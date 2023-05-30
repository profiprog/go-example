package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of Procs: ", runtime.GOROOT())
}
