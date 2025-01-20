package main

import "fmt"

// Should use interface!!!
type Intray struct {
	Title string
	Body  string
	Date  string
}

func (i *Intray) LogUp(s []string) {

	switch {
	case s[0] == "f":
		fmt.Println("Titil: %s", s[0])
	}
}
