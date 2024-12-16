package terminalpage


func (g *TerminalWriter) Add(edge map[string]*Node, parent *Node, value *Value, runInHover bool, click map[string]FunctionTp) *Node {
	var newNode = &Node{}
	if parent != nil {
		newNode = &Node{Level: parent.Level + 1, Parent: parent, Children: nil, Value: value, Edge: edge, RunInHover: runInHover, Click: click}
		for _, v := range g.NodeList {
			if v == parent {
				v.Children = append(v.Children, newNode)
			}
		}

	} else {
		newNode = &Node{Level: 0, Parent: nil, Children: nil, Value: value, Edge: edge, RunInHover: runInHover, Click: click}
	}
	for key, val := range edge {
		switch key {
		case "top":
			val.Edge["botton"] = newNode
		case "botton":
			val.Edge["top"] = newNode
		case "left":
			val.Edge["right"] = newNode
		case "right":
			val.Edge["left"] = newNode
		default:
			panic("Your Have Problem In Add Graph :" + key)
		}

	}

	g.NodeList = append(g.NodeList, newNode)


	return newNode
}
func (g *TerminalWriter) Clear() {
	g.NodeList = []*Node{}
}
func (g *TerminalWriter) Show(level int) []*Value {
	var arr []*Value
	for _, v := range g.NodeList {
		if v.Level == level {
			arr = append(arr, v.Value)
		}
	}
	return arr
}
