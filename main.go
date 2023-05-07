package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"os"
)

const (
	fieldSize = 50
	alive     = 1
	dead      = 0
	rectSize  = 10
	dimension = fieldSize * rectSize
)

var (
	table = make([][]int, fieldSize)
	rects []sdl.Rect
)

func initTable() {
	for i := 0; i < 50; i++ {
		table[i] = make([]int, fieldSize)
		for j := 0; j < 50; j++ {
			table[i][j] = rand.Intn(1-0+1) + 0
		}
	}
}

func neighbourCount(row, col int) int {
	sum := 0
	sRow := row - 1
	sCol := col - 1
	for r := 0; r < 3; r++ {
		if sRow+r < 0 || sRow+r >= fieldSize {
			continue
		}
		for c := 0; c < 3; c++ {
			if sCol+c < 0 || sCol+c >= fieldSize {
				continue
			}
			if sCol+c == col && sRow+r == row {
				continue
			}
			sum += table[sRow+r][sCol+c]
		}
	}
	return sum
}

func update() {
	tmpTable := make([][]int, fieldSize)
	for r := 0; r < fieldSize; r++ {
		tmpTable[r] = make([]int, fieldSize)
		for c := 0; c < fieldSize; c++ {
			nCount := neighbourCount(r, c)
			if table[r][c] == dead && nCount == 3 {
				tmpTable[r][c] = alive
			} else if table[r][c] == alive && (nCount == 3 || nCount == 2) {
				tmpTable[r][c] = alive
			} else {
				tmpTable[r][c] = dead
			}

		}
	}
	table = tmpTable
}

func main() {
	window, err := sdl.CreateWindow("GameOfGo", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		dimension, dimension, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return
	}
	defer renderer.Destroy()
	active := true

	initTable()
	for active {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				active = false
			}
		}
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		for r := 0; r < fieldSize; r++ {
			for c := 0; c < fieldSize; c++ {
				if table[r][c] == alive {
					rects = append(rects, sdl.Rect{
						X: int32(r * rectSize),
						Y: int32(c * rectSize),
						W: rectSize,
						H: rectSize,
					})
				}
			}
		}
		renderer.SetDrawColor(0, 255, 255, 255)
		renderer.DrawRects(rects)
		renderer.Present()
		sdl.Delay(150)
		update()
		rects = []sdl.Rect{}
	}
}
