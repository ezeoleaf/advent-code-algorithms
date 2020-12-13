package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type operations []string

func main() {
	start := time.Now()

	fmt.Printf("Result is %v \n", run())

	log.Printf("Code took %s", time.Since(start))
}

func run() int {

	ops := operations{}

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ops = append(ops, scanner.Text())
	}

	return getAccumulator(ops)
}

func getAccumulator(ops operations) int {
	var accumulator = 0
	executedOps := make(map[int]interface{})

	regNum, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(ops); i++ {
		o := ops[i]
		if _, ok := executedOps[i]; ok {
			break
		}

		executedOps[i] = nil

		amount, _ := strconv.Atoi(regNum.ReplaceAllString(o, ""))

		log.Println(i, o, accumulator)

		if strings.Contains(o, "nop") {
			continue
		} else if strings.Contains(o, "acc") {
			if strings.Contains(o, "+") {
				accumulator += amount
			} else if strings.Contains(o, "-") {
				accumulator -= amount
			}
		} else if strings.Contains(o, "jmp") {
			if strings.Contains(o, "+") {
				i += amount - 1
			} else if strings.Contains(o, "-") {
				i -= amount + 1
			}
		}
	}

	return accumulator
}
