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

	var bb Entry

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
	bb.WriteIntray()

}

func Curator(lookingFor string) {
	//Predefined jrnl!!!
	fmt.Println(lookingFor)
	var nameOfJrnl string = "test.txt"
	jrnl, err := os.Open(nameOfJrnl)
	if err != nil {
		fmt.Printf("couldn't open to jrnl. %v", err)
	}
	defer jrnl.Close()

	search := bufio.NewScanner(jrnl)
	// var lookingFor string
	for search.Scan() {
		if len(search.Text()) == 0 {
			continue
		} else {
			switch {
			case search.Text() == lookingFor:
				fmt.Printf("fount it!!!! [%s].", search.Text())
			default:
				fmt.Println(search.Text())
			}
		}

	}
}

// func Pre(lookingFor string) {
// 	switch {
// 	case lookingFor == "date", lookingFor == "date":
// 		return byDate == true
// 	}
// }
