package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
		rule := strings.ReplaceAll(scanner.Text(), "bags", "")
		rule = strings.ReplaceAll(rule, ".", "")
		rule = strings.ReplaceAll(rule, "bag", "")
		rule = strings.ReplaceAll(rule, " , ", ",")
		validRules = append(validRules, strings.TrimSpace(rule))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return countBags(validRules)
}

func countBags(validRules []string) int {
	// fmt.Println(validRules)
	reg, err := regexp.Compile("[^a-zA-Z] +")
	if err != nil {
		log.Fatal(err)
	}

	allBags := map[string]string{}
	search := true
	for search {
		newBags := map[string]string{}
		for _, rule := range validRules {
			sRule := strings.Split(rule, " contain")
			for k := range bagSearch {
				pk := reg.ReplaceAllString(k, "")
				if strings.Contains(sRule[0], pk) {
					newBags[strings.Split(sRule[0], " bags")[0]] = sRule[1]
				}
			}
		}

		search = false
		for k, v := range newBags {
			k = strings.TrimSpace(k)
			if _, ok := allBags[k]; !ok {
				allBags[k] = v
				sv := strings.Split(v, ",")
				for _, b := range sv {
					bagSearch[strings.TrimSpace(b)] = nil
				}
				search = true
			}
		}
	}

	return calculateBags(allBags) // Removing the shiny gold
}

func calculateBags(bags map[string]string) int {
	fmt.Println(bags)
	startBag := bags["shiny gold"]
	count := 1
	for k, v := range bags {

		fmt.Println(k, v)
	}
	fmt.Println(startBag)
	return count
}
