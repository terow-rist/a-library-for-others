package main

import (
	"fmt"
	"os"

	"a-library-for-others/csvparser"
)

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	csvparser := csvparser.CSVParser{}

	csvparser.ReadLine(file)
	// var csvparser CSVParser = YourCSVParser{}

	// for {
	// 	line, err := csvparser.ReadLine(file)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		fmt.Println("Error reading line:", err)
	// 		return
	// 	}
	// }
}
