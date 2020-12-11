package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	fmt.Printf("Result is %v \n", run())

	log.Printf("Code took %s", time.Since(start))
}

func run() int {
	var questionsAnswered int

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	questions := ""

	for scanner.Scan() {
		if scanner.Text() == "" {
			questionsAnswered += getQuestionsQuantity(questions)
			questions = ""
			continue
		}

		questions += scanner.Text()
	}

	questionsAnswered += getQuestionsQuantity(questions)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return questionsAnswered
}

func getQuestionsQuantity(qq string) int {
	var chars = map[string]interface{}{}

	for _, w := range qq {
		chars[string(w)] = nil
	}

	return len(chars)
}
