package terminalpage

import (
	"math"
	"time"

	"github.com/nsf/termbox-go"
)

type Task struct {
	Status     int
	Title      string
	Deedline   string
	Dic        string
	Ac         string
	RunInHover bool
	Node       *Node
	Function   func(*TerminalWriter, any)
}

type HomeInterface interface {
	TaskTab()
	AddTask()
}

func (tr *TerminalWriter) FirstPage() {
	tr.ImageWriter(tr.Width-46, tr.Height-8, 45, 5, "logo.jpeg")
	y := 4
	x := tr.Width - 46

	for _, v := range welcom {
		x++
		if x > tr.Width {
			x = tr.Width - 46
			y++
		}
		termbox.SetCell(x, y, v, termbox.ColorWhite, termbox.ColorDefault)
	}

}

func HomeHandel(tr *TerminalWriter, a any) {

	tr.AddTask()
	tr.TaskTab()
	tr.FirstPage()
}

var tasks = []Task{}

func RunInfoTask(terminal *TerminalWriter, a any) {

	for i := terminal.Y_start + 3; i < terminal.Y_end; i++ {
		for j := terminal.X_start + 45; j < terminal.X_end; j++ {

			termbox.SetCell(j, i, ' ', termbox.ColorWhite, termbox.ColorDefault)
		}
	}

	task, _ := terminal.NodeHover.Value.Val.(Task)

	for i, v := range task.Dic {
		// Y:=
		w := int(math.Floor(float64(i) / float64(terminal.X_end-55)))
		termbox.SetCell(45+i%(terminal.X_end-55), terminal.Y_start+w+3, v, termbox.ColorWhite, termbox.ColorDefault)

	}
}

func taskEnter(terminal *TerminalWriter, a any) {
	for i, v := range tasks {
		if v.Title == terminal.NodeHover.Value.Lable {
			if tasks[i].Status == 2 {
				tasks[i].Status = 0
			} else {
				tasks[i].Status = 2
			}
		}
	}
}

func (home *TerminalWriter) AddTask() {

	var node *Node
	leng := len(home.NodeList)

	for index, task := range tasks {
		if index != 0 {
			node = home.Add(map[string]*Node{"top": home.NodeList[index-1+leng]},
				home.NodeList[0], &Value{Lable: task.Title, Output: RunInfoTask, Val: task}, true, map[string]FunctionTp{"enter": taskEnter, "delete": deleteTask})
		} else {
			node = home.Add(map[string]*Node{}, home.NodeList[0],
				&Value{Lable: task.Title, Output: RunInfoTask, Val: task}, true, map[string]FunctionTp{"enter": taskEnter, "delete": deleteTask})
		}
		tasks[index].Function = RunInfoTask
		tasks[index].Node = node
	}

	home.NodeList[leng].Edge["top"] = home.NodeList[len(home.NodeList)-1]
	home.NodeList[len(home.NodeList)-1].Edge["botton"] = home.NodeList[leng]

}

