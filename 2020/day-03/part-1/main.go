package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	tree      = "#"
	stepRight = 3
	stepDown  = 1
)

func main() {
	start := time.Now()

	fmt.Printf("Result is %v \n", run())

	log.Printf("Code took %s", time.Since(start))
}

func run() int {
	trees := []string{}

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		trees = append(trees, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return calculate(trees)
}

func calculate(trees []string) int {
	var amountOfTrees, row, column int

	for {
		if row+stepDown >= len(trees) {
			break
		}

		row += stepDown

		rowOfTrees := trees[row]
		column = (column + stepRight) % len(rowOfTrees)

		spot := rowOfTrees[column : column+1]

		if spot == tree {
			amountOfTrees++
		}
	}

	return amountOfTrees
}
