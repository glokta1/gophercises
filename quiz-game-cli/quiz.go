package main

import (
	// "encoding/csv"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"flag"
)


func main() {
	filepath := flag.String("f", "problems.csv", "Path of CSV file")
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
		var response int
		fmt.Scanln(&response)
		ans, err := strconv.Atoi(line[1])
		if (err != nil) {
			panic(err)
		}
		if (response == ans) {
			correct++
		}
	}

	fmt.Printf("\nYou have answered %v out of %v questions correctly", correct, number_of_questions)
}