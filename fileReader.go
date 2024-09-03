package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadFile reads all data from standard input and returns it as a slice of strings.

func ReadFile() []string {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error: reading file", err)
		os.Exit(1)
	}
	br := strings.Fields(string(data))
	if len(br) == 0 {
		fmt.Println("Error: file is empty")
		os.Exit(1)
	}
	return br
}

// Check verifies whether the input is coming from a file or terminal.

func Check() bool {
	file, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error: reading file", err)
	}
	if file.Mode()&os.ModeCharDevice != 0 {
		return false
	} else {
		return true
	}
}
