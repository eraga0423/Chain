package main

import "flag"

// Command-line flags: 'w' for maximum number of words to generate,
// 'p' for the starting prefix,
// 'l' for the prefix length, and 'help' to display the help message.

var (
	w    = flag.Int("w", 100, "maximum number of words")
	p    = flag.String("p", "", "Starting prefix")
	l    = flag.Int("l", 2, "Prefix length")
	help = flag.Bool("help", false, "show this screen")
)
