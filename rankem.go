package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Print("Items to rank:\n\n")

	r := []string{}
	for {
		item := readline()
		if item == "" {
			break
		}
		r = append(r, item)
	}
	check(scanner.Err())
	BinaryInsertionSort(r, LessPrompt)
	for i, s := range r {
		fmt.Printf("%d. %s\n", i+1, s)
	}
}

var scanner = bufio.NewScanner(os.Stdin)

func readline() string {
	if !scanner.Scan() {
		return ""
	}
	return scanner.Text()
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

func LessPrompt(a, b string) bool {
	p := promptui.Select{Label: "Which ranks higher", Items: []string{a, b}}
	choice, _, err := p.Run()
	check(err)
	return choice == 0
}

func BinaryInsertionSort(a []string, less func(a, b string) bool) {
	var hi, lo int

	for i := 1; i < len(a); i++ {
		lo = 0
		hi = i
		m := i / 2

		for {
			if less(a[i], a[m]) {
				hi = m
			} else {
				lo = m + 1
			}

			m = lo + ((hi - lo) / 2)
			if lo >= hi {
				break
			}
		}

		if m < i {
			tmp := a[i]
			copy(a[m+1:], a[m:i])
			a[m] = tmp
		}
	}
}
