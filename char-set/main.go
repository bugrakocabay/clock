package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var start, stop int
	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}
	if start == 0 || stop == 0 {
		start = 'A'
		stop = 'Z'
	}

	fmt.Printf("%-10s %-10s %-10s %-10s\n%s\n", "literal", "dec", "hex", "encoded", strings.Repeat("-", 45))
	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(n))
	}
}
