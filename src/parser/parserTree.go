package parser

type Node struct {
	Symbol	string
	Branchs	map[string]*Node
	fn		func(*Data, []string) error
}

func initTree() Node {
	var tree Node
	var tmp *Node

	tree.Symbol = "root"
	tree.Branchs			= make(map[string]*Node)
	tree.Branchs["eof"]		=	&Node{"eof", nil, nil}
	tree.Branchs["o"]		=	&Node{"o", make(map[string]*Node), o}
	tree.Branchs["mtllib"]	=	&Node{"mtllib", make(map[string]*Node), mlib}

	tmp = tree.Branchs["mtllib"]
	tmp.Branchs["o"]	= tree.Branchs["o"]

	tmp = tree.Branchs["o"]
	tmp.Branchs["o"] = tmp
	tmp.Branchs["v"] = &Node{"v", make(map[string]*Node), v}

	tmp = tmp.Branchs["v"]
	tmp.Branchs["usemtl"] = &Node{"usemtl", make(map[string]*Node), nil}

	tmp = tmp.Branchs["usemtl"]
	tmp.Branchs["s"] = &Node{"s", make(map[string]*Node), nil}

	tmp = tmp.Branchs["s"]
	tmp.Branchs["f"] = &Node{"f", make(map[string]*Node), f}

	tmp = tmp.Branchs["f"]
	tmp.Branchs["eof"]	= &Node{"eof", nil, nil}

	return tree
}
