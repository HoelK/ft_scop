package parser

import "fmt"

var mtlCmds = map[string]func (*Material, []string) error {
	"Ns"		:	func(mtl *Material, args []string) error { return Ns(mtl, args) },
	"Ka"		:	func(mtl *Material, args []string) error { return Ka(mtl, args) },
	"Kd"		:	func(mtl *Material, args []string) error { return Kd(mtl, args) },
	"Ks"		:	func(mtl *Material, args []string) error { return Ks(mtl, args) },
	"Ni"		:	func(mtl *Material, args []string) error { return Ni(mtl, args) },	
	"d"			:	func(mtl *Material, args []string) error { return d(mtl, args) },
	"illum"		:	func(mtl *Material, args []string) error { return illum(mtl, args) },
}

func parseMtl(data *Data, name string) {
	var file	FILE
	var err		error
	var mtl		*Material

	file.Init(name)
	for {
		line, eof			:= file.getNextLine()
		tokenized, instr	:= checkLine(line, eof)

		if (instr == CONTINUE) { continue } else if (instr == BREAK) { break }
		if (tokenized[0] == "newmtl") {
			mtl, err = newmtl(data, tokenized)
			continue
		} else {
			fn := mtlCmds[tokenized[0]]
			if (fn == nil) {
				fmt.Println("[WARNING][", tokenized[0], "] Unsupported command")
				continue
			}
			err = fn(mtl, tokenized)
		}

		if (err != nil)	{ fmt.Println("[ERROR]", err) }
		if (eof)		{ break }
	}
}
