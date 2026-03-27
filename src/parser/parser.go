package parser

import "fmt"
//import "os"
//import "log"
import "strings"

const (
	CONTINUE = 1
	BREAK = 2
)

var cmds = map[string]func([]string) error {
	"mtllib" :	func(args []string) error { return mlib(args) },
	"o":		func(args []string) error { return o(args) },
	"v":		func(args []string) error { return v(args) },
	"f":		func(args []string) error { return f(args) },
}

func checkLine(line string) ([]string, int8) {
	tokenized := strings.Fields(line)

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
		if (fn == nil) {
			fmt.Println("[WARNING][", tokenized[0], "] Unsupported command")
			continue
		}
		err := fn(tokenized)
		if (err != nil) { fmt.Println("[ERROR]", err) }
		if (eof) { break }
	}
}

/*func main() {
	if len(os.Args) != 2 { log.Fatal("File required has argument") }

	var file FILE
	file.init(os.Args[1])
	defer file.fd.Close()
	parseObj(&file)
}*/
