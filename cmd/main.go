package main

import (
	"fmt"
	"os"
	"raven/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	//m := ui.SecondModel()
	//if _, err := tea.NewProgram(m).Run(); err != nil {
	//	fmt.Println("Oh no!", err)
	//	os.Exit(1)
	//}
}
