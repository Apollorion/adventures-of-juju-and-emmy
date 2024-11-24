package main

import (
	cl "colorlib"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type game struct {
	Player *Player

	World map[string]*Room

	view struct {
		width  int
		height int
		screen int
	}
}

func (g game) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	// clear the screen
	fmt.Println("\033[H\033[2J")
	return nil
}

func (g game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c":
			// clear the screen
			fmt.Println("\033[H\033[2J")
			return g, tea.Quit

		// The "up" and "w" keys move the cursor up
		case "up", "w":
			g.Player.Move(FORWARD)

		// The "down" and "s" keys move the cursor down
		case "down", "s":
			g.Player.Move(BACKWARD)

		// The "left" and "a" keys move the cursor left
		case "left", "a":
			g.Player.Move(LEFT)

		// The "right" and "d" keys move the cursor right
		case "right", "d":
			g.Player.Move(RIGHT)

		case "i":
			if g.view.screen == CONFIRM_PLACE {
				g.Player.PlaceItem(FORWARD)
				g.view.screen = ROOMS
			} else {
				g.Player.Interact(&g, FORWARD)
			}
		case "k":
			if g.view.screen == CONFIRM_PLACE {
				g.Player.PlaceItem(BACKWARD)
				g.view.screen = ROOMS
			} else {
				g.Player.Interact(&g, BACKWARD)
			}
		case "j":
			if g.view.screen == CONFIRM_PLACE {
				g.Player.PlaceItem(LEFT)
				g.view.screen = ROOMS
			} else {
				g.Player.Interact(&g, LEFT)
			}
		case "l":
			if g.view.screen == CONFIRM_PLACE {
				g.Player.PlaceItem(RIGHT)
				g.view.screen = ROOMS
			} else {
				g.Player.Interact(&g, RIGHT)
			}
		case "x":
			g.Player.Interact(&g, EXIT)

		case "o":
			if g.view.screen == ROOMS {
				g.view.screen = INVENTORY
			} else {
				g.view.screen = ROOMS
			}

		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			if g.view.screen == INVENTORY {
				i := int(msg.String()[0]) - 49
				if i < len(g.Player.Inventory) {
					g.view.screen = CONFIRM_PLACE
					g.Player.HeldItem = i + 1
				}
			}
		}

	case tea.WindowSizeMsg:
		// clear the screen
		fmt.Println("\033[H\033[2J")
		g.view.width = msg.Width
		g.view.height = msg.Height
	}

	return g, nil
}

func (g game) View() string {

	if g.view.screen == 0 {
		g.view.screen = ROOMS
	}

	s := ""
	if g.view.screen == ROOMS {
		if g.Player.Room == nil {
			g.Player.Room = g.GetRoom("house")
			g.Player.Location = g.Player.Room.StartingPosition
		}

		heightPadding := len(g.Player.Room.Layout) - g.view.height
		for h := 0; h < heightPadding; h++ {
			s += "\n"
		}

		if g.Player.Room.Status != "" {
			s += cl.SprintlnP(g.Player.Room.Status, defaultTextColor, g.view.width).GetString()
			g.Player.Room.Status = ""
		} else {
			s += "\n"
		}
		s += g.Player.DescribeLocation(g.view.width)
		s += g.Player.PrintRoom(g.view.width)
	} else if g.view.screen == INVENTORY {
		s += cl.SprintlnP("Inventory", defaultTextColor, g.view.width).GetString()
		s += cl.SprintlnP("---------", defaultTextColor, g.view.width).GetString()
		for i, item := range g.Player.Inventory {
			s += cl.SprintlnP(fmt.Sprintf("%d. %s", i+1, item.Name), defaultTextColor, g.view.width).GetString()
		}
	} else if g.view.screen == CONFIRM_PLACE {
		s += cl.SprintlnP("Place the item where?", defaultTextColor, g.view.width).GetString()
		s += cl.SprintlnP("[i] Front", defaultTextColor, g.view.width).GetString()
		s += cl.SprintlnP("[j] Left", defaultTextColor, g.view.width).GetString()
		s += cl.SprintlnP("[k] Back", defaultTextColor, g.view.width).GetString()
		s += cl.SprintlnP("[l] Right", defaultTextColor, g.view.width).GetString()
	}

	s += "\n"
	s += cl.SprintlnP("What would you like to do?", defaultTextColor, g.view.width).GetString()
	s += cl.SprintlnP("[w,a,s,d] Move", defaultTextColor, g.view.width).GetString()
	s += cl.SprintlnP("[i,j,k,l] Interact", defaultTextColor, g.view.width).GetString()
	s += cl.SprintlnP("[o] Open/Close Inventory", defaultTextColor, g.view.width).GetString()

	// Send the UI for rendering
	return s
}

func (g game) GetRoom(name string) *Room {
	if room, ok := g.World[name]; ok {
		return room
	} else {
		switch name {
		case "house":
			g.World[name] = NewHouse(&g)
		case "house_basement":
			g.World[name] = NewHouseBasement(&g)
		default:
			panic("Room not found")
		}
		return g.World[name]
	}
}
