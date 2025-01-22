package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Should use interface!!!
type Intray struct {
	Title string
	Body  []string
	Date  string
}

func (i *Intray) InputIntray(s []string) {

	for index, value := range s {

		switch {
		case index == 0:
			i.Title = fmt.Sprintf("** %s **", value)
		default:
			i.Body = append(i.Body, value)
		}
	}
}

func (i *Intray) TimeStamp() {

	date := time.Now().Format("2006/01/02")
	time := time.Now().UTC().Format("15:04:05 MST")
	dt := fmt.Sprintf("%s %s", date, time)
	i.Date = dt
}

// should it return an error?????????
func (i *Intray) WriteIntray() {
	fBody := strings.Join(i.Body, "\n")
	// home, err := os.UserHomeDir()
	// if err != nil {
	// 	fmt.Errorf("couldn't get the user's home directory: %v", err)
	// }

	var nameOfJrnl string = "test.txt"
	jrnl, err := os.OpenFile(nameOfJrnl, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0700)
	if err != nil {
		fmt.Printf("couldn't write to jrnl. %v", err)
	}
	defer jrnl.Close()

	jrnl.Write([]byte("\n" + i.Title + "\n"))
	jrnl.Write([]byte("\n" + fBody + "\n"))
	jrnl.Write([]byte("\n" + i.Date + "\n"))
	jrnl.Write([]byte("\n" + "\n" + "\n"))

}
