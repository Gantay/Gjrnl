package main

import ()

// Should use interface!!!
type Intray struct {
	Title       string
	Body        string
	Date        int32
	Attachments any
}

func (i *Intray) TheTitle() {}

func (i *Intray) TheBody() {}

func (i *Intray) TheDate() {}

func (i *Intray) TheAttachments() {}
