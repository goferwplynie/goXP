package cmdline

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	cursorPos int

	PlaceHolder string
	value       []rune
	cursor      cursor.Model
	Keybinds    KeyBinds
}

type KeyBinds struct {
	Quit            key.Binding
	EnterCommand    key.Binding
	DeleteCharacter key.Binding
	MoveForward     key.Binding
	MoveBackward    key.Binding
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return ""
}
