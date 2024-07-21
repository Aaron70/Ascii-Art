package terminal

import (
	"os"
	"os/exec"

	"golang.org/x/term"
)

const _CLEAR_COMMAND_LINUX string = "clear"

func ClearScreen() {
  exec.Command(_CLEAR_COMMAND_LINUX)
}

func GetScreenDimensions() (width, height int, err error) {
  return term.GetSize(int(os.Stdin.Fd()))
}
