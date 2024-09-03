package main

import (
	"fmt"
	"os"
	"strings"
)

// helpF displays the help message and exits the program.

func helpF() {
	fmt.Println(`Markov Chain text generator.

	Usage:
	  markovchain [-w <N>] [-p <S>] [-l <N>]
	  markovchain --help
	
	Options:
	  --help  Show this screen.
	  -w N    Number of maximum words
	  -p S    Starting prefix
	  -l N    Prefix length`)

	os.Exit(0)
}

// PDefault creates a default starting prefix based on the specified length.

func PDefault(lflag int, generalText []string) string {
	if lflag < 1 {
		fmt.Println("Prefix number does not match")
		os.Exit(1)
	}
	if lflag > len(generalText) {
		fmt.Println("Prefix number not equal")
		os.Exit(1)
	}
	br := generalText
	res := br[:lflag]
	res1 := strings.Join(res, " ")
	return res1
}
