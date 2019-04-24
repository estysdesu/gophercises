package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"log"
	"flag"
	"bufio"
)

func main() {

	fileLoc := flag.String("csv", "./problems.csv", "the path to the csv problems file")
	// timeLimit := flag.Int("time", 30, "the time limit for the quiz")
	// random := flag.Bool("rand", false, "randomize the questions")
	flag.Parse()

	// open the file and defer close
	csvFile, err := os.Open(*fileLoc)
	if err != nil {
		log.Println(err)
	}
	defer csvFile.Close()

	// create new csv reader then read the files content
	csvReader := csv.NewReader(csvFile)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Println(err)
	}

	// make the q&a slices
	q := make([]string, 0, len(csvData))
	a := make([]string, 0, len(csvData))
	for _, data := range csvData {
		q = append(q, data[0])
		a = append(a, data[1])
	}

	var usrAns string
	correct := 0
	buffScanner := bufio.NewScanner(os.Stdin)
	for indx, data := range csvData {
		fmt.Printf("Problem %d: %s = ", indx+1, data[0])
		buffScanner.Scan()
		usrAns = buffScanner.Text()

		if usrAns == data[1] {
			correct++
		}
	}

	fmt.Printf("Score: %d, Questions: %d, Percentage: %f", correct, len(csvData), float32(correct)/float32(len(csvData)))


}
