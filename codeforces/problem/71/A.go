package main

import (
	"fmt"
)

func solve(words []string) []string {
	var abbreviations []string
	for _, w := range words {
		wlen := len(w)
		if wlen > 10 {
			first := w[0:1]
			last := w[wlen-1:]

			abbreviation := fmt.Sprintf("%s%d%s", first, (wlen - 2), last)
			abbreviations = append(abbreviations, abbreviation)
		} else {
			abbreviations = append(abbreviations, w)
		}
	}
	return abbreviations
}

func read(n int) ([]string, error) {
	in := make([]string, n)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			return in[:i], err
		}
	}
	return in, nil
}

func printSlice(s []string) {
	if len(s) == 0 {
		return
	}
	fmt.Println(s[0])
	printSlice(s[1:])
}

func main() {
	var lines int
	var inputs []string

	fmt.Scan(&lines)
	if lines < 1 || lines > 100 {
		fmt.Println("Invalid number of lines")
		return
	}

	inputs, err := read(lines)
	if err != nil {
		return
	}

	printSlice(solve(inputs))
}

// Problem 71A: Way Too Long Words
// Time: 20m
