package caddycbrotli

import (
	"errors"
	"io"

	"github.com/google/brotli/go/cbrotli"
)

var errNoWriter = errors.New("no writer")

type encoder struct {
	writer  *cbrotli.Writer
	options cbrotli.WriterOptions
}

func newEncoder(options cbrotli.WriterOptions) *encoder {
	return &encoder{nil, options}
}

func (e *encoder) Close() error {
	if e.writer == nil {
		return errNoWriter
	}

	err := e.writer.Close()
	e.writer = nil

	return err
}

func (e *encoder) Flush() error {
	if e.writer == nil {
		return errNoWriter
	}

	return e.writer.Flush()
}

func (e *encoder) Write(p []byte) (n int, err error) {
	if e.writer == nil {
		return 0, errNoWriter
	}

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
