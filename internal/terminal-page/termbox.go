package terminalpage

import "github.com/nsf/termbox-go"

func LineX(y int, x int, width int, sp rune) {
	for i := 0; i < width-x*2; i++ {
		termbox.SetCell(x+i, y, sp, termbox.ColorWhite, termbox.ColorDefault)
	}
}

func LineY(y int, x int, height int) {
	for i := 0; i < height-y; i++ {
		termbox.SetCell(x, y+1, '|', termbox.ColorWhite, termbox.ColorDefault)
	}
}

func Clear(y int, x int, width int, height int) {
	for i := 0; i < width; i++ {
		for j := y; j < height; j++ {
			termbox.SetCell(i, j, ' ', termbox.ColorWhite, termbox.ColorDefault)
		}
	}
}
