```markdown
# A Library for Others

**Project Type:** Individual  
**Complexity:** Medium  
**Status:** In Progress  
**Deadline:** 09.09, 12:27  
**Max XP:** 2500 XP  
**User XP:** 0 XP

## Learning Objectives
- File Handling in Go
- CSV Parsing
- Interfaces
- Error Handling

## Abstract
This project involves building a CSV library in Go, focusing on key programming concepts such as interfaces, file handling, and error management. You'll design a library that simplifies CSV file handling, ensuring clean error reporting and proper resource management. The project aims to sharpen your skills in parsing, interface design, and efficient resource handling.

## Context
CSV files are a common format for storing tabular data. However, handling CSV files can be tricky due to inconsistencies in formatting and quoting conventions. This project will help you master CSV parsing in Go, with an emphasis on managing edge cases, such as complex data structures and error scenarios.

## Resources
- Go File I/O
- CSV Format
- Go Interfaces

## General Instructions
- Follow `gofumpt` formatting.
- Ensure the code compiles successfully and handles errors correctly.
- Only `io.Reader` is allowed.
- Create utility functions as needed.

## Mandatory Part

### CSVParser Interface
```go
type CSVParser interface {
    ReadLine(r io.Reader) (string, error)
    GetField(n int) (string, error)
    GetNumberOfFields() int
}

var (
    ErrQuote      = errors.New("excess or missing \" in quoted-field")
    ErrFieldCount = errors.New("wrong number of fields")
)
```

### ReadLine
- Reads one line from a CSV file.
- Strips the line terminator and returns the line.
- Handles errors related to quotes and line formatting.

### GetField
- Returns the nth field from the last line read by `ReadLine`.
- Handles quoting and field errors.

### GetNumberOfFields
- Returns the number of fields in the last line read.

## Testing
```go
func main() {
    file, err := os.Open("example.csv")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var csvparser CSVParser = YourCSVParser{}

    for {
        line, err := csvparser.ReadLine(file)
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("Error reading line:", err)
            return
        }
    }
}
```

## Support & Guidelines
- Test your code thoroughly with different file types and edge cases.
- Follow best practices in error handling and memory management.

## Author
**Adilyam Tilegenova**  
Software Developer at Doodocs.kz  
Email: adilyamt@gmail.com  
[GitHub](https://github.com)  
[LinkedIn](https://linkedin.com)
```