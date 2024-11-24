package main

import (
	cl "colorlib"
	"math/rand"
)

func NewToilet() Item {
	return Item{
		Name:         "Toilet",
		ShortName:    "T",
		Description:  "A toilet",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        cl.RGB{255, 255, 255},
		FontColor:    cl.RGB{0, 0, 1},
		Walkable:     false,
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			ir := rand.Intn(3)
			switch ir {
			case 0:
				p.Room.Status = "You used the toilet!"
			case 1:
				p.Room.Status = "You took a giant dump!"
			case 2:
				p.Room.Status = "You clogged it!"
			}
		},
	}
}

func NewShower() Item {
	return Item{
		Name:         "Shower",
		ShortName:    "S",
		Description:  "A shower",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        cl.RGB{255, 255, 255},
		FontColor:    cl.RGB{0, 0, 1},
		Walkable:     false,
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			ir := rand.Intn(3)
			switch ir {
			case 0:
				p.Room.Status = "Needs to be cleaned!"
			case 1:
				p.Room.Status = "Someone left the water on!"
			case 2:
				p.Room.Status = "You got wet!"
			}
		},
	}
}

func NewWasherDrier() Item {
	return Item{
		Name:         "Washer/Drier",
		ShortName:    "W",
		Description:  "A washer/drier",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        cl.RGB{255, 255, 255},
		FontColor:    cl.RGB{0, 0, 1},
		Walkable:     false,
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			ir := rand.Intn(3)
			switch ir {
			case 0:
				p.Room.Status = "You washed your clothes!"
			case 1:
				p.Room.Status = "You dried your clothes!"
			case 2:
				p.Room.Status = "You broke it!"
			}
		},
	}
}
