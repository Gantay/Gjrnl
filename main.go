package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("starting...")
	switch {
	case os.Args[1] == "-r":
		Output()
	default:
		input()
	}

	//use os.Args to switch between input & output
	// bb()
	// CSS()

}
