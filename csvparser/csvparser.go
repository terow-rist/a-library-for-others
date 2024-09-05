package csvparser

import (
	"errors"
	"fmt"
	"io"
)

// Struct & interface
type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	GetField(n int) (string, error)
	GetNumberOfFields() int
}

var (
	ErrField      = errors.New("unexpected value in field")
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
//
//	if prevChar == '"' && temp[0] == ',' && insideQuotes {
//		return "", ErrQuote
//	}
//
//	if temp[0] == ',' && insideQuotes {
//		return "", ErrQuote
//	}
func (c DataCSVParser) ReadLine(r io.Reader) (string, error) {
	var buffer []byte
	var prevChar byte
	var insideQuotes bool
	var incorrectField string
	for {
		temp := make([]byte, 1)
		_, err := r.Read(temp)
		if err != nil {
			if err == io.EOF {
				if insideQuotes {
					fmt.Println("SOMETHING")
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
		if temp[0] == ',' && insideQuotes && prevChar == '"' {
			fmt.Println("End qoute ,")
			return "", ErrQuote
		}
		if !insideQuotes {
			if temp[0] == '\n' {
				incorrectField += string(buffer) + string(temp[0])
				break
			}
		}

		buffer = append(buffer, temp[0])
		prevChar = temp[0]

	}

	line := string(buffer)
	if incorrectField == "\r\n" || incorrectField == "\n" {
		return "", ErrField
	}
	c.fields = separateLine(line)
	c.line = line
	return line, nil
}

// Utils function \ /
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
