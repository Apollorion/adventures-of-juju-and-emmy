package main

import (
	cl "colorlib"
	"fmt"
)

type Coordinate struct {
	X int
	Y int
}

type Player struct {
	// Player name
	Name string

	// Player description
	Description string

	// Player health
	Health int

	// Player strength
	Strength int

	// Player inventory
	Inventory []*Item

	// Player location
	Location Coordinate

	Room *Room

	HeldItem int
}

func NewPlayer() *Player {
	return &Player{
		Name:        "Player",
		Description: "The player",
		Health:      100,
		Strength:    10,
		Inventory:   []*Item{},
	}
}

func PrintDirection(tile interface{}, direction int, width int) string {

	var itemDescription string
	var interactiveItemDescription string
	var roomDescription string
	var interactiveRoomDescription string

	switch direction {
	case FORWARD:
		itemDescription = "There is a %s in front of you."
		interactiveItemDescription = "There is a %s[i] in front of you."
		roomDescription = "You are facing a %s."
		interactiveRoomDescription = "You are facing a %s[i]."
	case BACKWARD:
		itemDescription = "There is a %s behind you."
		interactiveItemDescription = "There is a %s[k] behind you."
		roomDescription = "Behind you is a %s."
		interactiveRoomDescription = "Behind you is a %s[k]."
	case LEFT:
		itemDescription = "There is a %s to your left."
		interactiveItemDescription = "There is a %s[j] to your left."
		roomDescription = "To your left is a %s."
		interactiveRoomDescription = "To your left is a %s[j]."
	case RIGHT:
		itemDescription = "There is a %s to your right."
		interactiveItemDescription = "There is a %s[l] to your right."
		roomDescription = "To your right is a %s."
		interactiveRoomDescription = "To your right is a %s[l]."
	}

	returnable := ""
	var ll cl.LogLine
	if t, ok := tile.(Item); ok {
		if t.Interactable {
			ll = cl.LogLine{
				Text:        fmt.Sprintf(interactiveItemDescription, t.Name),
				BgRGB:       t.Color,
				FontRGB:     t.FontColor,
				WithNewLine: true,
				WithPadding: true,
				Padding:     width,
			}
			returnable += ll.GetString()
		} else {
			ll = cl.LogLine{
				Text:        fmt.Sprintf(itemDescription, t.Name),
				BgRGB:       t.Color,
				FontRGB:     t.FontColor,
				WithNewLine: true,
				WithPadding: true,
				Padding:     width,
			}
			returnable += ll.GetString()
		}
	} else if t, ok := tile.(string); ok {
		switch t {
		case "D":
			ll = cl.LogLine{
				Text:        fmt.Sprintf(interactiveRoomDescription, "door"),
				BgRGB:       doorColor,
				WithNewLine: true,
				WithPadding: true,
				Padding:     width,
			}
			returnable += ll.GetString()
		case "W":
			ll = cl.LogLine{
				Text:        fmt.Sprintf(interactiveRoomDescription, "window"),
				BgRGB:       windowColor,
				WithNewLine: true,
				WithPadding: true,
				Padding:     width,
			}
			returnable += ll.GetString()
		case "#":
			ll = cl.LogLine{
				Text:        fmt.Sprintf(roomDescription, "wall"),
				BgRGB:       wallColor,
				WithNewLine: true,
				WithPadding: true,
				Padding:     width,
			}
			returnable += ll.GetString()
		}
	}

	return returnable
}

func (p *Player) Move(direction int) {

	attemptedX := p.Location.X
	attemptedY := p.Location.Y

	switch direction {
	case FORWARD:
		attemptedY--
	case BACKWARD:
		attemptedY++
	case LEFT:
		attemptedX--
	case RIGHT:
		attemptedX++
	}

	if attemptedY >= 0 && attemptedY < len(p.Room.Layout) {
		if attemptedX >= 0 && attemptedX < len(p.Room.Layout[attemptedY]) {
			thisTileL := p.Room.Layout[p.Location.Y][p.Location.X]
			nextTileL := p.Room.Layout[attemptedY][attemptedX]
			if thisTile, ok := thisTileL.(Item); ok {
				if nextTile, ok2 := nextTileL.(Item); ok2 {
					if nextTile.Walkable && thisTile.Name == nextTile.Name {
						p.Location.Y = attemptedY
						p.Location.X = attemptedX
					}
				}
			} else if t, ok := nextTileL.(string); ok {
				if t == "." {
					p.Location.Y = attemptedY
					p.Location.X = attemptedX
				}
			}
		}
	}
}

