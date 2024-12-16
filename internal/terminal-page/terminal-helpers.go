package terminalpage

import (

	"github.com/nsf/termbox-go"
	"golang.org/x/exp/slices"
)

type TerminalPanelHelper interface {
	KeyTabShift()
	KeyTab()
	KeyArrowLeft()
	KeyArrowRight()
	KeyEnter()
	KeyDelete()
	KeyCtrlA()
	KeyBackspace2()
	KeyCtrlC(chan int)
	GetKey(int)
}

var _ TerminalPanelHelper = (*TerminalWriter)(nil)

func (terminal *TerminalWriter) KeyTabShift() {
	terminal.NodeHover = terminal.NodeActive
	if terminal.NodeHover.Parent != nil {
		terminal.NodeHover = terminal.NodeHover.Parent
	} else {
		var child = &Node{}
		child = terminal.NodeHover
		for len(child.Children) != 0 {
			child = child.Children[0]
		}
		terminal.NodeHover = child
	}
}

func (terminal *TerminalWriter) KeyTab() {
	if len(terminal.NodeHover.Children) != 0 && !slices.Contains(terminal.NodeActive.Children, terminal.NodeHover) {
		terminal.NodeHover = terminal.NodeActive.Children[0]
	} else {
		terminal.NodeHover = terminal.NodeActive
	}

}

func (terminal *TerminalWriter) KeyArrowLeft() {
	if terminal.NodeHover.Edge["left"] != nil {
		terminal.NodeHover = terminal.NodeHover.Edge["left"]
	}
}

func (terminal *TerminalWriter) KeyArrowRight() {
	if terminal.NodeHover.Edge["right"] != nil {

		terminal.NodeHover = terminal.NodeHover.Edge["right"]
	}

}

func (terminal *TerminalWriter) KeyArrowUp() {
	if terminal.NodeActive.Click["keyup"] != nil {
		terminal.NodeActive.Click["keyup"](terminal, nil)
		return
	}

	if terminal.NodeHover.Edge["top"] != nil {

		terminal.NodeHover = terminal.NodeHover.Edge["top"]
	}
}

func (terminal *TerminalWriter) KeyArrowDown() {

	if terminal.NodeActive.Click["keydown"] != nil {
		terminal.NodeActive.Click["keydown"](terminal, nil)
		return
	}

	if terminal.NodeHover.Edge["botton"] != nil {
		terminal.NodeHover = terminal.NodeHover.Edge["botton"]
	}

}

func (terminal *TerminalWriter) KeyEnter() {
	
	if terminal.Writer {
		terminal.NodeActive.Click["cntrla"](terminal, -2)

		return
	}
	if terminal.NodeHover.Click["enter"] != nil {
		terminal.NodeHover.Click["enter"](terminal, nil)
	} else {
		terminal.NodeActive = terminal.NodeHover
		Clear(terminal.X_start, terminal.Y_start, terminal.Width, terminal.Height)
		if len(terminal.NodeActive.Children) != 0 {

			terminal.NodeHover = terminal.NodeActive.Children[0]
		}
	}

}

func (terminal *TerminalWriter) KeyDelete() {
	if terminal.NodeHover.Click["delete"] != nil {
		terminal.NodeHover.Click["delete"](terminal,nil)
	}
}

func (terminal *TerminalWriter) KeyCtrlA() {

	if terminal.Writer {
		terminal.Writer = false
		Clear(terminal.Y_start+1, 0, terminal.Width, terminal.Height)
		return

	}
	if terminal.NodeActive.Click["cntrla"] != nil {
		terminal.NodeActive.Click["cntrla"](terminal, nil)
		terminal.Writer = true
	}

}

func (terminal *TerminalWriter) KeyBackspace2() {
	if terminal.Writer {
		terminal.NodeActive.Click["cntrla"](terminal, -1)
	}
}

func (terminal *TerminalWriter) KeyCtrlC(keypress chan int) {
	keypress <- 11111111
	termbox.Close()
}

func (terminal *TerminalWriter) GetKey(key int) {
	if terminal.Writer {
		terminal.NodeActive.Click["cntrla"](terminal, key)
	}
}
