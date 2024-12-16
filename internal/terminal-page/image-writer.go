package terminalpage

import (
	"github.com/nsf/termbox-go"
	imagetodot "github.com/sajad-dev/commit-duration-checker/pkg/imageto-dot"

)

func (t *TerminalWriter) ImageWriter(w int, h int, x_start int, y_start int,imgname string) {

	img := imagetodot.Handel(w/2, h,imgname)
	symbols := "@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@             "
	var arrsymbols = []rune(symbols)
	for index, value := range *img {

		for ind, val := range value {
			termbox.SetCell(ind+x_start, index+y_start, arrsymbols[val], termbox.ColorRed, termbox.ColorDefault)
		}
	}

}
