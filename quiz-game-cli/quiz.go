package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)


func main() {
	filepath := flag.String("f", "problems.csv", "Path of CSV file")
	time_limit := flag.Int("t", 30, "Time limit for quiz (in seconds)")

	_ = time_limit

	flag.Parse()

	file, err := os.Open(*filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close();
	r := csv.NewReader(file)
	correct := 0
	number_of_questions := 0;
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		number_of_questions++;
		fmt.Print("What is " + line[0] + "? ")
		var response string
		fmt.Scanln(&response)
		ans := strings.Trim(line[1], " /*")
		if (response == ans) {
			correct++
		}
	}

	fmt.Printf("\nYou have answered %v out of %v questions correctly", correct, number_of_questions)
}