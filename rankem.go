package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Print("Items to rank:\n\n")

	r := newByRank()
	for {
		item := readline()
		if item == "" {
			break
		}
		r.Push(item)
	}
	check(scanner.Err())
	sort.Sort(r)
	for i, s := range r.arr {
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

type byRank struct {
	arr []string
	s   map[[2]string]bool
}

func newByRank() *byRank {
	return &byRank{s: make(map[[2]string]bool)}
}

func (r *byRank) Len() int {
	return len(r.arr)
}

func (r *byRank) Swap(i, j int) {
	r.arr[i], r.arr[j] = r.arr[j], r.arr[i]
}

func (r *byRank) Less(i, j int) bool {
	pair := [2]string{r.arr[i], r.arr[j]}
	if b, ok := r.s[pair]; ok {
		return b
	}

	oppPair := [2]string{r.arr[j], r.arr[i]}
	if b, ok := r.s[oppPair]; ok {
		return !b
	}

	p := promptui.Select{Label: "Which ranks higher", Items: pair[:]}
	choice, _, err := p.Run()
	check(err)
	r.s[pair] = choice == 0
	return choice == 0
}

func (r *byRank) Push(x string) {
	r.arr = append(r.arr, x)
}
