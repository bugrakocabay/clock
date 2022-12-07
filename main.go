package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	usage := "[int] [-v]?"
	rand.Seed(time.Now().UnixMicro())
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println(usage)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(usage)
		return
	}
	if guess < 0 {
		fmt.Println(usage)
		return
	}
	/*
			randM1 := "Good job!"
			randM2 := "Impressive!"
			randM3 := "Pure luck!"
			randM4 := "Good shit booii!"

		randL1 := "get lost"
		randL2 := "ghahaha noob"
		randL3 := "go play with toys"
		randL4 := "u gonna cry??"
	*/
	var line string
	for turn := 0; turn < 8; turn++ {
		n := rand.Intn(11)
		if turn == 0 && n == guess {
			fmt.Println("WOOWW YOU WON ON FIRST TURN!!!!")
			return
		}

		if args[1] == "-v" {
			line += fmt.Sprintf("%d ", n)
		}

		if n == guess {
			fmt.Printf("%s YOU WON", line)
			return
		}
	}
	fmt.Printf("%s YOU lostt", line)
}
