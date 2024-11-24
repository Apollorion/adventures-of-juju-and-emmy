package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	g := game{
		Player: NewPlayer(),
		World:  make(map[string]*Room),
	}

	p := tea.NewProgram(g)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
