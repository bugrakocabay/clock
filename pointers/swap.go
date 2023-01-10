package main

import (
	"fmt"
	"strings"
)

func main() {
	first, second := 3.14, 2.99
	fmt.Printf("Before swapping %f , %f\n", first, second)
	swap(&first, &second)
	fmt.Printf("After swapping %f , %f\n", first, second)

	fmt.Println(strings.Repeat("-", 45))
	pa, pb := &first, &second
	swapAddr(pa, pb)
	fmt.Printf("first's address = %p, value = %f\n", &first, first)
	fmt.Printf("second's address = %p, value = %f\n", &second, second)
}

func swap(i, j *float64) {
	*i, *j = *j, *i
}

func swapAddr(a, b *float64) (*float64, *float64) {
	return b, a
}
