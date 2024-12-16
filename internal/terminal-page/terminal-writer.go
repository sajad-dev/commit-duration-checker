package terminalpage

import (
	"github.com/nsf/termbox-go"
)

type TerminalWriterInterface interface {
	menu()
	footer()
	paginate()
}

var menu_list = []*Value{
	&Value{Lable: " Home ", Output: HomeHandel},
	&Value{Lable: " Activity ", Output: ActivityHandel},
	&Value{Lable: " Contact ", Output: AbouHandel},
}

func AddItemMenu(page *TerminalWriter) {
	page.Clear()
	
	page.Add(map[string]*Node{}, nil, &Value{Lable: menu_list[0].Lable, Output: menu_list[0].Output}, false, map[string]FunctionTp{"cntrla": addHomeTask})
	
	page.Add(map[string]*Node{"left": page.NodeList[0]}, nil, &Value{Lable: menu_list[1].Lable, Output: menu_list[1].Output}, false, map[string]FunctionTp{"keydown": scrollDown,"keyup": scrollUp})
	
	page.Add(map[string]*Node{"left": page.NodeList[1]}, nil, &Value{Lable: menu_list[2].Lable, Output: menu_list[2].Output}, false, map[string]FunctionTp{})

	page.NodeList[0].Edge["left"] = page.NodeList[len(page.NodeList)-1]
	page.NodeList[len(page.NodeList)-1].Edge["right"] = page.NodeList[0]
	if page.NodeActive == nil {
		page.NodeHover = page.NodeList[0]
		page.NodeActive = page.NodeList[0]

	}
}

func (terminal *TerminalWriter) menu() {
	var lenarr int
	for _, item := range menu_list {
		lenarr = lenarr + len(item.Lable)
	}

	var last int
	var menu_list = terminal.Show(0)
	for i, v := range menu_list {
		switch {
		case terminal.NodeList[i].Value.Lable == terminal.NodeHover.Value.Lable:
			for index, ch := range v.Lable {
				termbox.SetCell(terminal.X_start+index+last, terminal.Y_start, ch, termbox.ColorWhite, termbox.ColorBlue)
			}
		case terminal.NodeList[i].Value.Lable == terminal.NodeActive.Value.Lable:
			for index, ch := range v.Lable {
				termbox.SetCell(terminal.X_start+index+last, terminal.Y_start, ch, termbox.ColorWhite, termbox.ColorRed)
			}
		default:
			for index, ch := range v.Lable {
				termbox.SetCell(terminal.X_start+index+last, terminal.Y_start, ch, termbox.ColorWhite, termbox.ColorDefault)
			}
		}

		last = len(v.Lable) + last + (((terminal.Width - terminal.X_start*2) - lenarr) / (len(menu_list) - 1))
	}
	LineX(terminal.Y_start+1, terminal.X_start, terminal.Width, '_')
	terminal.Y_start = terminal.Y_start + 1
}

