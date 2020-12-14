package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type lines [][]string

var w, h = 0, 0

func main() {
	start := time.Now()

	fmt.Printf("Result is %v \n", run())

	log.Printf("Code took %s", time.Since(start))
}

func run() int {

	lns := lines{}

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		vals := []string{}
		for _, v := range scanner.Text() {
			vals = append(vals, string(v))
		}
		w = len(vals)
		lns = append(lns, vals)
	}

	h = len(lns)

	return calculateSeats(lns)
}

func calculateSeats(lns lines) int {
	move := true
	for move {
		lns, move = iterate(lns)
		// move = false
	}

	return countSeats(lns)
}

func countSeats(lns lines) int {
	seats := 0

	for _, v := range lns {
		for _, chair := range v {
			if chair == "#" {
				seats++
			}
		}
	}

	return seats
}

func iterate(lns lines) (lines, bool) {
	toChange := make(map[string]string)
	changed := false
	for i, v := range lns {
		for j, chair := range v {
			if chair == "." {
				continue
			}

			key := fmt.Sprintf("%v-%v", i, j)
			oc := getOccupiedChairs(lns, i, j)
			if chair == "L" && oc == 0 {
				toChange[key] = "#"
			} else if chair == "#" && oc >= 4 {
				toChange[key] = "L"
			}
		}
	}

	if len(toChange) > 0 {
		changed = true
		for toChangeKey, v := range toChange {
			values := strings.Split(toChangeKey, "-")
			i, _ := strconv.Atoi(values[0])
			j, _ := strconv.Atoi(values[1])

			lns[i][j] = v
		}
	}

	return lns, changed
}

func getOccupiedChairs(lns lines, i, j int) int {
	oc := 0

	for k := -1; k <= 1; k++ {
		for l := -1; l <= 1; l++ {
			ni := i + k
			nj := j + l

			if (ni < 0 || nj < 0) || (ni > h-1 || nj > w-1) {
				continue
			}

			if ni == i && nj == j {
				continue
			}

			c := lns[ni][nj]
			if c == "#" {
				oc++
			}
		}
	}

	return oc
}
