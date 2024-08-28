package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func input() (userInputs FormFormat) {

	// Create a bufio scanner to read user input.
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

func write() {
	//Read config file
	conf, err := os.ReadFile("conf.json")
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}

	//Converting JSON
	var data map[string]interface{}
	json.Unmarshal([]byte(conf), &data)
	var note = data["name_Gjrnl"].(string)

	// Open the file for writing
	jrnl, err := os.OpenFile(note, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0700)
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return
	}

	// Write the user input to the file
	for _, k := range input().Intrary {

		_, err = jrnl.WriteString(k + "\n")
		if err != nil {
			fmt.Printf("Error writing to the file: %v\n", err)
			return
		}
	}

	// Append time
	var now = time.Now()
	t := now.Format(time.ANSIC)
	_, err = jrnl.WriteString(t + "\n")
	if err != nil {

		fmt.Printf("Error writing timeStamp: %v\n", err)
	}

	defer jrnl.Close()

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
