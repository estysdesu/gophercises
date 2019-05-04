package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	fileLoc := flag.String("csv", "./problems.csv", "the path to the csv problems file")
	timeLimit := flag.Int("time", 30, "the time limit for the quiz")
	random := flag.Bool("rand", false, "randomize the questions")
	flag.Parse()

	// print out quiz settings
	fmt.Printf("\nSettings:\nFile Location: %v\nTime Limit: %v\nRandomize: %v\n\n", *fileLoc, *timeLimit, *random)

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

	// if random is true then shuffle the array
	if *random == true {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(csvData), func(i, j int) {
			csvData[i], csvData[j] = csvData[j], csvData[i]
		})
	}

	// ask if the user is ready
	fmt.Print("Ready? (hit 'Enter' to start)")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// create time limit channel
	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))
	usrFin := make(chan bool)

	// iterate the questions and check their output
	var usrAns string
	correct := 0
	buffScanner := bufio.NewScanner(os.Stdin)
	go func() {
		for indx, data := range csvData {
			qu, ans := data[0], data[1]

			fmt.Printf("Problem %d: %s = ", indx+1, qu)
			buffScanner.Scan()
			usrAns = buffScanner.Text()

			usrAns = strings.TrimSpace(usrAns)
			usrAns = strings.ToLower(usrAns)
			ans = strings.ToLower(ans)

			if usrAns == ans {
				correct++
			}
		}
		usrFin <- true
	}()

	select {
	case <-usrFin:
	case <-timer.C:
		fmt.Println("\nTime's up!")
	}

	fmt.Printf("Score: %d, Questions: %d, Percentage: %f%%\n", correct, len(csvData), math.Round(100*(float64(correct)/float64(len(csvData)))))

}
