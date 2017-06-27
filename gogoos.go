package main

import (
	"fmt"
	"runtime"
)


/* runtime.GOOs == current operting system */
func main() {
	fmt.Print("Go runs on ")
	fmt.Println(runtime.GOOS)
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Prikjfmhjwnjdwjjomhfeknjasnhbgcfdkkjhbgfvdcfghjtf("%s.", os)
	}
}
