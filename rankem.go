package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Items to rank:\n\n")

	var items []string
	for {
		item := readline("")
		if item == "" {
			break
		}
		items = append(items, item)
	}
	checkErr()

	var r byRank
	for _, g := range items {
		heap.Push(&r, g)
		checkErr()
	}

	fmt.Println()
	for i := range r {
		fmt.Printf("%d. %s\n", i+1, r[i])
	}
}

var scanner = bufio.NewScanner(os.Stdin)

func readline(prompt string, args ...interface{}) string {
	fmt.Printf(prompt, args...)
	if !scanner.Scan() {
		return ""
	}
	return scanner.Text()
}

func checkErr() {
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

type byRank []string

func (r byRank) Len() int {
	return len(r)
}

func (r byRank) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r byRank) Less(i, j int) bool {
	return readline(`1. "%s" or 2. "%s"? `, r[i], r[j]) == "1"
}

func (r *byRank) Push(x interface{}) {
	*r = append(*r, x.(string))
}

func (r *byRank) Pop() interface{} {
	old := *r
	n := len(old)
	x := old[n-1]
	*r = old[0 : n-1]
	return x
}
