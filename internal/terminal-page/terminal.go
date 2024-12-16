package terminalpage

import (
	"fmt"
	"io"
	"os"

	"github.com/nsf/termbox-go"
)

type TerminalPanelInterface interface {
	KeyHandel(keypress chan int)
	write()
}

var welcom = ""

func (terminal *TerminalWriter) KeyHandel(keypress chan int) {
	termbox.Init()
	AddWelcom()
	readFile()
	for {
		if len(os.Args) == 1 {
			terminal.write()
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyCtrlC:
					terminal.KeyCtrlC(keypress)
				case termbox.KeyArrowLeft:
					terminal.KeyArrowLeft()
				case termbox.KeyArrowUp:
					terminal.KeyArrowUp()
				case termbox.KeyArrowRight:
					terminal.KeyArrowRight()
				case termbox.KeyArrowDown:
					terminal.KeyArrowDown()
				case termbox.KeyTab:
					terminal.KeyTab()
				case termbox.KeyCtrlV:
					terminal.KeyTabShift()
				case termbox.KeyEnter:
					terminal.KeyEnter()
				case termbox.KeyDelete:
					terminal.KeyDelete()
				case termbox.KeyCtrlA:
					terminal.KeyCtrlA()
				case termbox.KeyBackspace2:
					terminal.KeyBackspace2()
				default:
					terminal.GetKey(int(ev.Ch))
					// page.GetKey(int(ev.Ch))
				}
			}
		} else {
			terminal.noPanel()
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyCtrlC:
					terminal.KeyCtrlC(keypress)
				}
			}
		}
	}
}

func AddWelcom() {
	os.Chdir("static")
	file, err := os.Open("welcome.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	data, _ := io.ReadAll(file)
	welcom = string(data)
}

func (terminal *TerminalWriter) noPanel() {
	width, height := termbox.Size()
	terminal.ImageWriter(width, height, 0, 0, "logo.jpeg")
	y := 1
	x := width/2 + 1
	for _, v := range welcom {
		x++
		if x > width {
			x = width/2 + 1
			y++
		}
		termbox.SetCell(x, y, v, termbox.ColorWhite, termbox.ColorDefault)
	}
	termbox.Flush()
}

func (terminal *TerminalWriter) paginate() {

	terminal.NodeActive.Value.Output(terminal, nil)
	if terminal.NodeHover.RunInHover {
		terminal.NodeHover.Value.Output(terminal, nil)
	}
}

func (terminal *TerminalWriter) write() {

	for _, v := range terminal.NodeList {
		if v.Value.Lable == terminal.NodeHover.Value.Lable {
			terminal.NodeHover = v
		}
	}

	AddItemMenu(terminal)
	width, height := termbox.Size()

	terminal.Height = height
	terminal.Width = width
	terminal.X_end = width
	terminal.Y_end = height
	terminal.X_start = 1
	terminal.Y_start = 1

	terminal.menu()

	terminal.paginate()

	// terminal.footer()

	termbox.Flush()

}

func Handel(keypress chan int) {
	var terminal TerminalPanelInterface

	terminal = &TerminalWriter{}

	terminal.KeyHandel(keypress)
}
