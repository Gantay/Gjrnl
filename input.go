package main

import (
	"bufio"
	"fmt"
	"os"
)

func input() {

	scanner := bufio.NewScanner(os.Stdin)
	toBeSent := []string{}

	fmt.Println("Enter multiple lines of input (press Ctrl+D when finished):")

	var bb Intray

	// // //Read user input
	for scanner.Scan() {
		line := scanner.Text()
		toBeSent = append(toBeSent, line)
	}
	// how can i deal with empty spaces????????
	// me no know
	if len(toBeSent) == 0 {
		fmt.Println("no user input.")
		os.Exit(0)
	} else if toBeSent[0] == "\t" {
		fmt.Println("no Title.")
		os.Exit(0)
	}

	bb.TimeStamp()
	bb.InputIntray(toBeSent)

}
