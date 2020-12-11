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

type stepConf struct {
	StepRight int
	StepDown  int
}

var traverseMap = []stepConf{
	stepConf{StepRight: 1, StepDown: 1},
	stepConf{StepRight: 3, StepDown: 1},
	stepConf{StepRight: 5, StepDown: 1},
	stepConf{StepRight: 7, StepDown: 1},
	stepConf{StepRight: 1, StepDown: 2},
}

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

	return calculatePath(trees)
}

func calculatePath(trees []string) int {
	var result int

	for _, path := range traverseMap {
		calcResult := calculate(trees, path)
		if result == 0 {
			result = calcResult
		} else {
			result *= calcResult
		}
	}

	return result
}

func calculate(trees []string, conf stepConf) int {
	var amountOfTrees, row, column int

	for {
		if row+conf.StepDown >= len(trees) {
			break
		}

		row += conf.StepDown

		rowOfTrees := trees[row]
		column = (column + conf.StepRight) % len(rowOfTrees)

		spot := rowOfTrees[column : column+1]

		if spot == tree {
			amountOfTrees++
		}
	}

	return amountOfTrees
}
