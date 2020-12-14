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
			} else if chair == "#" && oc >= 5 {
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

//TODO: Improve this
func getOccupiedChairs(lns lines, i, j int) int {
	oc := 0

	// Get chairs in cross(+)
	foundUp, foundDown, foundLeft, foundRight := false, false, false, false
	ix := 0
	for !foundUp || !foundDown { //Up and down
		ix++
		if !foundUp {
			ni := i - ix
			if ni < 0 {
				foundUp = true
			} else {
				c := lns[ni][j]
				if c != "." {
					if c == "#" {
						oc++
					}
					foundUp = true
				}
			}
		}

		if !foundDown {
			ni := i + ix

			if ni > h-1 {
				foundDown = true
			} else {
				c := lns[ni][j]
				if c != "." {
					if c == "#" {
						oc++
					}
					foundDown = true
				}
			}
		}
	}

	ix = 0
	for !foundLeft || !foundRight { //Left and right
		ix++
		if !foundLeft {
			nj := j - ix
			if nj < 0 {
				foundLeft = true
			} else {
				c := lns[i][nj]
				if c != "." {
					if c == "#" {
						oc++
					}
					foundLeft = true
				}
			}
		}

		if !foundRight {
			nj := j + ix

			if nj > w-1 {
				foundRight = true
			} else {
				c := lns[i][nj]
				if c != "." {
					if c == "#" {
						oc++
					}
					foundRight = true
				}
			}
		}
	}

	// Get chairs in x
	foundLU, foundRU, foundLL, foundRL := false, false, false, false
	ix = 0
	for !foundLU || !foundRU { //Left Upper and Right Upper
		ix++
		ni := i - ix

		if ni < 0 {
			foundLU = true
			foundRU = true
			continue
		}

		if !foundLU {
			nj := j - ix
			if nj < 0 {
				foundLU = true
			} else {
				c := lns[ni][nj]
				if c != "." {
					if c == "#" {
						oc++
					}
					foundLU = true
				}
			}
		}

		if !foundRU {
			nj := j + ix
			if nj > w-1 {
				foundRU = true
			} else {
				c := lns[ni][nj]
				if c != "." {
					if c == "#" {
						oc++
					}
					foundRU = true
				}
			}
		}
	}

	ix = 0
	for !foundLL || !foundRL { //Left Lower and Right Lower
		ix++
		ni := i + ix
		if ni > h-1 {
			foundLL = true
			foundRL = true
			continue
		}

		if !foundLL {
			nj := j - ix
			if nj < 0 {
				foundLL = true
			} else {
				c := lns[ni][nj]
				if c != "." {
					if c == "#" {
						oc++
					}
					foundLL = true
				}
			}
		}

		if !foundRL {
			nj := j + ix
			if nj > w-1 {
				foundRL = true
			} else {
				c := lns[ni][nj]
				if c != "." {
					if c == "#" {
						oc++
					}
					foundRL = true
				}
			}
		}
	}

	return oc
}
