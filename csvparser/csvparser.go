package csvparser

import (
	"errors"
	"fmt"
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
	var buffer []byte
	for {
		// Create a small buffer to read byte by byte
		temp := make([]byte, 1)
		_, err := r.Read(temp)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return "", nil
		}

		// Collect bytes until a newline is found
		if temp[0] == '\n' {
			fmt.Println(string(buffer))
			buffer = []byte{} // Reset buffer after printing the line
		} else {
			buffer = append(buffer, temp[0])
		}
	}

	// Print any remaining data in the buffer
	if len(buffer) > 0 {
		fmt.Println(string(buffer))
	}
	return "", nil
}
