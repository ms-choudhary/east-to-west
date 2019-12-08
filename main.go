package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode"
)

var WesternNotation = []string{"a", "a+", "b", "c", "c+", "d", "d+", "e", "f", "f+", "g", "g+"}
var IndianNotation = []string{"s", "s+", "r", "r+", "g", "m", "m+", "p", "p+", "d", "d+", "n"}

func getIndex(note string, notations []string) int {
	for i, n := range notations {
		if n == note {
			return i
		}
	}
	return -1
}

func generateTranslationMaps(saNote string) (indianToWestern, westernToIndian map[string]string) {
	i := getIndex(saNote, WesternNotation)

	indianToWestern = map[string]string{}
	westernToIndian = map[string]string{}
	for _, indianNote := range IndianNotation {
		westernNote := WesternNotation[i]
		indianToWestern[indianNote] = westernNote
		westernToIndian[westernNote] = indianNote
		i = (i + 1) % len(WesternNotation)
	}
	return
}

func splitNote(s string) (note, duration string) {
	i := len(s) - 1

	for ; i >= 0; i-- {
		if !unicode.IsNumber(rune(s[i])) {
			break
		}
	}

	note = s[:i+1]
	duration = s[i+1:]
	return
}

func main() {

	from := flag.String("from", "indian", "indian or western")
	sa := flag.String("sa", "a", "which western note is sa")

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "required filename")
		os.Exit(1)
	}

	filename := flag.Args()[0]

	indianToWestern, westernToIndian := generateTranslationMaps(*sa)

	var m map[string]string
	if *from == "western" {
		m = westernToIndian
	} else {
		m = indianToWestern
	}

	infile, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer infile.Close()

	scanner := bufio.NewScanner(infile)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()

		n, d := splitNote(word)

		if _, ok := m[n]; ok {
			fmt.Printf("%s%s ", m[n], d)
		} else {
			fmt.Printf("%s ", word)
		}
	}
}
