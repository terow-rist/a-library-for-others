package main

import (
	"a-library-for-others/csvparser"
	"fmt"
	"io"
	"os"
)

func main() {
	content := "John\r\n122,ofkomcs,escsc,,dmrdmrvr,,sefsvm, "
	CreateFile(content)
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	csvparser := csvparser.DataCSVParser{}
	ind := 1
	for {
		// fmt.Println(csvparser.GetNumberOfFields())
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
		fmt.Println("GetNumbeROfFields", csvparser.GetNumberOfFields())
		for i := 0; i < csvparser.GetNumberOfFields(); i++ {
			field, errf := csvparser.GetField(i)
			if errf != nil {
				fmt.Println("Error reading field VIA line:", errf)
				return
			}
			fmt.Println("Got field: ", field)

		}
		fmt.Println("-----------------------")

	}
}

func CreateFile(content string) {
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Always remember to close the file after you're done
	defer file.Close()

	// Write the string content to the file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
