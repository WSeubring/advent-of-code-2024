package main

import (
	"bufio"
	"os"
)

type FileLineReader struct {
	Reader *bufio.Reader
}

func NewFileLineReader(path string) (*FileLineReader, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return &FileLineReader{
		Reader: bufio.NewReader(file),
	}, nil
}

func (r *FileLineReader) ReadLine() (string, error) {
	return r.Reader.ReadLine()
}
