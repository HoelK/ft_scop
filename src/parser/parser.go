package parser

import "fmt"

const (
	CONTINUE = 1
	BREAK = 2
)

var cmds = map[string]func(*Data, []string) error {
	"mtllib" :	func(data *Data, args []string) error { return mlib(data, args) },
	"o":		func(data *Data, args []string) error { return o(data, args) },
	"v":		func(data *Data, args []string) error { return v(data, args) },
	"f":		func(data *Data, args []string) error { return f(data, args) },
}

func printData(data *Data) {
	fmt.Println("[LOG] Data info :")
	for i := 0; i < len(data.Objs); i++ {
		fmt.Println(data.Objs[i].Name)
	}
}

func ParseObj(file *FILE) (Data) {
	var data Data
	data.Path = file.Path

	for {
		line, eof := file.getNextLine()
		tokenized, instr := checkLine(line, eof)

		if (instr == CONTINUE) { continue } else if (instr == BREAK) { break }
		fn := cmds[tokenized[0]]
		if (fn == nil) {
			fmt.Println("[WARNING][", tokenized[0], "] Unsupported command")
			continue
		}
		err := fn(&data, tokenized)
		if (err != nil)	{ fmt.Println("[ERROR]", err) }
		if (eof)		{ break }
	}

	return data
}
