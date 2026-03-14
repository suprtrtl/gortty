package main

import (
	"flag"
)

type Flags struct {
	screensaver *bool
}

func NewFlags() Flags {
	return Flags{
		screensaver: flag.Bool("screensaver", false, "Set Screensaver Mode or Not"),
	}
}

func (f Flags) Init() {
	flag.Parse()
}
