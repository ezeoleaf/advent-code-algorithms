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

type game struct {
	deck1 []int
	deck2 []int
	score int
}

func main() {
	start := time.Now()

	fmt.Printf("Result is %v \n", run())

	log.Printf("Code took %s", time.Since(start))
}

func run() int {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	g := game{}
	deck1 := true

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if strings.Contains(scanner.Text(), "Player 2") {
			deck1 = false
			continue
		} else if strings.Contains(scanner.Text(), "Player 1") {
			deck1 = true
			continue
		}

		v, _ := strconv.Atoi(scanner.Text())
		if deck1 {
			g.deck1 = append(g.deck1, v)
		} else {
			g.deck2 = append(g.deck2, v)
		}
	}

	g.play()

	return g.calculateScore()
}

func (g *game) play() {
	end := false

	for !end {
		if g.deck1[0] > g.deck2[0] {
			g.deck1 = append(g.deck1, g.deck1[0], g.deck2[0])
		} else {
			g.deck2 = append(g.deck2, g.deck2[0], g.deck1[0])
		}

		if len(g.deck1) == 1 || len(g.deck2) == 1 {
			end = true
			if len(g.deck1) == 1 {
				g.deck1 = []int{}
				g.deck2 = g.deck2[1:]
			} else {
				g.deck2 = []int{}
				g.deck1 = g.deck1[1:]
			}
		} else {
			g.deck1 = g.deck1[1:]
			g.deck2 = g.deck2[1:]
		}
	}
}

func (g *game) calculateScore() int {
	score := 0
	winner := ""
	var d []int

	if len(g.deck1) == 0 {
		winner = "Player 2"
		d = g.deck2
	} else {
		winner = "Player 1"
		d = g.deck1
	}

	total := len(d)
	for _, v := range d {
		score += v * total
		total--
	}

	fmt.Printf("The winner is %s with a score of %v\n", winner, score)

	return score
}
