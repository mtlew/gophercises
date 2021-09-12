package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
)

type question struct {
	question string
	answer   string
}

func main() {
	csvFilenameDefault := "quiz.csv"
	csvFilename := flag.String("csv", csvFilenameDefault, "Questions filename ("+csvFilenameDefault+" by default)")
	//timeLimitSeconds := flag.Int("limit", 30, "Time limit in seconds")
	flag.Parse()

	questions, err := csv2structs(*csvFilename)

	if err != nil {
		exit(err.Error())
	}

	//timer := time.NewTimer(time.Duration(*timeLimitSeconds) * time.Second)

	correct := 0
	for i, question := range questions {
		fmt.Printf("Question #%d: %s = \n", i+1, question.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == question.answer {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(questions))
}

func csv2structs(filename string) ([]question, error) {

	file, err := os.Open(filename)

	if err != nil {
		return nil, errors.New("Can't open csv file" + filename)
	}

	reader := csv.NewReader(file)
	reader.Comma = ';'

	var questions []question

	for {
		line, err := reader.Read()

		if err != nil {
			break
		}

		questions = append(questions, question{
			question: line[0],
			answer:   line[1],
		})
	}

	return questions, nil
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
