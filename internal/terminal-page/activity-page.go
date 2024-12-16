package terminalpage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/nsf/termbox-go"
)

var columns = [][]string{}

var scrollnum = 0

var nup_pag = 2

func (terminal *TerminalWriter) AddToColumns() {
	os.Chdir("static")

	file, _ := os.Open("data.json")

	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var data []map[string]interface{}
	json.Unmarshal(byteValue, &data)

	for _, v := range data {
		hash_map := []string{}
		hash_map = append(hash_map, v["window"].(string))
		hash_map = append(hash_map, v["secend"].(string))
		hash_map = append(hash_map, v["repository"].(string))
		hash_map = append(hash_map, v["commit"].(string))
		columns = append(columns, hash_map)
	}

}

func ActivityHandel(terminal *TerminalWriter, a any) {
	columns = [][]string{}
	terminal.AddToColumns()
	terminal.WriteActivity()

	if nup_pag <= len(columns)-1 {

		terminal.ScrollBar()
	}
}

func scrollDown(terminal *TerminalWriter, a any) {
	if scrollnum+nup_pag <= len(columns)-1 {
		scrollnum++

	}
	Clear(terminal.Y_start, 0, terminal.Width, terminal.Height)

}

func scrollUp(terminal *TerminalWriter, a any) {
	if scrollnum-1 >= 0 {
		scrollnum--
	}
	Clear(terminal.Y_start, 0, terminal.Width, terminal.Height)

}

func (terminal *TerminalWriter) ScrollBar() {
	x := float32(float32(terminal.Height-terminal.Y_start) / float32(len(columns)+1-nup_pag))
	for i := terminal.Y_start; i < terminal.Height-2; i++ {
		if x*float32(scrollnum) < float32(i) && x*float32((scrollnum+1)) > float32(i) {
			termbox.SetCell(terminal.X_end-8, i+2, ' ', termbox.ColorWhite, termbox.ColorGreen)
			termbox.SetCell(terminal.X_end-9, i+2, ' ', termbox.ColorWhite, termbox.ColorGreen)

		} else {
			termbox.SetCell(terminal.X_end-8, i+2, ' ', termbox.ColorWhite, termbox.ColorLightGreen)
			termbox.SetCell(terminal.X_end-9, i+2, ' ', termbox.ColorWhite, termbox.ColorLightGreen)

		}
	}
}

func (terminal *TerminalWriter) WriteActivity() {
	nup_pag = terminal.Height - terminal.Y_start
	scroll := [][]string{}
	if nup_pag >= len(columns) {
		scroll = columns
	} else {
		scroll = columns[scrollnum : scrollnum+nup_pag]

	}

	max_width_column := (terminal.Width - 100) / 4
	for i, v := range scroll {
		x_start := 0
		for j, val := range v {
			milsecend, err := strconv.Atoi(val)
			if err == nil {
				scend := milsecend / 1000
				min := 0
				if scend >= 60 {
					min = scend / 60
					scend = scend % 60

				}
				val = fmt.Sprintf("%d:%d",min,scend)
			}
			for index, value := range val {
				termbox.SetCell(index+x_start, terminal.Y_start+i+2, value, termbox.ColorWhite, termbox.ColorDefault)
			}

			if j == 0 {
				x_start += max_width_column + 100

			} else {

				x_start += max_width_column + 2
			}
		}
	}
}
