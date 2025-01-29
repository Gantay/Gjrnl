package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("starting...")
	if len(os.Args) >= 2 {
		switch {
		case os.Args[1] == "-on":
			Curator(os.Args[2])
		case os.Args[1] == "-from":
			log.Fatal("from arg")

		case os.Args[1] == "-to":
			log.Fatal("to arg")

		case os.Args[1] == "-n":
			log.Fatal("n arg")

		case os.Args[1] == "-contains":
			log.Fatal("contains arg")

		case os.Args[1] == "@":
			log.Fatal("tags")

		default:
			log.Fatal("Unknowin command")
		}
	} else {
		input()
	}

}
