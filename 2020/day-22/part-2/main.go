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
	finalDeck []int
	score     int
	winner    int
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
	loadingDeck1 := true
	deck1 := []int{}
	deck2 := []int{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if strings.Contains(scanner.Text(), "Player 2") {
			loadingDeck1 = false
			continue
		} else if strings.Contains(scanner.Text(), "Player 1") {
			loadingDeck1 = true
			continue
		}

		v, _ := strconv.Atoi(scanner.Text())
		if loadingDeck1 {
			deck1 = append(deck1, v)
		} else {
			deck2 = append(deck2, v)
		}
	}

	g.winner = g.playGame(deck1, deck2)

	return g.calculateScore()
}

func (g *game) playGame(deck1, deck2 []int) int {
	end := false

	winner := 0
	w := 0

	var plays = map[int]map[string]interface{}{}
	plays[1] = make(map[string]interface{})
	plays[2] = make(map[string]interface{})

	for !end {
		if deck1[0] > len(deck1)-1 || deck2[0] > len(deck2)-1 {
			if deck1[0] > deck2[0] {
				w = 1
			} else {
				w = 2
			}
		} else {
			ndeck1 := []int{}
			ndeck1 = append(ndeck1, deck1[1:1+deck1[0]]...)
			ndeck2 := []int{}
			ndeck2 = append(ndeck2, deck2[1:1+deck2[0]]...)

			w = g.playGame(ndeck1, ndeck2)
		}

		if w == 1 {
			deck1 = append(deck1, deck1[0], deck2[0])
		} else {
			deck2 = append(deck2, deck2[0], deck1[0])
		}

		if len(deck1) == 1 || len(deck2) == 1 {
			end = true
			if len(deck1) == 1 {
				deck2 = deck2[1:]
				winner = 2
			} else {
				deck1 = deck1[1:]
				winner = 1
			}
			continue
		} else {
			deck1 = deck1[1:]
			deck2 = deck2[1:]
		}

		move1 := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(deck1)), "-"), "[]")
		move2 := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(deck2)), "-"), "[]")

		_, ok1 := plays[1][move1]
		_, ok2 := plays[2][move2]

		if ok1 && ok2 {
			end = true
			winner = 1
		} else {
			plays[1][move1] = nil
			plays[2][move2] = nil
		}
	}

	if winner == 1 {
		g.finalDeck = deck1
	} else {
		g.finalDeck = deck2
	}

	return winner
}

func (g *game) calculateScore() int {
	score := 0

	total := len(g.finalDeck)
	for _, v := range g.finalDeck {
		score += v * total
		total--
	}

	fmt.Printf("The winner is %v with a score of %v\n", g.winner, score)

	return score
}
