package command

import (
	"fmt"
	"os"

	"github.com/sajad-dev/commit-duration-checker/internal/clear"
	"github.com/sajad-dev/commit-duration-checker/internal/git"
)



func help()  {
fmt.Print(`
git             Your Git Commands
clear           Clear Logs Activity

`)	
}

func Command() {
	switch os.Args[1] {
	case "git":
		git := git.Git{}
		git.HandelCommand()
	case "clear":
		clear.Clear()
	case "help":
		help()
	}
}
