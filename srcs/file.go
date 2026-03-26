package main

import "io"
import "os"
import "log"
import "bufio"

type FILE struct {
	path	string
	fd		*os.File
	reader	*bufio.Reader
}

func (this *FILE) init(path string) {
	var err error
	this.fd, err = os.Open(os.Args[1])
	if err != nil { log.Fatal(err) }
	this.reader = bufio.NewReader(this.fd)
}

func (this *FILE) getNextLine() (string, bool) {
	line, err := this.reader.ReadString('\n')
	if err == io.EOF { return string(line), true }
	if err != nil { log.Fatal(err) }

	return string(line), false
}
