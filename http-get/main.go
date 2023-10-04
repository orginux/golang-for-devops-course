package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	content string
	pos     int
}

func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.pos < len(m.content) {
		fmt.Println("reading...")
		n := copy(p, m.content[m.pos:m.pos+1])
		m.pos++
		return n, nil
	}
	return 0, io.EOF
}

func main() {
	mySlowReaderInstance := &MySlowReader{
		content: "oleg",
	}

	out, err := io.ReadAll(mySlowReaderInstance)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(out))
}