func (home *TerminalWriter) TaskTab() {
	for i, v := range tasks {
		inputDate, _ := time.Parse("2006/1/2", v.Deedline)
		now := time.Now()
		if inputDate.Before(now) && v.Status == 0 {
			tasks[i].Status = 3
		}
	}

	for index, task := range tasks {
		title := task.Title
		switch task.Status {
		case 0:
			title = "⬜" + task.Title
			for i, chac := range title {
				if home.NodeHover != nil && home.NodeHover.Value.Lable == task.Node.Value.Lable {
					termbox.SetCell(home.X_start+i, home.Y_start+index*3+2, chac, termbox.ColorBlue, termbox.ColorGreen)

				} else {

					termbox.SetCell(home.X_start+i, home.Y_start+index*3+2, chac, termbox.ColorWhite, termbox.ColorDefault)
				}
			}
		case 1:

			title = "⧗" + task.Title
			for i, chac := range title {
				if home.NodeHover != nil && home.NodeHover.Value.Lable == task.Node.Value.Lable {
					termbox.SetCell(home.X_start+i, home.Y_start+index*3+2, chac, termbox.ColorBlue, termbox.ColorGreen)

				} else {

					termbox.SetCell(home.X_start+i, home.Y_start+index*3+2, chac, termbox.ColorWhite, termbox.ColorDefault)
				}
			}
		case 2:
			title = "✅" + task.Title
			for i, chac := range title {
				if home.NodeHover != nil && home.NodeHover.Value.Lable == task.Node.Value.Lable {
					termbox.SetCell(home.X_start+i, home.Y_start+index*3+2, chac, termbox.ColorBlue, termbox.ColorGreen)

				} else {

					termbox.SetCell(home.X_start+i, home.Y_start+index*3+2, chac, termbox.ColorWhite, termbox.ColorDefault)
				}
			}
		case 3:
			title = "❌" + task.Title
			for i, chac := range title {
				if home.NodeHover != nil && home.NodeHover.Value.Lable == task.Node.Value.Lable {
					termbox.SetCell(home.X_start+i, home.Y_start+index*3+2, chac, termbox.ColorBlue, termbox.ColorGreen)

				} else {

					termbox.SetCell(home.X_start+i, home.Y_start+index*3+2, chac, termbox.ColorWhite, termbox.ColorDefault)
				}
			}
		}
		for i, chac := range " -> " + task.Deedline {
			termbox.SetCell(home.X_start+i+len(task.Title)+3, home.Y_start+index*3+2, chac, termbox.ColorDefault, termbox.ColorDefault)
		}
		for i := home.Y_end - 4; i > home.Y_start; i-- {
			termbox.SetCell(40, i, '|', termbox.ColorWhite, termbox.ColorDefault)
		}
	}

}

type AddTask struct {
	cursorX int
	cursorY int
	step    []string
	qu      []string
	laststr string
}

var ObjAddTask = AddTask{qu: []string{
	"Write Title : ",
	"Write Dis : ",
	"Write Deed : ",
	"Write AC : ",
}}

func AddToTask(terminal *TerminalWriter) {
	task := Task{Title: ObjAddTask.step[0], Dic: ObjAddTask.step[1], Deedline: ObjAddTask.step[2], Ac: ObjAddTask.laststr, Function: RunInfoTask}
	Clear(terminal.Y_start+1, 0, terminal.Width, terminal.Height)
	terminal.Writer = false
	tasks = append(tasks, task)
	addFile()


}

