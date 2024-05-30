package main

import (
	"fmt"
)

type FormFormat struct {
	Intrary []string
}

func JrnalForm(title string, intrary []string) FormFormat {
	I := FormFormat{

		Intrary: intrary,
	}

	fmt.Println(I)
	return I
}

func (i FormFormat) format() (title []string, body []string) {

	title = i.Intrary[:1]
	body = i.Intrary[1:]

	return title, body

}
