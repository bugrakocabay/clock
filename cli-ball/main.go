package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {
	const (
		width  = 50
		height = 10
	)
	cellEmpty := ' '
	cellBall := 'O'
	var (
		cell      rune
		px, py    int
		maxFrames = 1200
		vx, vy    = 1, 1
		speed     = time.Second / 10
	)
	buf := make([]rune, 0, width*height)
	board := make([][]bool, width)
	screen.Clear()
	for column := range board {
		board[column] = make([]bool, height)
	}
	for i := 0; i < maxFrames; i++ {

		px += vx
		py += vy
		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}
		board[px][py] = true

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		screen.MoveTopLeft()
		time.Sleep(speed)
		fmt.Print(string(buf))

	}
}
