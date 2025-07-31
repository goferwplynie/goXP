package cmdline

import (
	"fmt"
	"os"
	"strings"

	"github.com/goferwplynie/goXP/internal/modules/filepicker"
)

type Command struct {
	Name string
	Args []string
}

type CommandFunc func(args []string, fp filepicker.FilePickerApi) error

type CommandRegistry map[string]CommandFunc

func parseInput(command string) Command {
	splitted := strings.Split(command, " ")
	return Command{
		Name: splitted[0],
		Args: splitted[1:],
	}
}

func ChangeDir(args []string, fp filepicker.FilePickerApi) error {
	if len(args) <= 0 {
		return fmt.Errorf("cd needs at least 1 argument")
	}
	path := args[0]
	_, err := os.Stat(path)
	if err != nil {
		// path doesn't exist or some other error
		if os.IsNotExist(err) {
			return fmt.Errorf("directory does not exist!")
		} else {
			return fmt.Errorf("Other error: %v", err)
		}
	}

	return nil
}

func DefaultCommands() CommandRegistry {
	return CommandRegistry{
		"cd": ChangeDir,
	}
}
