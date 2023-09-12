package main

import (
	"math/rand"
)

type Cell struct {
	char     rune
	prevChar rune
	xPos     int
	yPos     int
	speed    uint8
}

func (c *Cell) step() {
	c.yPos += int(c.speed)
	c.change()
}

func (c *Cell) check() {
	if c.yPos > screen_height+int(c.speed) {
		c.yPos = -int(c.speed)
	}
}

func (c *Cell) change() {
	c.prevChar = c.char
	c.char = randomChar()
}

func randomChar() rune {
	runes := []rune{'#', '%', '$', '@', '+', '&', '?', ':', '*', '!', 'o', '^', '-', '>'}
	return runes[rand.Intn(len(runes))]
}

func InitCells() []*Cell {
	cells := []*Cell{}
	for i := 0; i < number_of_cells; i++ {
		cells = append(cells, &Cell{
			char:     randomChar(),
			prevChar: '0',
			xPos:     i * size,
			yPos:     rand.Intn(screen_height),
			speed:    uint8(rand.Intn(4) + 1*size),
		})
	}
	return cells
}
