package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me input")
		return
	}

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)
	)

	for i := 0; i < size; i++ {
		if i == 0 {
			buf = append(buf, text[i])
			continue
		}

		if i%40 == 0 {
			for j := i; ; j-- {
				if text[j] == ' ' {
					buf[j] = '\n'
					break
				}
			}
		}

		buf = append(buf, text[i])
	}

	fmt.Println(string(buf))
}
