package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			continue
		}

		if IsPrime(num) {
			fmt.Printf("%d ", num)
		}
	}
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
