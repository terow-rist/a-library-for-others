package csvparser

import (
	"errors"
	"io"
)

type CSVParser struct {
	line   string
	fields []string
}

var (
	ErrQuote      = errors.New("excess or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")
)

func (c CSVParser) ReadLine(r io.Reader) (string, error) {
	// start with io.Reader.Read
	return "", nil
}
