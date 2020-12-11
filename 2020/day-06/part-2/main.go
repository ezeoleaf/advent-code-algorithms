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

	questions := []string{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			questionsAnswered += getQuestionsQuantity(questions)
			questions = []string{}
			continue
		}

		questions = append(questions, scanner.Text())
	}

	questionsAnswered += getQuestionsQuantity(questions)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return questionsAnswered
}

func getQuestionsQuantity(questions []string) int {
	var chars = map[string]int{}
	var valids int

	for _, question := range questions {
		for _, w := range question {
			w := string(w)
			if _, ok := chars[w]; ok {
				chars[w]++
			} else {
				chars[w] = 1
			}
		}
	}

	lenQ := len(questions)
	for _, v := range chars {
		if v == lenQ {
			valids++
		}
	}

	return valids
}
