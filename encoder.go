package caddycbrotli

import (
	"io"

	"github.com/google/brotli/go/cbrotli"
)

type encoder struct {
	writer  *cbrotli.Writer
	options cbrotli.WriterOptions
}

func newEncoder(options cbrotli.WriterOptions) *encoder {
	return &encoder{nil, options}
}

func (e *encoder) Close() error {
	err := e.writer.Close()
	e.writer = nil

	return err
}

func (e *encoder) Flush() error {
	return e.writer.Flush()
}

func (e *encoder) Write(p []byte) (n int, err error) {
	return e.writer.Write(p)
}

// see https://github.com/google/brotli/issues/679
func (e *encoder) Reset(w io.Writer) {
	if e.writer != nil {
		e.writer.Close()
	}

	if w == nil {
		e.writer = nil
	} else {
		e.writer = cbrotli.NewWriter(w, e.options)
	}
}
