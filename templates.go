package main

import "time"

type FormFormat struct {
	Title   string
	Intrary string
	Date    time.Time
}

func jrnalForm(title string, intrary string) FormFormat {
	I := FormFormat{
		Title:   title,
		Intrary: intrary,
	}
	return I
}