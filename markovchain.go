package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// CreatorMarkovchain generates text based on the prefix map and specified parameters.

func CreatorMarkovchain(mapa map[string][]string, generalText []string, pflag string, lflag int, wflag int) []string {
	if pflag == "" || pflag == " " {
		pflag = PDefault(lflag, generalText)
	}
	if mapa == nil {
		fmt.Println("Empty text")
		os.Exit(1)
	}
	if len(strings.Fields(pflag)) < lflag {
		fmt.Println("Error: flags do not match")
		os.Exit(1)
	} else if len(strings.Fields(pflag)) > lflag {
		a := strings.Fields(pflag)
		pflag = strings.Join(a[(len(a)-lflag):], " ")
	}
	count := lflag
	txt := []string{}
	suffixes, ok := mapa[pflag]
	if !ok {
		fmt.Println("Error: prefix not found")
		os.Exit(1)
	}
	randomizer := rand.Intn(len(suffixes))
	splittext := strings.Fields(pflag)
	txt = append(splittext[1:], suffixes[randomizer])
	result := append(splittext, suffixes[randomizer])

	for {
		if wflag <= lflag {
			return result[:wflag]
		}
		count++
		if len(splittext) > lflag {
			result = result[:lflag]
			break
		}
		suffixes1, ok := mapa[strings.Join(txt, " ")]
		if !ok || count == wflag {
			break
		}
		randomizer1 := (rand.Intn(len(suffixes1)))
		txt = append(txt[1:], suffixes1[randomizer1])
		result = append(result, txt[lflag-1:]...)
	}
	return result
}

// CreateMap creates a map of prefixes and suffixes from the text for building a Markov chain.

func CreateMap(generalText []string, lflag int) map[string][]string {
	if generalText == nil {
		fmt.Println("Empty text")
		os.Exit(1)
	}
	mapa := make(map[string][]string)
	br := generalText
	lens := (len(br) - lflag) - 2
	for i := 0; i <= lens+1; i++ {
		pref := strings.Join(br[i:i+lflag], " ")
		mapa[pref] = append(mapa[pref], br[i+lflag])
	}
	return mapa
}
