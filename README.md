# a-library-for-others
	package csvparser

import (
	"errors"
	"io"
)

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
	quotes struct {
		start bool
		end   bool
	}
}

// Field struct {
// 	values   []byte
// 	quoted   bool
// 	closedBy string
// }

// "", ""\n ""\r ""\r\n ""EOF

// insideQoutes := false

// if !c.checkQuotes() {
// 	return "", ErrQuote
// } else {
// 	c.quotes.end = false
// 	c.quotes.start = false
// }

func (c DataCSVParser) ReadLine(r io.Reader) (string, error) {
	var buffer []byte
	var prevChar byte
	for {
		temp := make([]byte, 1)
		_, err := r.Read(temp)
		if err != nil {
			if err == io.EOF {
				if len(buffer) > 0 {
					break
				}
				//
				return "", io.EOF
			}
			return "", err
		}

		if temp[0] == '\n' || (temp[0] == '\r' && prevChar != '\n') {
			//
			break
		}
		// if temp[0] == ',' {
		// 	if !c.checkQuotes() {
		// 		return "", ErrQuote
		// 	} else {
		// 		c.quotes.end = false
		// 		c.quotes.start = false
		// 	}
		// }

		buffer = append(buffer, temp[0])
		prevChar = temp[0]
	}

	line := string(buffer)

	c.fields = separateLine(line)
	c.line = line
	return line, nil
}

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

func (c DataCSVParser) checkQuotes() bool {if ch == '"' {
				
}
		if c.quotes.end {
			return false
		}
	}
	return true
}

func (c DataCSVParser) lineCheck(line string) {
	if line[0] != '"' || line[len(line)-1] != '"' {
		
	}
	for i := 0; i < len(line); i++ {
		if 
		if line[i] == '"' {
			if c.quotes.start {
				c.quotes.end = true
			}
			c.quotes.start = true
		}
	}
}

