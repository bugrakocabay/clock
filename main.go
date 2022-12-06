package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMicro())
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Give input")
	}

	guess, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Give an integer")
	}
	if guess < 0 {
		fmt.Println("Give an positive integer")
	}
	randM1 := "Good job!"
	randM2 := "Impressive!"
	randM3 := "Pure luck!"
	randM4 := "Good shit booii!"

	randL1 := "get lost"
	randL2 := "ghahaha noob"
	randL3 := "go play with toys"
	randL4 := "u gonna cry??"

	for turn := 0; turn < 3; turn++ {
		n := rand.Intn(guess + 1)
		if turn == 0 && n == guess {
			fmt.Println("WOOWW YOU WON ON FIRST TURN!!!!")
			return
		}

		if n == guess {
			switch rand.Intn(4) {
			case 0:
				fmt.Println(randM1)
			case 2:
				fmt.Println(randM2)
			case 3:
				fmt.Println(randM3)
			case 1:
				fmt.Println(randM4)
			}
			return
		}
	}

	switch rand.Intn(4) {
	case 0:
		fmt.Println(randL1)
	case 1:
		fmt.Println(randL2)
	case 2:
		fmt.Println(randL3)
	case 3:
		fmt.Println(randL4)
	}
}
