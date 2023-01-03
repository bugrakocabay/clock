package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Give input")
		return
	}

	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)

	var (
		text  = args[0]
		size  = len(text)
		buf   = make([]byte, 0, size)
		start int
		stop  int
	)

	for i := 0; i < size; i++ {
		buf = append(buf, text[i])

		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			start = i + nlink
		}

		switch text[i] {
		case ' ', '\t', '\n':
			stop = i
		}
	}

	for i := 0; i < len(buf); i++ {
		if i >= start && i < stop {
			buf[i] = mask
		}
	}
	fmt.Println(string(buf))
}
