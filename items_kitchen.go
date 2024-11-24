package main

import (
	cl "colorlib"
	"math/rand"
)

func NewCounterTop() Item {
	return Item{
		Name:         "CounterTop",
		ShortName:    "C",
		Description:  "A countertop",
		Weight:       100,
		Value:        100,
		Interactable: false,
		Carryable:    false,
		Color:        cl.RGB{99, 99, 99},
	}
}

func NewCounterTopWithKey() Item {
	return Item{
		Name:         "CounterTop",
		ShortName:    "C",
		Description:  "A countertop",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        cl.RGB{99, 99, 99},
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			if p.Room.State["key_from_countertop_found"].(bool) {
				p.Room.Status = "Just an empty drawer."
				return
			}

			p.Room.Status = "You found a key!"
			key := NewKey("key_from_countertop", "A")
			p.Inventory = append(p.Inventory, &key)
			p.Room.State["key_from_countertop_found"] = true
		},
	}
}

func NewFridge() Item {
	return Item{
		Name:         "Fridge",
		ShortName:    "F",
		Description:  "A fridge",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        cl.RGB{255, 255, 255},
		FontColor:    cl.RGB{0, 0, 1},
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			p.Room.Status = "You opened the fridge!"
		},
	}
}

func NewTable() Item {
	return Item{
		Name:         "Table",
		ShortName:    "T",
		Description:  "A table",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    true,
		Color:        cl.RGB{139, 69, 19},
		Interaction:  PickUpItem("You picked up the table!"),
	}
}

func NewStove() Item {
	return Item{
		Name:         "Stove",
		ShortName:    "S",
		Description:  "A stove",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        cl.RGB{0, 0, 1},
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			ir := rand.Intn(3)
			switch ir {
			case 0:
				p.Room.Status = "Ouch! Its hot!"
			case 1:
				p.Room.Status = "Whoops, it was left on, I turned it off!"
			case 2:
				p.Room.Status = "This is making me hungry!"
			}
		},
	}
}

func NewWaterDispenser() Item {
	return Item{
		Name:         "WaterDispenser",
		ShortName:    "W",
		Description:  "A water dispenser",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        cl.RGB{0, 0, 255},
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			p.Room.Status = "You got a drink of water!"
		},
	}
}
