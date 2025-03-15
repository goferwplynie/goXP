package main

import (
	"fmt"
	"os"

	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/goferwplynie/goXP/config"
	"github.com/goferwplynie/goXP/internal/modules/filepicker"
)

var wg sync.WaitGroup

func main() {
	fp := Setup()
	p := tea.NewProgram(fp)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func Setup() tea.Model {
	// confCh := make(chan config.Config)
	// go func(c chan config.Config) {
	// 	defer wg.Done()
	// 	c <- conf
	// }(confCh)
	conf := config.ConfigLoader()
	fp := filepicker.New()
	fp.CurrentDir = filepicker.SetupPath()
	fp.ReadDir()()
	fpStyles := conf.FilePickerConfig.Styles
	fp.Styles = filepicker.CustomStyle(fpStyles)
	fp.Keybinds = filepicker.CustomKeybinds(conf.FilePickerConfig.Keybinds)

	return fp
}
