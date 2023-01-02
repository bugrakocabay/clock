package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {
	zero := [5]string{
		"▉▉▉",
		"▉ ▉",
		"▉ ▉",
		"▉ ▉",
		"▉▉▉",
	}
	one := [5]string{
		"▉▉ ",
		" ▉ ",
		" ▉ ",
		" ▉ ",
		"▉▉▉",
	}
	two := [5]string{
		"▉▉▉",
		"  ▉",
		"▉▉▉",
		"▉  ",
		"▉▉▉",
	}
	three := [5]string{
		"▉▉▉",
		"  ▉",
		"▉▉▉",
		"  ▉",
		"▉▉▉",
	}
	four := [5]string{
		"▉ ▉",
		"▉ ▉",
		"▉▉▉",
		"  ▉",
		"  ▉",
	}
	five := [5]string{
		"▉▉▉",
		"▉  ",
		"▉▉▉",
		"  ▉",
		"▉▉▉",
	}
	six := [5]string{
		"▉▉▉",
		"▉  ",
		"▉▉▉",
		"▉ ▉",
		"▉▉▉",
	}
	seven := [5]string{
		"▉▉▉",
		"  ▉",
		"  ▉",
		"  ▉",
		"  ▉",
	}
	eight := [5]string{
		"▉▉▉",
		"▉ ▉",
		"▉▉▉",
		"▉ ▉",
		"▉▉▉",
	}
	nine := [5]string{
		"▉▉▉",
		"▉ ▉",
		"▉▉▉",
		"  ▉",
		"▉▉▉",
	}
	type placeholder [5]string
	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}

				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}
