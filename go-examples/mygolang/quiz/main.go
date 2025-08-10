package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
)

var score int = 0

func main() {

	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	// load the csv file

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Error opening CSV file: %s\n", *csvFilename))
	}

	// Read the CSV file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse reading CSV file: %s\n", *csvFilename))
	}

	problems := parseLines(lines)
	// Shuffle the problems
	shuffleProblems(problems)

	for i, p := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer) // by adding \n to the format string can take the whole line as input and spaces can be taken as well
		if answer == p.answer {
			score++
		}
	}
	fmt.Printf("Your score: %d out of %d\n", score, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func shuffleProblems(problems []problem) {
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})
}
