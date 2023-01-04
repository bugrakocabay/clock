package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	students := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "scorpius"},
	}
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("give me 1 input")
		return
	}

	query := args[0]
	arr, ok := students[query]
	if !ok {
		fmt.Printf("Sorry. I don't know anything about \"%s\".\n", query)
		return
	}

	sort.Strings(arr)
	fmt.Printf("~~~ %s students ~~~\n", query)
	for _, v := range arr {
		fmt.Printf("+ %s\n", v)
	}
}
