package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var bagSearch = map[string]interface{}{"shiny gold": nil}

func main() {
	start := time.Now()

	fmt.Printf("Result is %v \n", run())

	log.Printf("Code took %s", time.Since(start))
}

func run() int {
	var validRules = []string{}

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "contain no other bags") {
			continue
		}

		validRules = append(validRules, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return countBags(validRules)
}

func countBags(validRules []string) int {
	search := true
	for search {
		newBags := map[string]interface{}{}
		for _, rule := range validRules {
			sRule := strings.Split(rule, "contain")
			for k := range bagSearch {
				if strings.Contains(sRule[1], k) {
					newBags[strings.Split(sRule[0], " bags")[0]] = nil
				}
			}
		}

		search = false
		for k := range newBags {
			if _, ok := bagSearch[k]; !ok {
				bagSearch[k] = nil
				search = true
			}
		}
	}

	return len(bagSearch) - 1 // Removing the shiny gold
}
