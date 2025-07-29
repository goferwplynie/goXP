package main

import (
	"fmt"
	"os"

	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/goferwplynie/goXP/config"
	"github.com/goferwplynie/goXP/internal/modules/cmdline"
	"github.com/goferwplynie/goXP/internal/modules/filepicker"
)

var wg sync.WaitGroup

//test czy git dziala

func main() {
	fp := Setup()
	p := tea.NewProgram(fp)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func Setup() tea.Model {
	conf := config.ConfigLoader()
	fp := filepicker.New()
	fp.CurrentDir = filepicker.SetupPath()
	fp.ReadDir()()
	fpStyles := conf.FilePickerConfig.Styles
	fp.Styles = filepicker.CustomStyle(fpStyles)
	fp.Keybinds = filepicker.CustomKeybinds(conf.FilePickerConfig.Keybinds)

	cmd := cmdline.New()
	//to test filepicker just return fp
	return cmd
}
