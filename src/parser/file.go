package parser

import "io"
import "os"
import "log"
import "bufio"

type FILE struct {
	Path	string
	Fd		*os.File
	Reader	*bufio.Reader
}

func (this *FILE) Init(path string) {
	var err error
	this.Fd, err = os.Open(os.Args[1])
	if (err != nil) { log.Fatal(err) }
	this.Reader = bufio.NewReader(this.Fd)
}

func (this *FILE) getNextLine() (string, bool) {
	line, err := this.Reader.ReadString('\n')
	if (err == io.EOF) { return string(line), true }
	if (err != nil) { log.Fatal(err) }

	return string(line), false
}
