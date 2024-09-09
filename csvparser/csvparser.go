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

		// Handle the quotes logic
		if char == '"' {
			insideQuotes = !insideQuotes
		}

		// Handle line breaks outside of quotes
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
	c.fields = separateLine(string(buffer))
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
			// Handle opening and closing of quotes
			if openQuotes && i+1 < len(line) && line[i+1] == '"' {
				// If it's a double quote within a quoted field, add a single quote
				tempStr += `"`
				i++ // Skip the next quote
			} else {
				// Toggle the openQuotes state
				openQuotes = !openQuotes
			}

		case ',':
			// If we're inside quotes, consider the comma as part of the field
			if openQuotes {
				tempStr += string(char)
			} else {
				// If not inside quotes, the comma marks the end of a field
				fields = append(fields, tempStr)
				tempStr = ""
			}

		default:
			// Add the character to the current field buffer
			tempStr += string(char)
		}
	}

	// Append the last field after the loop finishes
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
