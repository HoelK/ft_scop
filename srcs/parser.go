package main

import "strings"
import "fmt"
import "os"
import "log"

const (
	CONTINUE = 1
	BREAK = 2
)

func v(args []string) ([3]float32, error) {
	var a [3]float32
	a[0] = 0
	if len(args) != 0 { return a, nil }
	return a, nil
}

func f(args []string) ([]int32, error) {
	a := make([]int32, 3)
	a[0] = 0
	if len(args) != 0 { return a, nil }
	return a, nil
}

var cmds = map[string]func([]string) (any, error) {
	"v": func(args []string) (any, error) { return v(args) },
	"f": func(args []string) (any, error) { return f(args) },
}

func checkLine(line string) ([]string, int8) {
	tokenized := strings.Fields(line)
	fmt.Println(tokenized)

	if (len(tokenized)) <= 0 {
		fmt.Println("Empty line")
		return tokenized, BREAK
	}
	if (tokenized[0] == "#") { return tokenized, CONTINUE }
	if len(tokenized) != 4 {
		if (len(tokenized)) > 4 { fmt.Println("Too much arguments on line") 
		} else if (len(tokenized)) < 4 { fmt.Println("Missing arguments on line")}
		return tokenized, CONTINUE
	}
	return tokenized, 0
}

func parseObj(file *FILE) {
	for {
		line, eof := file.getNextLine()
		tokenized, instr := checkLine(line)

		if instr == CONTINUE { continue } else if instr == BREAK { break }
		cmds[tokenized[0]](tokenized)
		if eof { break }
	}
}

func main() {
	if len(os.Args) != 2 { log.Fatal("File required has argument") }

	var file FILE
	file.init(os.Args[1])
	defer file.fd.Close()
	parseObj(&file)
}
