package cmdline

import (
	"errors"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/goferwplynie/goXP/internal/modules/filepicker"
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

type CommandFunc func(args []string, m filepicker.Model) error

func deleteFile(args []string, m filepicker.Model) error {
	if len(args) > 0 {
		if m.DeleteFile(args[0]) {
			return nil
		} else {
			return errors.New("error deleting file")
		}
	} else if len(m.GetSelected()) > 0 {
		for _, v := range m.GetSelected() {
			if m.DeleteFile(v.Name()) {
				return nil
			} else {
				return errors.New("error deleting file")
			}
		}
	}
	return errors.New("error deleting file")
}

func New() Model {
	return Model{
		cursor:   cursor.New(),
		Keybinds: GetKeybinds(),
	}
}

func GetKeybinds() KeyBinds {
	return KeyBinds{
		Quit:            key.NewBinding(key.WithKeys("esc")),
		EnterCommand:    key.NewBinding(key.WithKeys("enter")),
		DeleteCharacter: key.NewBinding(key.WithKeys("backspace")),
		MoveForward:     key.NewBinding(key.WithKeys("right")),
		MoveBackward:    key.NewBinding(key.WithKeys("left")),
	}
}

func (m *Model) readRunes(msg tea.KeyMsg) {
	//crazy work
	m.value = append(m.value[:m.cursorPos], append(msg.Runes, m.value[m.cursorPos:]...)...)
	m.cursorPos++
}

func handleCommand(command string) {

}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keybinds.Quit):
			break
		case key.Matches(msg, m.Keybinds.EnterCommand):
			handleCommand(string(m.value))
		case key.Matches(msg, m.Keybinds.DeleteCharacter):
			if len(m.value) > 0 {
				m.value = append(m.value[:m.cursorPos-1], m.value[m.cursorPos:]...)
				m.cursorPos--
			}
		case key.Matches(msg, m.Keybinds.MoveBackward):
			m.cursorPos--
			if m.cursorPos > 0 {
			}
		case key.Matches(msg, m.Keybinds.MoveForward):
			if m.cursorPos < len(m.value) {
				m.cursorPos++
			}
		default:
			m.readRunes(msg)
		}
	}

	return m, nil
}

func (m Model) View() string {
	return string(m.value[:m.cursorPos]) + "|" + string(m.value[m.cursorPos:])
}
