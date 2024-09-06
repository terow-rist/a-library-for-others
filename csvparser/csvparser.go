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
	ErrQuote      = errors.New("excess or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")
)

type DataCSVParser struct {
	fields []string
	line   string
}

// Interface methods:
func (c DataCSVParser) ReadLine(r io.Reader) (string, error) {
	var buffer []byte
	var prevChar byte
	var insideQuotes bool
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
				break
			}
		}

		buffer = append(buffer, temp[0])
		prevChar = temp[0]

	}

	line := quoteFix(string(buffer))
	// if !quoteCheck(line) {
	// 	fmt.Println("wf")
	// 	return "", ErrQuote
	// }
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

// func quoteCheck(line string) bool {
// 	var field string
// 	var openQuote bool
// 	for i := 0; i < len(line); i++ {
// 		if line[i] == '"' {
// 			openQuote = !openQuote
// 		}
// 		if line[i] == ',' && !openQuote {
// 			for j, char := range field {
// 				if char == '"' && j != 0 && j != len(field)-1 {
// 					return false
// 				}
// 			}
// 			field = ""
// 			continue
// 		}
// 		field += string(line[i])
// 	}
// 	if len(field) > 0 {
// 		for j, char := range field {
// 			if char == '"' && j != 0 && j != len(field)-1 {
// 				return false
// 			}
// 		}
// 	}
// 	if openQuote {
// 		return false
// 	}
// 	return true
// }
