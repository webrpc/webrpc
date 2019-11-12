package schema

import (
	"io"
)

func NewReader(r io.Reader, path string) *Reader {
	return &Reader{File: path, r: r}
}

type Reader struct {
	r io.Reader

	File string
}

func (r *Reader) Read(p []byte) (n int, err error) {
	return r.r.Read(p)
}

var _ = io.Reader(&Reader{})
