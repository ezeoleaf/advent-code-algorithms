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

type game map[int][]int
type turns []int

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

	g := game{}

	scanner := bufio.NewScanner(f)

	scanner.Scan()

	return g.play(scanner.Text())
}

func (g game) play(conf string) int {
	turn := 0
	prevNumber := -1
	var spokenNumber int
	sConf := strings.Split(conf, ",")

	for i := 0; i <= len(sConf); i++ {
		turn++

		if i > len(sConf)-1 {
			spokenNumber = 0
		} else {
			spokenNumber, _ = strconv.Atoi(sConf[i])
		}

		g[spokenNumber] = append(g[spokenNumber], turn)

		prevNumber = spokenNumber
	}

	for {
		turn++
		if _, ok := g[prevNumber]; ok && len(g[prevNumber]) > 1 {
			turns := g[prevNumber]
			spokenNumber = turns[len(turns)-1] - turns[len(turns)-2]
		} else {
			spokenNumber = 0
		}

		g[spokenNumber] = append(g[spokenNumber], turn)

		prevNumber = spokenNumber

		if turn == 2020 {
			return spokenNumber
		}
	}
}
