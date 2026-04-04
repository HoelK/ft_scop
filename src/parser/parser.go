package parser

import "fmt"

const (
	CONTINUE = 1
	BREAK = 2
)

var cmdsSupported = map[string]bool {
	"o" :		true,
	"g" :		false,
	"f" :		true,
	"v" :		true,
	"s" :		true,
	"mtllib" :	true,
	"usemtl" :	true,
}

func ParseObj(file *FILE) (Data) {
	var data	Data
	var pTree	Node
	var curr	*Node

	pTree = initTree()
	curr = &pTree
	data.Path = file.Path

	for {
		line, eof := file.getNextLine()
		tokenized, instr := checkLine(line, eof)

		if (instr == CONTINUE) { continue } else if (instr == BREAK) {
			if (curr.Branchs["eof"] == nil) { fmt.Println("[ERROR][EOF] Unexpected - Incomplete object") }
			break
		}
		_, ok := cmdsSupported[tokenized[0]]

		if (!ok) {
			fmt.Println("[ERROR][", tokenized[0], "] Undefined commmand")
			continue
		}
		if (curr.Branchs[tokenized[0]] == nil) {
			fmt.Println("[ERROR][", tokenized[0], "] Command in Wrong Section")
			continue
		}
		curr = curr.Branchs[tokenized[0]]
		if (curr.fn == nil) {
			fmt.Println("[WARNING][", tokenized[0], "] Unsupported command")
			continue
		}

		err	:= curr.fn(&data, tokenized)
		if (err != nil)	{ fmt.Println("[ERROR]", err) }
		if (eof) {
			if (curr.Branchs["eof"] == nil) { fmt.Println("[WARNING][EOF] Unexpected - Incomplete object") }
			break
		}
	}

	return data
}
