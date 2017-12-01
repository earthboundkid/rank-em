package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Print("Items to rank:\n\n")

	var items []string
	for {
		item := readline()
		if item == "" {
			break
		}
		items = append(items, item)
	}
	check(scanner.Err())

	r := make(byRank, 0, len(items))
	for _, g := range items {
		heap.Push(&r, g)
	}

	items = items[0:0]
	for i := 0; len(r) > 0; i++ {
		item := heap.Pop(&r).(string)
		items = append(items, item)
	}

	for i := range items {
		fmt.Printf("%d. %s\n", i+1, items[i])
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

type byRank []string

func (r byRank) Len() int {
	return len(r)
}

func (r byRank) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r byRank) Less(i, j int) bool {
	p := promptui.Select{Label: "Which ranks higher", Items: []string{r[i], r[j]}}
	choice, _, err := p.Run()
	check(err)
	return choice == 0
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
