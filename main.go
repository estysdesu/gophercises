package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"log"
)

func main() {

	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	csvData, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(csvData)
	fmt.Println(len(csvData))

	q := make([]string, len(csvData))
	a := make([]string, len(csvData))
	for indx, data := range csvData {
		q[indx] = data[0]
		a[indx] = data[1]
		fmt.Sprintf("q = %s & a = %s", q, a)
	}


}
