package main

import (
	cl "colorlib"
)

type Item struct {
	// Item name
	Name string

	// ShortName string
	ShortName string

	// Item description
	Description string

	// Item weight
	Weight int

	// Item value
	Value int

	// Interactable
	Interactable bool

	// Carryable
	Carryable bool

	Color     cl.RGB
	FontColor cl.RGB

	Walkable bool

	Interaction func(g *game, p *Player, direction int, c Coordinate, i *Item)

	State map[string]interface{}
}
