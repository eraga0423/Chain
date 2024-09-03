package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// The main function of the program, processing flags and executing core functions.

func main() {
	flag.Parse()
	p := *p
	l := *l
	w := *w
	if *help {
		helpF()
	}

	if l < 0 || l > 5 {
		fmt.Println("Error: length of prefix can be from 0 to 5")
		os.Exit(1)
	}
	if w <= 0 || w > 10000 {
		fmt.Println("Error: number of words must be in between 1 and 10000")
		os.Exit(1)
	}

	if len(os.Args) == 0 || !Check() {
		fmt.Println("Error: no input text")
		os.Exit(1)
	}
	generalText := ReadFile()
	createMap := CreateMap(generalText, l)
	result := strings.Join(CreatorMarkovchain(createMap, generalText, p, l, w), " ")
	fmt.Println(result)
}
