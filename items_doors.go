package main

func NewDoor(doorLocation Coordinate, upAndDown bool) Item {

	var a, b Coordinate

	if upAndDown {
		a = Coordinate{doorLocation.X, doorLocation.Y - 1}
		b = Coordinate{doorLocation.X, doorLocation.Y + 1}
	} else {
		a = Coordinate{doorLocation.X - 1, doorLocation.Y}
		b = Coordinate{doorLocation.X + 1, doorLocation.Y}
	}

	return Item{
		Name:         "Door",
		ShortName:    "D",
		Description:  "A door",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        doorColor,
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			// move player to other side of door
			if p.Location.X == a.X && p.Location.Y == a.Y {
				p.Location.X = b.X
				p.Location.Y = b.Y
			} else {
				p.Location.X = a.X
				p.Location.Y = a.Y
			}
		},
	}
}

func NewLocationChangingDoor(newLocation string, symbol string) Item {
	return Item{
		Name:         "Door",
		ShortName:    symbol,
		Description:  "A door",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        doorColor,
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			location := g.GetRoom(newLocation)

			p.Room = location
			p.Location = location.StartingPosition
		},
	}
}

func NewLocationChangingDoorWithCoordinateOverride(newLocation string, symbol string, override Coordinate) Item {
	return Item{
		Name:         "Door",
		ShortName:    symbol,
		Description:  "A door",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        doorColor,
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			location := g.GetRoom(newLocation)

			p.Room = location
			p.Location = override
		},
	}
}

func NewLocationChangingLockableDoor(newLocation string, symbol string, requiredKeySig string) Item {
	return Item{
		Name:         "Door",
		ShortName:    symbol,
		Description:  "A door",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        doorColor,
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {

			playerHasCorrectKey := false
			for _, item := range p.Inventory {
				if sig, ok := item.State["keySig"]; ok {
					if sig == requiredKeySig {
						playerHasCorrectKey = true
						break
					}
				}
			}

			if playerHasCorrectKey {
				location := g.GetRoom(newLocation)
				p.Room = location
				p.Location = location.StartingPosition
			} else {
				p.Room.Status = "The door is locked!"
			}
		},
	}
}

func NewDescriptiveDoor(description string, symbol string) Item {
	return Item{
		Name:         "Door",
		ShortName:    symbol,
		Description:  "A door",
		Weight:       100,
		Value:        100,
		Interactable: true,
		Carryable:    false,
		Color:        doorColor,
		Interaction: func(g *game, p *Player, direction int, c Coordinate, i *Item) {
			p.Room.Status = description
		},
	}
}
