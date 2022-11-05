package main

import (
	"io"
)

type stdin struct {
	stdin    io.Reader
	_quitCmd string // use method quitCmd to get the quit command
}

func newStdin(quitCmd string, reader io.Reader) *stdin {
	if reader == nil {
		panic("the reader can not be nil")
	}

	if quitCmd == "" {
		quitCmd = "quit"
	}
	return &stdin{reader, quitCmd}
}
func (s *stdin) Read(b []byte) (int, error) {
	n, err := s.stdin.Read(b)
	bytesRead := b[:n]
	if string(bytesRead) == s.quitCmd() {
		return 0, io.EOF
	}
	return n, err
}
func (s *stdin) quitCmd() string {
	if s._quitCmd[len(s._quitCmd)-1] != '\n' {
		return s._quitCmd + "\n"
	}
	return s._quitCmd
}
