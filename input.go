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
	write()
	return userInputs

}

func banner() {
	asciiLogo := `
  ____      _     _     
 / ___|__ _| |__ | |__  
| |  _ / _' | '_ \| '_ \ 
| |_| | (_| | | | | | | |
 \____|\__,_|_| |_|_| |_|
`
	fmt.Println(asciiLogo)
}