func (p *Player) Interact(g *game, direction int) {

	tileUp := p.Room.Layout[p.Location.Y-1][p.Location.X]
	tileDown := p.Room.Layout[p.Location.Y+1][p.Location.X]
	tileLeft := p.Room.Layout[p.Location.Y][p.Location.X-1]
	tileRight := p.Room.Layout[p.Location.Y][p.Location.X+1]

	if direction == EXIT {
		thisTile := p.Room.Layout[p.Location.Y][p.Location.X]
		if t, ok := thisTile.(Item); ok {
			if t.Interactable {
				t.Interaction(g, p, direction, p.Location, &t)
			}
		}
	}

	if direction == FORWARD {
		if t, ok := tileUp.(Item); ok {
			if t.Interactable {
				t.Interaction(g, p, direction, Coordinate{X: p.Location.X, Y: p.Location.Y - 1}, &t)
			}
		}
	}

	if direction == BACKWARD {
		if t, ok := tileDown.(Item); ok {
			if t.Interactable {
				t.Interaction(g, p, direction, Coordinate{X: p.Location.X, Y: p.Location.Y + 1}, &t)
			}
		}
	}

	if direction == LEFT {
		if t, ok := tileLeft.(Item); ok {
			if t.Interactable {
				t.Interaction(g, p, direction, Coordinate{X: p.Location.X - 1, Y: p.Location.Y}, &t)
			}
		}
	}

	if direction == RIGHT {
		if t, ok := tileRight.(Item); ok {
			if t.Interactable {
				t.Interaction(g, p, direction, Coordinate{X: p.Location.X + 1, Y: p.Location.Y}, &t)
			}
		}
	}
}

func (p *Player) PlaceItem(direction int) {

	if p.HeldItem == 0 {
		p.Room.Status = "Youre not holding anything."
		return
	}

	tileUp := p.Room.Layout[p.Location.Y-1][p.Location.X]
	tileDown := p.Room.Layout[p.Location.Y+1][p.Location.X]
	tileLeft := p.Room.Layout[p.Location.Y][p.Location.X-1]
	tileRight := p.Room.Layout[p.Location.Y][p.Location.X+1]

	if direction == FORWARD {
		if t, ok := tileUp.(string); ok {
			if t == "." {
				p.Room.Layout[p.Location.Y-1][p.Location.X] = *p.Inventory[p.HeldItem-1]
				p.Inventory = RemoveFromInventory(p.Inventory, p.HeldItem-1)
				p.HeldItem = 0
			}
		}
	} else if direction == BACKWARD {
		if t, ok := tileDown.(string); ok {
			if t == "." {
				p.Room.Layout[p.Location.Y+1][p.Location.X] = *p.Inventory[p.HeldItem-1]
				p.Inventory = RemoveFromInventory(p.Inventory, p.HeldItem-1)
				p.HeldItem = 0
			}
		}
	} else if direction == LEFT {
		if t, ok := tileLeft.(string); ok {
			if t == "." {
				p.Room.Layout[p.Location.Y][p.Location.X-1] = *p.Inventory[p.HeldItem-1]
				p.Inventory = RemoveFromInventory(p.Inventory, p.HeldItem-1)
				p.HeldItem = 0
			}
		}
	} else if direction == RIGHT {
		if t, ok := tileRight.(string); ok {
			if t == "." {
				p.Room.Layout[p.Location.Y][p.Location.X+1] = *p.Inventory[p.HeldItem-1]
				p.Inventory = RemoveFromInventory(p.Inventory, p.HeldItem-1)
				p.HeldItem = 0
			}
		}
	}
}
