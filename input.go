package main

import (
	"bufio"
	"fmt"
	"os"
)

func input() (userInputs FormFormat) {

	// Create a bufio scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter multiple lines of input (press Ctrl+D when finished):")

	// //Read user input
	for scanner.Scan() {
		intrary := scanner.Text()
		userInputs.Intrary = append(userInputs.Intrary, intrary)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return

	}
	fmt.Println(userInputs.format())
	return userInputs
}

type Readeyto interface {
	format() []string
}

func write(k Readeyto) {

	// Open the file for writing
	jrnl, err := os.OpenFile("note", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0700)
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return
	}

	// Write the user input to the file
	_, err = jrnl.WriteString(string(k.format()))
	if err != nil {
		fmt.Printf("Error writing to the file: %v\n", err)
		return
	}

}
