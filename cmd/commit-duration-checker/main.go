package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sajad-dev/commit-duration-checker/internal/command"
	"github.com/sajad-dev/commit-duration-checker/internal/keybord"
	terminalpanel "github.com/sajad-dev/commit-duration-checker/internal/terminal-page"
	timeactivity "github.com/sajad-dev/commit-duration-checker/internal/time-activity"
)

func main() {


	if len(os.Args) != 1 && os.Args[1] != "-np" && os.Args[1] != "--no-pandel" {
		command.Command()
		return
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	keypress := make(chan int, 10)
	go terminalpanel.Handel(keypress)

	go keybord.Keyboard(keypress)

	for {
		if keybord.CheckEndKey(keypress) {

			return
		}

		timeactivity.TimeCa(keypress)

	}

}
