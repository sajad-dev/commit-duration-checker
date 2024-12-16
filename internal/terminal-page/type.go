package terminalpage

type TerminalWriter struct {
	Writer          bool
	X_start         int
	X_end           int
	Y_start         int
	Y_end           int
	Width           int
	Height          int
	NodeActive      *Node
	NodeChildActive *Node
	NodeHover       *Node
	NodeList        []*Node
}
type FunctionTp func(*TerminalWriter, interface{})
type Node struct {
	Level      int
	Parent     *Node
	RunInHover bool
	Click      map[string]FunctionTp
	Children   []*Node
	Value      *Value
	Edge       map[string]*Node
}

type Value struct {
	Lable  string
	Val    interface{}
	Output FunctionTp
}
