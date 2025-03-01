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
	line        string
	fields      []string
	prevChar    byte
	eatenSlashR bool
}

// Interface methods:
func (c *DataCSVParser) ReadLine(r io.Reader) (string, error) {
	var buffer []byte
	var insideQuotes bool
	for {
		if c.eatenSlashR {
			buffer = append(buffer, c.prevChar)
			c.eatenSlashR = false
		}
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

		char := temp[0]

		if char == '"' {
			insideQuotes = !insideQuotes
		}

		if !insideQuotes {
			if char == '\r' {
				temp = make([]byte, 1)
				_, err := r.Read(temp)
				if err != nil {
					if err == io.EOF {
						if len(buffer) > 0 {
							break
						}
						return "", io.EOF
					}
					return "", err
				}
				if temp[0] == '\n' {
					break
				} else {
					c.prevChar = temp[0]
					c.eatenSlashR = true
					break
				}
			}
			if char == '\n' {
				break
			}
		}

		buffer = append(buffer, char)
		c.prevChar = char
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

	field = trimWhitespace(field)

	if len(field) > 1 && field[0] == '"' && field[len(field)-1] == '"' {
		return field[1 : len(field)-1], nil
	}

	return field, nil
}

func (c *DataCSVParser) GetNumberOfFields() int {
	return len(c.fields)
}

// Utils function \ /
func separateLine(line string) []string {
	var tempStr string
	var fields []string
	var openQuotes bool

	for i := 0; i < len(line); i++ {
		char := line[i]

		switch char {
		case '"':
			openQuotes = !openQuotes
			tempStr += string(line[i])

		case ',':
			if openQuotes {
				tempStr += string(char)
			} else {
				fields = append(fields, tempStr)
				tempStr = ""
			}

		default:
			tempStr += string(char)
		}
	}

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

func trimWhitespace(s string) string {
	start, end := 0, len(s)-1

	for start <= end && (s[start] == ' ' || s[start] == '\n' || s[start] == '\r' || s[start] == '\t') {
		start++
	}

	for end >= start && (s[end] == ' ' || s[end] == '\n' || s[end] == '\r' || s[end] == '\t') {
		end--
	}

	return s[start : end+1]
}
