package parser

import "fmt"
import "errors"
import "strings"

func checkArgs(args []string, amount int) error {
	if len(args) > amount { return errors.New("[" + args[0] + "] Too much arguments (" + string(amount) + " needed)") }
	if len(args) < amount { return errors.New("[" + args[0] + "] Missing arguments (" + string(amount) + " needed)") }
	return nil
}


func checkLine(line string, eof bool) ([]string, int8) {
	tokenized := strings.Fields(line)

	if (len(tokenized)) <= 0 {
		if (eof) {
			fmt.Println("[INFO] End Of File")
			return tokenized, BREAK
		}
		fmt.Println("[INFO] Empty line")
		return tokenized, CONTINUE
	}
	if (tokenized[0] == "#") { return tokenized, CONTINUE }
	return tokenized, 0
}

