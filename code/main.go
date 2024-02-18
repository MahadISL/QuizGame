package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Printf("working")
	file, err := os.Open("problems.csv")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	record, err := reader.ReadAll()

	for _, value := range record {
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Print("err")
		}

		fmt.Print(value)
	}
}
