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
	fields []string
	line   string
	// quotes struct {
	// 	start bool
	// 	end   bool
	// }
}

// Interface methods:

func (c DataCSVParser) ReadLine(r io.Reader) (string, error) {
	var buffer []byte
	// var prevChar byte
	var insideQuotes bool
	for {
		temp := make([]byte, 1)
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

		// if prevChar == '"' && temp[0] == ',' && insideQuotes {
		// 	return "", ErrQuote
		// }
		if temp[0] == ',' && insideQuotes {
			return "", ErrQuote
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
		// prevChar = temp[0]
	}

	line := string(buffer)

	c.fields = separateLine(line)
	c.line = line
	return line, nil
}

// Utils function \ /
//
//	|
//	^
func separateLine(line string) []string {
	tempStr := ""
	fields := []string{}
	for i := 0; i < len(line); i++ {
		if line[i] == ',' {
			fields = append(fields, tempStr)
			tempStr = ""
		}
		tempStr += string(line[i])
	}
	if tempStr != "" {
		fields = append(fields, tempStr)
	}
	return fields
}

func invalidAmountQuotes(line string) bool {
	counter := 0
	for _, char := range line {
		if char == '"' {
			counter++
		}
	}
	if counter%2 != 0 {
		return true
	}
	return false
}
