package main

import "strings"
import "fmt"
import "os"
import "log"

const (
	CONTINUE = 1
	BREAK = 2
)

var cmds = map[string]func([]string) (any, error) {
	"mtllib" :	func(args []string) (any, error) { return mlib(args) },
	"o":		func(args []string) (any, error) { return o(args) },
	"v":		func(args []string) (any, error) { return v(args) },
	"f":		func(args []string) (any, error) { return f(args) },
}

func checkLine(line string) ([]string, int8) {
	tokenized := strings.Fields(line)
	//fmt.Println(tokenized)

	if (len(tokenized)) <= 0 {
		fmt.Println("[INFO] Empty line")
		return tokenized, BREAK
	}
	if (tokenized[0] == "#") { return tokenized, CONTINUE }
	//if (len(tokenized) != 4) {
	//	if (len(tokenized)) > 4 { fmt.Println("Too much arguments on line") 
	//	} else { fmt.Println("Missing arguments on line") }
	//	return tokenized, CONTINUE
	//}
	return tokenized, 0
}

func parseObj(file *FILE) {
	for {
		line, eof := file.getNextLine()
		tokenized, instr := checkLine(line)

		if (instr == CONTINUE) { continue } else if (instr == BREAK) { break }
		fn := cmds[tokenized[0]]
		if (fn != nil) { fn(tokenized) } else { fmt.Println("[WARNING][", tokenized[0], "] Unsupported command") }
		if (eof) { break }
	}
}

func main() {
	if len(os.Args) != 2 { log.Fatal("File required has argument") }

	var file FILE
	file.init(os.Args[1])
	defer file.fd.Close()
	parseObj(&file)
}
