package cmdline

type Model struct {
	cursorPos int

	PlaceHolder string
	value       []rune
	Commands    map[string]func() bool
}
