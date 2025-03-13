package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/goferwplynie/goXP/internal/modules/filepicker"
)

func main() {
	fp := filepicker.New()
	fp.CurrentDir = filepicker.SetupPath()
	fp.ReadDir()()
	p := tea.NewProgram(fp)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
