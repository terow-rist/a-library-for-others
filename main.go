package main

import (
	"fmt"
	"io"
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

	csvparser := csvparser.DataCSVParser{}
	ind := 1
	for {
		line, err := csvparser.ReadLine(file)
		if err != nil {
			if err == io.EOF {
				fmt.Println(csvparser.GetNumberOfFields())
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}
		fmt.Println("Line", ind, ":", line)
		ind++
		var n int
		fmt.Scanf("%d", &n)
		field, errf := csvparser.GetField(n)
		if errf != nil {
			fmt.Println("Error reading field VIA line:", errf)
			return
		}
		fmt.Println("Got field: ", field)
		fmt.Println("GetNumbeROfFields", csvparser.GetNumberOfFields())
		fmt.Println("-----------------------")
	}
}
