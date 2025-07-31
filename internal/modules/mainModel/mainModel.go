package mainmodel

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/goferwplynie/goXP/internal/modules/cmdline"
	"github.com/goferwplynie/goXP/internal/modules/filepicker"
)

type Focus int

const (
	filePickerFocus Focus = iota
	cmdLineFocus
)

type Model struct {
	Filepicker filepicker.Model
	Cmdline    cmdline.Model
	focus      Focus
	KeyBinds   KeyBinds
}

type KeyBinds struct {
	FocusFilePicker key.Binding
	FocusCmd        key.Binding
}

func NewKeyBinds() KeyBinds {
	return KeyBinds{
		FocusFilePicker: key.NewBinding(key.WithKeys("esc")),
		FocusCmd:        key.NewBinding(key.WithKeys(":")),
	}
}

func New(fp filepicker.Model, cmdline cmdline.Model) Model {
	return Model{
		Filepicker: fp,
		Cmdline:    cmdline,
		focus:      cmdLineFocus,
		KeyBinds:   NewKeyBinds(),
	}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyBinds.FocusFilePicker):
			m.focus = filePickerFocus
			return m, nil
		case key.Matches(msg, m.KeyBinds.FocusCmd):
			m.focus = cmdLineFocus
			return m, nil
		}
	}

	switch m.focus {
	case filePickerFocus:
		updated, cmd := m.Filepicker.Update(msg)
		if fp, ok := updated.(filepicker.Model); ok {
			m.Filepicker = fp
		}
		return m, cmd
	case cmdLineFocus:
		updated, cmd := m.Cmdline.Update(msg)
		if cmdline, ok := updated.(cmdline.Model); ok {
			m.Cmdline = cmdline
		}
		return m, cmd
	}
	return m, nil
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m Model) View() string {
	main := lipgloss.JoinVertical(lipgloss.Top, m.Filepicker.View(), m.Cmdline.View())
	return main
}
