package main

func RemoveFromInventory(s []*Item, i int) []*Item {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func PickUpItem(roomStatus string) func(g *game, p *Player, direction int, c Coordinate, i *Item) {
	return func(g *game, p *Player, direction int, c Coordinate, i *Item) {

		if len(p.Inventory) == 9 {
			p.Room.Status = "You can't carry any more items."
			return
		}

		tileUp := p.Room.Layout[p.Location.Y-1][p.Location.X]
		tileDown := p.Room.Layout[p.Location.Y+1][p.Location.X]
		tileLeft := p.Room.Layout[p.Location.Y][p.Location.X-1]
		tileRight := p.Room.Layout[p.Location.Y][p.Location.X+1]

		switch direction {
		case FORWARD:
			if t, ok := tileUp.(Item); ok {
				if t.Carryable {
					p.Inventory = append(p.Inventory, &t)
					p.Room.Layout[p.Location.Y-1][p.Location.X] = "."
				}
			}
		case BACKWARD:
			if t, ok := tileDown.(Item); ok {
				if t.Carryable {
					p.Inventory = append(p.Inventory, &t)
					p.Room.Layout[p.Location.Y+1][p.Location.X] = "."
				}
			}
		case LEFT:
			if t, ok := tileLeft.(Item); ok {
				if t.Carryable {
					p.Inventory = append(p.Inventory, &t)
					p.Room.Layout[p.Location.Y][p.Location.X-1] = "."
				}
			}
		case RIGHT:
			if t, ok := tileRight.(Item); ok {
				if t.Carryable {
					p.Inventory = append(p.Inventory, &t)
					p.Room.Layout[p.Location.Y][p.Location.X+1] = "."
				}
			}
		}

		p.Room.Status = roomStatus
	}
}
