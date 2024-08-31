package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	w    = flag.Int("w", 100, "maximum number of words")
	p    = flag.String("p", "lya lya", "Starting prefix")
	l    = flag.Int("l", 2, "Prefix length")
	help = flag.Bool("help", false, "show this screen")
)

func CreateMap() map[string][]string {
	if len(strings.Split(*p, " ")) != *l {
		fmt.Println("Flags do not match")
		os.Exit(1)
	}

	data, err := io.ReadAll(os.Stdin)
	// fmt.Println(string(content))
	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}
	mapa := make(map[string][]string)
	br := strings.Fields(string(data))
	lens := (len(br) - *l) - 2

	for i := 0; i <= lens+1; i++ {
		pref := strings.Join(br[i:i+*l], " ")
		mapa[pref] = append(mapa[pref], br[i+*l])
	}
	return mapa
}

func CreatorMarkovchain(mapa map[string][]string) []string {
	count := *l
	txt := []string{}
	suffixes, _ := mapa[*p]
	randomizer := 0
	// (rand.Intn(len(suffixes)))
	splittext := strings.Split(*p, " ")
	txt = append(splittext[1:], suffixes[randomizer])
	result := append(splittext, suffixes[randomizer])

	for {
		count++
		if len(*p) >= *l {
			result = result[:*l]
			break
		}
		suffixes1, ok := mapa[strings.Join(txt, " ")]
		if !ok || count == *w {
			break
		}
		randomizer1 := 0
		// (rand.Intn(len(suffixes1)))
		txt = append(txt[1:], suffixes1[randomizer1])
		result = append(result, txt[*l-1:]...)

	}
	return result
}

func main() {
	flag.Parse()
	if len(os.Args) == 1 {
		a := CreateMap()
		fmt.Println(CreatorMarkovchain(a))

	}

	// CreateMap()
}
