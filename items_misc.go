package main

import (
	cl "colorlib"
	"fmt"
)

func NewKey(keySig string, identifier string) Item {
	return Item{
		Name:         fmt.Sprintf("Key - %s", identifier),
		ShortName:    "K",
		Description:  "A key",
		Weight:       1,
		Value:        0,
		Interactable: true,
		Carryable:    true,
		Color:        cl.RGB{R: 255, G: 215, B: 0},
		FontColor:    cl.RGB{R: 0, G: 0, B: 0},
		Walkable:     true,
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			p.Inventory = append(p.Inventory, i)
			p.Room.Layout[c.Y][c.X] = "."
		},
		State: map[string]interface{}{
			"keySig": keySig,
		},
	}
}

func NewBox(withKey bool, keySig string) Item {
	return Item{
		Name:         "Box",
		ShortName:    "B",
		Description:  "A box",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        cl.RGB{R: 0, G: 0, B: 0},
		FontColor:    cl.RGB{R: 255, G: 255, B: 255},
		Walkable:     false,
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			if withKey {
				if p.Room.State[fmt.Sprintf("%s_found", keySig)].(bool) {
					p.Room.Status = "Just an empty Box."
					return
				}
				p.Room.State[fmt.Sprintf("%s_found", keySig)] = true
				p.Room.Status = "You found another key!"

				key := NewKey(keySig, "B")
				p.Inventory = append(p.Inventory, &key)
			} else {
				p.Room.Status = "Just an empty Box."
			}

		},
	}
}

func NewLockBox(keySig string) Item {
	return Item{
		Name:         "LockBox",
		ShortName:    "L",
		Description:  "A lock box",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        cl.RGB{R: 0, G: 0, B: 0},
		FontColor:    cl.RGB{R: 255, G: 255, B: 255},
		Walkable:     false,
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			playerHasCorrectKey := false
			for _, item := range p.Inventory {
				if sig, ok := item.State["keySig"].(string); ok {
					if sig == keySig {
						playerHasCorrectKey = true
						break
					}
				}
			}

			if playerHasCorrectKey {
				if openedBefore, ok := p.Room.State["lock_box_opened"].(bool); ok {
					if openedBefore {
						p.Room.Status = "Just an empty box!"
						return
					}
				}
				p.Room.Status = "You opened the lock box! Julian was inside!"
				julian := NewBotPlayer("Julian")
				p.Inventory = append(p.Inventory, &julian)
				p.Room.State["lock_box_opened"] = true
			} else {
				p.Room.Status = "The lock box is locked!"
			}
		},
		State: map[string]interface{}{
			"keySig": keySig,
		},
	}
}

func NewBotPlayer(name string) Item {
	return Item{
		Name:         name,
		ShortName:    "P",
		Description:  "The player",
		Weight:       100,
		Value:        100,
		Interactable: false,
		Carryable:    true,
		Color:        cl.RGB{R: 0, G: 0, B: 0},
		FontColor:    cl.RGB{R: 255, G: 255, B: 255},
		Walkable:     false,
	}
}
