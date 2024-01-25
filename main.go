package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	//create an empty slice to store inputs
	var inputs []string

	var now = time.Now()

	// Create a bufio scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user for input
	fmt.Println("Enter multiple lines of input (press Ctrl+D when finished):")

	//Read user input
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

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
	Gjrnl, err := os.OpenFile(note, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return
	}

	// Write the user input to the file

	for _, k := range inputs {
		_, err = Gjrnl.WriteString(k + "\n")
		if err != nil {
			fmt.Printf("Error writing to the file: %v\n", err)
			return
		}
	}
	//Append time
	t := now.Format(time.ANSIC)
	_, err = Gjrnl.WriteString(t + "\n")
	if err != nil {
		fmt.Printf("Error writing timeStamp: %v\n", err)
	}

	defer Gjrnl.Close()

}
