package main

import cl "colorlib"

func NewBed(inCoordinate Coordinate, exitCoordinate Coordinate, stateKey string) Item {
	return Item{
		Name:         "Bed",
		ShortName:    "B",
		Description:  "A bed",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Walkable:     true,
		Color:        cl.RGB{R: 139, G: 69, B: 19},
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			if p.Room.State[stateKey].(bool) {
				p.Location = exitCoordinate
			} else {
				p.Location = inCoordinate
			}
			p.Room.State[stateKey] = !p.Room.State[stateKey].(bool)
		},
	}
}

func NewTV(stateKey string, c Coordinate, color cl.RGB) Item {
	return Item{
		Name:         "TV",
		ShortName:    "T",
		Description:  "A TV",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        color,
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			if p.Room.State[stateKey].(bool) {
				p.Room.Status = "Turned the TV off!"
				p.Room.Layout[c.Y][c.X] = NewTV(stateKey, c, tv_off)
			} else {
				p.Room.Status = "Turned the TV on!"
				p.Room.Layout[c.Y][c.X] = NewTV(stateKey, c, tv_on)
			}
			p.Room.State[stateKey] = !p.Room.State[stateKey].(bool)
		},
	}
}

func NewCouch(stateKey string, inCoordinate Coordinate, exitCoordinate Coordinate) Item {
	return Item{
		Name:         "Couch",
		ShortName:    "C",
		Description:  "A couch",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Walkable:     true,
		Carryable:    false,
		Color:        cl.RGB{R: 139, G: 69, B: 19},
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			if p.Room.State[stateKey].(bool) {
				p.Location = exitCoordinate
			} else {
				p.Location = inCoordinate
			}
			p.Room.State[stateKey] = !p.Room.State[stateKey].(bool)
		},
	}
}
