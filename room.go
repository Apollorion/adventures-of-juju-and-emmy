package main

import (
	cl "colorlib"
	"fmt"
)

type Room struct {
	// Room name
	Name string

	// Room description
	Description string

	// Room layout
	// # = wall
	// D = door
	// W = window
	// . = floor
	// Type(Item) = Item
	// N = No Render
	Layout [][]interface{}

	State map[string]interface{}

	StartingPosition Coordinate

	OtherRooms map[string]*Room

	Status string
}

func (p *Player) PrintRoom(width int) string {
	returnable := ""
	for y := 0; y < len(p.Room.Layout); y++ {
		row := ""
		for x := 0; x < len(p.Room.Layout[y]); x++ {
			if x == p.Location.X && y == p.Location.Y {
				row += player
			} else {
				tile := p.Room.Layout[y][x]
				if t, ok := tile.(Item); ok {
					ll := cl.LogLine{
						Text:        fmt.Sprintf(" %s ", t.ShortName),
						BgRGB:       t.Color,
						FontRGB:     t.FontColor,
						WithNewLine: false,
						WithPadding: false,
					}
					row += ll.GetString()
				} else if t, ok := tile.(string); ok {
					switch t {
					case "#":
						row += wall
					case "D":
						row += door
					case "W":
						row += window
					case ".":
						row += floor
					case "N":
						row += fmt.Sprintf("   ")

					}
				}
			}
		}
		returnable += cl.PadWithOverride(fmt.Sprintf("%s\n", row), width, (len(p.Room.Layout[y]) * 3))
	}

	return returnable
}

func (p *Player) DescribeLocation(width int) string {

	currentTile := p.Room.Layout[p.Location.Y][p.Location.X]
	if t, ok := currentTile.(Item); ok {
		if t.Walkable {
			ll := cl.LogLine{
				Text:        fmt.Sprintf("[x] exit the %s", t.Name),
				BgRGB:       t.Color,
				FontRGB:     t.FontColor,
				WithNewLine: true,
				WithPadding: true,
				Padding:     width,
			}

			returnable := ll.GetString()
			returnable += "\n\n\n\n"

			return returnable
		}
	}

	front := p.Location.Y - 1
	back := p.Location.Y + 1
	left := p.Location.X - 1
	right := p.Location.X + 1

	frontTile := p.Room.Layout[front][p.Location.X]
	backTile := p.Room.Layout[back][p.Location.X]
	leftTile := p.Room.Layout[p.Location.Y][left]
	rightTile := p.Room.Layout[p.Location.Y][right]

	returnable := ""
	newlinesToAdd := 0
	forwardDescription := PrintDirection(frontTile, FORWARD, width)
	if forwardDescription == "" {
		newlinesToAdd++
	} else {
		returnable += forwardDescription
	}
	backwardDescription := PrintDirection(backTile, BACKWARD, width)
	if backwardDescription == "" {
		newlinesToAdd++
	} else {
		returnable += backwardDescription
	}
	lefDescription := PrintDirection(leftTile, LEFT, width)
	if lefDescription == "" {
		newlinesToAdd++
	} else {
		returnable += lefDescription
	}
	rightDescription := PrintDirection(rightTile, RIGHT, width)
	if rightDescription == "" {
		newlinesToAdd++
	} else {
		returnable += rightDescription
	}

	for i := 0; i < newlinesToAdd; i++ {
		returnable += "\n"
	}

	if returnable != "" {
		returnable = fmt.Sprintf("\n%s", returnable)
	}

	return returnable
}
