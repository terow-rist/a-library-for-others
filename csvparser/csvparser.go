package csvparser

import (
	"errors"
	"io"
)

// Struct & interface
type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	GetField(n int) (string, error)
	GetNumberOfFields() int
}

var (
	ErrQuote      = errors.New("excess or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")
)

type DataCSVParser struct {
	line   string
	fields []string
}

// Interface methods:
func (c *DataCSVParser) ReadLine(r io.Reader) (string, error) {
	var buffer []byte
	var insideQuotes bool
	for {
		temp := make([]byte, 1)
		_, err := r.Read(temp)
		if err != nil {
			if err == io.EOF {
				if insideQuotes {
					return "", ErrQuote
				}
				if len(buffer) > 0 {
					break
				}
				return "", io.EOF
			}
			return "", err
		}

		if temp[0] == '"' {
			insideQuotes = !insideQuotes
		}

		if !insideQuotes {
			if temp[0] == '\n' {
				break
			}
		}

		buffer = append(buffer, temp[0])

	}

	line := quoteFix(string(buffer))
	c.fields = separateLine(line)
	c.line = line
	return line, nil
}

func (c *DataCSVParser) GetField(n int) (string, error) {
	fields := c.fields
	if n < 0 || n >= len(fields) {
		return "", ErrFieldCount
	}
	field := fields[n]
	if len(field) > 1 && field[0] == '"' && field[len(field)-1] == '"' {
		return field[1 : len(field)-1], nil
	}

	if len(field) > 2 && field[0] == '"' {
		if (field[len(field)-1] == '\n' || field[len(field)-1] == '\r') && field[len(field)-2] == '"' {
			return field[1 : len(field)-2], nil
		}
	}
	return field, nil
}

func (c *DataCSVParser) GetNumberOfFields() int {
	// if c.
	return len(c.fields)
}

// Utils function \ /
func separateLine(line string) []string {
	var tempStr string
	var fields []string
	var openQuotes bool

	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '"':
			openQuotes = !openQuotes
			tempStr += string(line[i])

		case ',':
			if openQuotes {
				tempStr += string(line[i])
			} else {
				fields = append(fields, tempStr)
				tempStr = ""
			}

		default:
			tempStr += string(line[i])
		}
	}

	// Add the last field even if it's empty
	fields = append(fields, tempStr)

	return fields
}

func quoteFix(line string) string {
	var newLine string
	for i := 0; i < len(line); i++ {
		if line[i] == '"' {
			if i == 0 {
				newLine += "\""
				continue
			} else if line[i-1] == '"' {
				continue
			}
		}
		newLine += string(line[i])
	}
	return newLine
}
