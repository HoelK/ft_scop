package parser

import "io"
import "os"
import "log"
import "bufio"
import "strings"

type FILE struct {
	Name		string
	Path		string
	Fullname	string
	Fd			*os.File
	Reader		*bufio.Reader
}

func (this *FILE) Init(path string) {
	var err error

	this.Fullname = path
	spl := strings.Split(path, "/")
	this.Name = spl[len(spl) - 1]
	spl = spl[:len(spl) - 1]
	this.Path = strings.Join(spl, "/")
	this.Path += "/"

	this.Fd, err = os.Open(path)
	if (err != nil) { log.Fatal(err) }
	this.Reader = bufio.NewReader(this.Fd)
}

func (this *FILE) getNextLine() (string, bool) {
	line, err := this.Reader.ReadString('\n')
	if (err == io.EOF) { return string(line), true }
	if (err != nil) { log.Fatal(err) }

	return string(line), false
}
