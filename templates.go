package main

import (
	"fmt"
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

	// for testing. delete later
	fmt.Println(i.Title, i.Body, len(i.Body), i.Date)
}

func (i *Intray) TimeStamp() {

	date := time.Now().Format("2006/01/02")
	time := time.Now().Format("15:04:05 MST")
	dt := fmt.Sprintf("%s %s", date, time)
	i.Date = dt
}