func addHomeTask(terminal *TerminalWriter, key any) {
	if key == nil {

		for i := 0; i < terminal.Width; i++ {
			termbox.SetCell(i, terminal.Height-3, '_', termbox.ColorWhite, termbox.ColorDefault)

		}
		for i := 0; i < terminal.Width; i++ {
			for j := terminal.Height - 2; j <= terminal.Height; j++ {

				termbox.SetCell(i, terminal.Height-2, ' ', termbox.ColorWhite, termbox.ColorDefault)
			}
		}
		for i, v := range ObjAddTask.qu[len(ObjAddTask.step)] {
			termbox.SetCell(i, terminal.Height-2, v, termbox.ColorWhite, termbox.ColorDefault)

		}
		termbox.SetCursor(len(ObjAddTask.qu[len(ObjAddTask.step)])+1, terminal.Height-2)

	} else {
		qu := ObjAddTask.qu[len(ObjAddTask.step)]

		if key == -2 {

			ObjAddTask.step = append(ObjAddTask.step, ObjAddTask.laststr)
			for i := 0; i < terminal.Width; i++ {
				termbox.SetCell(i, terminal.Height-2, ' ', termbox.ColorWhite, termbox.ColorDefault)

			}
			termbox.SetCursor(len(ObjAddTask.qu[len(ObjAddTask.step)])+1, terminal.Height-2)
			for i, v := range ObjAddTask.qu[len(ObjAddTask.step)] {
				termbox.SetCell(i, terminal.Height-2, v, termbox.ColorWhite, termbox.ColorDefault)

			}
			ObjAddTask.cursorX = 0
			if len(ObjAddTask.step)+1 == len(ObjAddTask.qu) {
				for i := 0; i < terminal.Width; i++ {
					termbox.SetCell(i, terminal.Height-3, ' ', termbox.ColorWhite, termbox.ColorDefault)

				}
				AddToTask(terminal)
				ObjAddTask = AddTask{qu: []string{
					"Write Title : ",
					"Write Dis : ",
					"Write Deed : ",
					"Write AC : "}}
			}
			ObjAddTask.laststr = ""
			termbox.HideCursor()
			return
		}

		if key == -1 {
			for i, _ := range ObjAddTask.laststr {
				termbox.SetCell(i+len(qu), terminal.Height-2, ' ', termbox.ColorWhite, termbox.ColorDefault)

			}
			ObjAddTask.cursorX--
			termbox.SetCursor(ObjAddTask.cursorX+len(qu), terminal.Height-2)
			ObjAddTask.laststr = ObjAddTask.laststr[:len(ObjAddTask.laststr)-1]
		} else {

			ch := key.(int)

			ObjAddTask.laststr += string(rune(ch))
			ObjAddTask.cursorX++
			termbox.SetCursor(ObjAddTask.cursorX+len(qu), terminal.Height-2)
		}

		for i, v := range qu {
			termbox.SetCell(i, terminal.Height-2, v, termbox.ColorWhite, termbox.ColorDefault)

		}
		for i, v := range ObjAddTask.laststr {
			termbox.SetCell(i+len(qu), terminal.Height-2, v, termbox.ColorWhite, termbox.ColorDefault)

		}

	}
}

func deleteTask(terminal *TerminalWriter, a any) {
	tasksupdate := []Task{}

	for _, v := range tasks {
		if v.Node.Value.Lable != terminal.NodeHover.Value.Lable {
			tasksupdate = append(tasksupdate, v)
		}
		if v.Node.Edge["botton"].Value.Lable == terminal.NodeHover.Value.Lable {
			v.Node.Edge["botton"] = nil
		}
		if v.Node.Edge["top"].Value.Lable == terminal.NodeHover.Value.Lable {
			v.Node.Edge["top"] = nil
		}

	}

	nodelist := []*Node{}

	for i, v := range terminal.NodeList {
		lastlen := 3 + i
		length := 3
		maxlen := len(tasks) + length
		if v.Edge["botton"] != nil && v.Edge["botton"].Value.Lable == terminal.NodeHover.Value.Lable {

			if i+2 <= len(terminal.NodeList)-1 {
				v.Edge["botton"] = terminal.NodeList[lastlen+2]
			} else {
				v.Edge["botton"] = terminal.NodeList[length]

			}
		}
		if v.Edge["top"] != nil && v.Edge["top"].Value.Lable == terminal.NodeHover.Value.Lable {
			if lastlen-2 >= length {
				v.Edge["top"] = terminal.NodeList[lastlen-2]
			} else {
				v.Edge["top"] = terminal.NodeList[maxlen-1]

			}
		}
		if v.Value.Lable != terminal.NodeHover.Value.Lable {
			nodelist = append(nodelist, v)
		} else {
			if lastlen-1 >= length && lastlen-1 < maxlen {
				terminal.NodeHover = terminal.NodeList[lastlen-1]
			} else {
				if lastlen-1 >= maxlen {
					terminal.NodeHover = terminal.NodeList[length]
				} else {
					terminal.NodeHover = terminal.NodeList[lastlen+1]
				}

			}
		}
	}
	Clear(3, 0, 45, terminal.Height)
	terminal.NodeList = nodelist
	tasks = tasksupdate

	addFile()
}
