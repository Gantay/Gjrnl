package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Prompt the user for input
	fmt.Print("Enter some text: \n")

	// Create a bufio scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		userInput := scanner.Text() + "\n"

		// Prompt the user for the file name
		//fmt.Print("Enter the file name: ")
		var fileName = "test"
		//fmt.Scanln(&fileName)

		// Open the file for writing
		file, err := os.OpenFile("test", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Printf("Error opening the file: %v\n", err)
			return
		}
		defer file.Close()

		// Write the user input to the file
		_, err = file.WriteString(userInput)
		if err != nil {
			fmt.Printf("Error writing to the file: %v\n", err)
			return
		}

		fmt.Printf("User input has been written to %s\n", fileName)
	} else {
		fmt.Println("Failed to read user input.")
	}
}

// You can also use the `Write` function to write a byte slice:
// _, err = file.Write([]byte(newText))
// if err != nil {
//     log.Fatal(err)
// }
