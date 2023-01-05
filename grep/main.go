package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		if len(args) == 0 {
			fmt.Println(in.Text())
		} else {
			if strings.Contains(strings.ToLower(in.Text()), args[0]) {
				fmt.Println(in.Text())
			}
		}
	}
}
